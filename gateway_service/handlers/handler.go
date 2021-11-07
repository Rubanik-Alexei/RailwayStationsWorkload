package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	workloadservice "RailwayStationsWorkload_micro/workload_service"
	"RailwayStationsWorkload_micro/workload_service/protobuff"

	"github.com/gorilla/mux"
	"github.com/xuri/excelize/v2"
	"google.golang.org/grpc"
)

type MyLog struct {
	l *log.Logger
}

func NewLog(l *log.Logger) *MyLog {
	return &MyLog{l}
}

//struct for unmarshaling json response from wl service
type StatResp struct {
	StatName string `json:"RespstationName"`
	StatWl   map[string]struct {
		DayWl map[int32]string `json:"DayWorkload"`
	} `json:"respWorkload"`
	Err string `json:"Error"`
}

func (p *MyLog) GetWorkload(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	flag, err := strconv.ParseBool(vars["dbflag"])
	if err != nil {
		http.Error(rw, "Incorrect last argument", http.StatusBadRequest)
		return
	}
	station := vars["station"]
	conn, err := grpc.Dial(workloadservice.WLport, grpc.WithBlock(), grpc.WithInsecure())
	if err != nil {
		http.Error(rw, "Couldn't connect to service", http.StatusInternalServerError)
		return
	}
	defer conn.Close()
	wl_client := protobuff.NewWorkloadServiceClient(conn)
	stream, err := wl_client.GetStationWorkload(context.Background(), &protobuff.GetStationWorkloadRequest{StationName: station, IsUpdateDB: flag})
	if err != nil {
		http.Error(rw, "Bad response from some service", http.StatusInternalServerError)
		return
	}
	cnt := 2
	//presetting excel file
	xlsx := excelize.NewFile()
	xlsx.NewSheet("Sunday")
	xlsx.NewSheet("Monday")
	xlsx.NewSheet("Tuesday")
	xlsx.NewSheet("Wednesday")
	xlsx.NewSheet("Thursday")
	xlsx.NewSheet("Friday")
	xlsx.NewSheet("Saturday")
	xlsx.DeleteSheet("Sheet1")
	xlsx.SetSheetRow("Sunday", "A1", &[]string{"Станция", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23"})
	xlsx.SetSheetRow("Monday", "A1", &[]string{"Станция", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23"})
	xlsx.SetSheetRow("Tuesday", "A1", &[]string{"Станция", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23"})
	xlsx.SetSheetRow("Wednesday", "A1", &[]string{"Станция", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23"})
	xlsx.SetSheetRow("Thursday", "A1", &[]string{"Станция", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23"})
	xlsx.SetSheetRow("Friday", "A1", &[]string{"Станция", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23"})
	xlsx.SetSheetRow("Saturday", "A1", &[]string{"Станция", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23"})
	//waiting for messages and storing them into excel file
	for {
		feature, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.ListFeatures(_) = _, %v", wl_client, err)
		}
		tmpResp := StatResp{}
		tmp, err := json.Marshal(feature)
		if err != nil {
			fmt.Print(err)
		}
		err = json.Unmarshal(tmp, &tmpResp)
		if err != nil {
			fmt.Print(err)
		}
		// rw.Write([]byte(tmpResp.StatName + "\n"))
		for k, v := range tmpResp.StatWl {
			excl_str := []string{tmpResp.StatName}
			// rw.Write([]byte(k + "\n"))
			for i := int32(0); i <= 23; i++ {
				// rw.Write([]byte(strconv.Itoa(int(i)) + ":"))
				// rw.Write([]byte(v.DayWl[i] + "\n"))
				excl_str = append(excl_str, v.DayWl[i])
			}
			xlsx.SetSheetRow(k, "A"+strconv.Itoa(cnt), &excl_str)
		}
		// rw.Write([]byte(tmpResp.Err))
		// e := json.NewEncoder(rw)
		// e.Encode(feature)
		cnt++
	}
	rw.Header().Set("Content-Disposition", "attachment; filename="+"WL"+".xlsx")
	rw.Header().Set("Content-Transfer-Encoding", "binary")
	rw.Header().Set("Expires", "0")
	xlsx.Write(rw)
}

//unimplemented for now
func (p *MyLog) GetFromDB(rw http.ResponseWriter, req *http.Request) {

}
