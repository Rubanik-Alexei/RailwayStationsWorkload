package handlers

import (
	"RailwayStationsWorkload_micro/config"
	redisservice "RailwayStationsWorkload_micro/internal/redis_service"
	redisProtobuff "RailwayStationsWorkload_micro/internal/redis_service/protobuff"
	workloadservice "RailwayStationsWorkload_micro/internal/workload_service"
	"RailwayStationsWorkload_micro/internal/workload_service/protobuff"
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	//"github.com/gorilla/mux"
	"github.com/gorilla/mux"
	"github.com/xuri/excelize/v2"
	"google.golang.org/grpc"
)

//Paths to hmtls
const wlTemplateName = "internal/gateway_service/htmls/wl.html"
const dbTemplateName = "internal/gateway_service/htmls/db.html"

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
type DBResp struct {
	StatName string `json:"StationName"`
	Workload string `json:"Workload"`
}

func (p *MyLog) LoadingForm(rw http.ResponseWriter, req *http.Request) {
	_, cancel := context.WithTimeout(req.Context(), 5*time.Second)
	defer cancel()
	t, err := template.ParseFiles(wlTemplateName)
	if err != nil {
		http.Error(rw, "Cannot parse form", http.StatusBadRequest)
		return
	}
	t.Execute(rw, "Enter station's names separated by colon")
}
func (p *MyLog) LoadingFormIfErr(rw http.ResponseWriter, req *http.Request) {
	_, cancel := context.WithTimeout(req.Context(), 5*time.Second)
	defer cancel()
	vars := mux.Vars(req)
	t, err := template.ParseFiles(wlTemplateName)
	if err != nil {
		http.Error(rw, "Cannot parse form", http.StatusBadRequest)
		return
	}
	t.Execute(rw, "The following stations either been typed wrong or just unimplemented:"+vars["ErrStations"])
}

func createExcelWorkloadFile() *excelize.File {
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
	return xlsx
}

//For now every request get a one minute timeout

func (p *MyLog) GetWorkload(rw http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(req.Context(), 60*time.Second)
	defer cancel()
	Err_stations := ""
	t, err := template.ParseFiles(wlTemplateName)
	if err != nil {
		http.Error(rw, "Cannot parse html", http.StatusBadRequest)
		return
	}
	var station string
	var dbflag bool
	req.ParseForm()

	if len(req.Form["Stations"][0]) == 0 && len(req.Form["AllStationsFlag"]) == 0 {
		t.Execute(rw, "No stations was entered! ")
		return
	}
	fmt.Println("Station:", req.Form["Stations"])
	station = req.Form["Stations"][0]
	if len(req.Form["DBFlag"]) > 0 {
		if req.Form["DBFlag"][0] == "AddToDB" {
			dbflag = true
			_ = dbflag
		} else {
			http.Error(rw, "Wrong request", http.StatusBadRequest)
			return
		}
	}
	// if err != nil {
	// 	http.Error(rw, "Incorrect last argument", http.StatusBadRequest)
	// 	return
	// }
	conn, err := grpc.Dial(config.WLport, grpc.WithBlock(), grpc.WithInsecure())
	if err != nil {
		http.Error(rw, "Couldn't connect to service", http.StatusInternalServerError)
		return
	}
	defer conn.Close()
	wl_client := protobuff.NewWorkloadServiceClient(conn)
	var stream protobuff.WorkloadService_GetStationWorkloadClient
	if len(req.Form["AllStationsFlag"]) > 0 && req.Form["AllStationsFlag"][0] == "AllStations" {
		allStat, err := workloadservice.ReturnAllStations(os.Getenv("STATIONSURLS"))
		if err != nil {
			http.Error(rw, "Couldn't get all stations", http.StatusInternalServerError)
			return
		}
		stream, err = wl_client.GetStationWorkload(ctx, workloadservice.CreateWorkloadRequest(allStat, dbflag))
	} else {
		stream, err = wl_client.GetStationWorkload(ctx, workloadservice.CreateWorkloadRequest(station, dbflag))
		if err != nil {
			http.Error(rw, "Bad response from some service", http.StatusInternalServerError)
			return
		}
	}
	cnt := 2
	xlsx := createExcelWorkloadFile()
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
		if len(tmpResp.StatWl) == 0 || tmpResp.Err != "OK" {
			Err_stations += tmpResp.StatName + ", "
		} else {
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
	}
	if cnt == 2 {
		t.Execute(rw, "All requested stations either have wrong names or unimplemented")
		return
	}
	fmt.Println(Err_stations)
	rw.Header().Set("Content-Disposition", "attachment; filename="+"WL"+".xlsx")
	rw.Header().Set("Content-Type", req.Header.Get("Content-Type"))
	rw.Header().Set("Content-Transfer-Encoding", "binary")
	rw.Header().Set("Expires", "0")
	xlsx.Write(rw)
	// if Err_stations != "" {
	// 	http.Redirect(rw, req, "/wl/"+Err_stations, http.StatusSeeOther)
	// }

	// 	t.Execute(rw, "Wrong names or unimplemented for these stations: "+Err_stations)
	// 	return
	// } else {
	// 	t.Execute(rw, "All workloads have been successfully collected for required stations ")
	// 	return
	// }
}

func (p *MyLog) LoadingDBForm(rw http.ResponseWriter, req *http.Request) {
	t, err := template.ParseFiles(dbTemplateName)
	if err != nil {
		http.Error(rw, "Cannot parse form", http.StatusBadRequest)
		return
	}
	t.Execute(rw, "Enter station's names separated by colon")
}

//unimplemented for now
func (p *MyLog) GetFromDB(rw http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(req.Context(), 60*time.Second)
	defer cancel()
	t, err := template.ParseFiles(dbTemplateName)
	if err != nil {
		http.Error(rw, "Cannot parse html", http.StatusBadRequest)
		return
	}
	var station string
	req.ParseForm()
	if len(req.Form["Stations"][0]) == 0 {
		t.Execute(rw, "No stations was entered! ")
		return
	}
	// logic part of log in
	fmt.Println("Station:", req.Form["Stations"])
	station = req.Form["Stations"][0]
	conn, err := grpc.Dial(config.Redisport, grpc.WithBlock(), grpc.WithInsecure())
	if err != nil {
		http.Error(rw, "Couldn't connect to service", http.StatusInternalServerError)
		return
	}
	defer conn.Close()
	redis_client := redisProtobuff.NewRedisServiceClient(conn)
	stream, err := redis_client.SearchWorkload(ctx, redisservice.CreateSearchRequest(station))
	if err != nil {
		http.Error(rw, "Bad response from some service", http.StatusInternalServerError)
		return
	}
	cnt := 2
	xlsx := createExcelWorkloadFile()
	//waiting for messages and storing them into excel file
	for {
		feature, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.ListFeatures(_) = _, %v", redis_client, err)
		}
		tmpResp := &DBResp{}
		tmp, err := json.Marshal(feature)
		if err != nil {
			fmt.Print(err)
		}
		err = json.Unmarshal(tmp, &tmpResp)
		if err != nil {
			fmt.Print(err)
		}
		// rw.Write([]byte(tmpResp.StatName + "\n"))
		tmpWl := &StatResp{}
		err = json.Unmarshal([]byte(tmpResp.Workload), &tmpWl.StatWl)
		if err != nil {
			fmt.Print(err)
		}
		for k, v := range tmpWl.StatWl {
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

	if cnt == 2 {
		t.Execute(rw, "All requested stations either have wrong names or unimplemented")
		return
	}
	rw.Header().Set("Content-Disposition", "attachment; filename="+"WL"+".xlsx")
	rw.Header().Set("Content-Type", req.Header.Get("Content-Type"))
	rw.Header().Set("Content-Transfer-Encoding", "binary")
	rw.Header().Set("Expires", "0")
	xlsx.Write(rw)
}
