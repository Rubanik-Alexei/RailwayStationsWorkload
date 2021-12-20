package workloadservice

import (
	redisworkload "RailwayStationsWorkload_micro/redis_service"
	redisProtobuff "RailwayStationsWorkload_micro/redis_service/protobuff"
	wlProtobuff "RailwayStationsWorkload_micro/workload_service/protobuff"
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
)

var reStart, reEnd, reWorkload = regexp.MustCompile(`\[\[\[7`), regexp.MustCompile(`0\]\]`), regexp.MustCompile("загруженность")

type Server struct {
	log hclog.Logger
	wlProtobuff.UnimplementedWorkloadServiceServer
}

func NewMyServer(l hclog.Logger) *Server {
	return &Server{l, wlProtobuff.UnimplementedWorkloadServiceServer{}}
}

type MyError struct {
	error_msg string
}

func (m MyError) Error() string {
	return m.error_msg
}

//function to isolate internal structure of service
func CreateWorkloadRequest(station string, dbflag bool) *wlProtobuff.GetStationWorkloadRequest {
	return &wlProtobuff.GetStationWorkloadRequest{StationName: station, IsUpdateDB: dbflag}
}

//Checking if required station name is available for scrapping workload
func ReadCsvFile(filePath string, station string) (string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return "Cannot Open Urls File", err
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		return "Cannot parse Urls file as CSV", err
	}
	for _, v := range records {
		//fmt.Println(v)
		//if name is found then return it's google maps url
		if v[5] == station {
			return v[2], nil
		}
	}
	return "Incorrect station name or this station is not supported for now :(", MyError{error_msg: ""}
}

func ReturnAllStations(filePath string) (string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return "Cannot Open Urls File", err
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		return "Cannot parse Urls file as CSV", err
	}
	allStations := ""
	for _, v := range records {
		//fmt.Println(v)
		//if name is found then return it's google maps url
		allStations += v[5] + ","
	}
	return allStations[:len(allStations)-1], nil
}

//helper function for finding beginning/ending of workload
func FindIndex(re *regexp.Regexp, start_ind int, split_body []string) int {
	for i := start_ind; i < len(split_body); i++ {
		if re.MatchString(split_body[i]) {
			return i
		}
	}
	return 0
}

//Collecting workload
func GetMap(uurl string, wait time.Duration) (map[string]*wlProtobuff.DayWork, error) {
	client := &http.Client{}
	result := make(map[string]*wlProtobuff.DayWork)
	req, err := http.NewRequest("GET", uurl, nil)
	if err != nil {
		return result, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.54 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	resp, err := client.Do(req)
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()
	time.Sleep(wait)
	body, err1 := io.ReadAll(resp.Body)
	if err1 != nil {
		return result, err1
	}
	str_body := string(body)
	if !reWorkload.MatchString(str_body) {
		os.WriteFile("resp", []byte(str_body), 0666)
		return result, MyError{error_msg: "No workload at this moment"}
	}
	//Now we have the body of our station page and starting to retrieve data

	//fmt.Println(str_body)
	split_body := strings.Split(str_body, ",")
	//fmt.Println(split_body)

	//Finding the begging of workload info

	//re := regexp.MustCompile(`\[\[\[7`)

	start_ind := FindIndex(reStart, 0, split_body)
	if start_ind == 0 {
		return result, MyError{error_msg: "Cannot find workload on google map page"}
	}
	//Finding the ending of workload info

	//re1 := regexp.MustCompile(`0\]\]`)

	end_ind := FindIndex(reEnd, start_ind, split_body)
	if end_ind == 0 {
		return result, MyError{error_msg: "Cannot find workload on google map page"}
	}
	cnt := 0
	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	//Loop for retrieving hour:percentage pairs for each day
	for i := start_ind + 1; i <= end_ind; {
		endDay := true
		tmp_map := make(map[int32]string)
		for endDay {
			tmp := days[cnt]
			str_hour := strings.TrimLeft(split_body[i], "[")
			hour, err := strconv.Atoi(str_hour)
			if err != nil {
				return result, err
			}
			percentage := split_body[i+1]
			tmp_map[int32(hour)] = percentage
			if hour == 3 {
				endDay = false
				result[tmp] = &wlProtobuff.DayWork{DayWorkload: tmp_map}
				cnt++
				//skipping to next day
				i += 9
				break
			}
			//skipping to next pair
			i += 7
		}
	}
	return result, nil
}

func (s *Server) GetStationWorkload(req *wlProtobuff.GetStationWorkloadRequest, srv wlProtobuff.WorkloadService_GetStationWorkloadServer) error {
	stations_array := strings.Split(req.GetStationName(), ",")
	var wg sync.WaitGroup
	for _, v := range stations_array {
		wg.Add(1)
		go func(v string) {
			resp_msg := "OK"
			dbflag := req.GetIsUpdateDB()
			defer wg.Done()
			url_file := os.Getenv("STATIONSURLS")
			url, err := ReadCsvFile(url_file, v)
			if err != nil {
				s.log.Error(url, "error", err)
				srv.Send(&wlProtobuff.StationData{RespstationName: v, RespWorkLoad: map[string]*wlProtobuff.DayWork{}, Error: err.Error()})
			}
			res, err := GetMap(url, 2)
			if err != nil {
				s.log.Error(url, "error", err)
				srv.Send(&wlProtobuff.StationData{RespstationName: v, RespWorkLoad: map[string]*wlProtobuff.DayWork{}, Error: err.Error()})
			}
			//Making call to Redis service if needed
			for dbflag == true {
				//Needed to marshal response struct to json to be able to store it
				tmpres, err := json.Marshal(res)
				if err != nil {
					resp_msg = "Cannot parse response"
					break
				}
				//initializing redisService Client and making request to it
				redisConn, err := grpc.Dial(redisworkload.Redisport, grpc.WithBlock(), grpc.WithInsecure())
				if err != nil {
					resp_msg = "Couldn't connect to service"
					break
				}
				defer redisConn.Close()
				redis_client := redisProtobuff.NewRedisServiceClient(redisConn)
				redisresp, err := redis_client.StoreWorkload(context.Background(), redisworkload.CreateStoreRequest(v, string(tmpres)))
				if err != nil {
					resp_msg = redisresp.Error
					break
				}
				fmt.Println(v)
				dbflag = false
			}
			resp := &wlProtobuff.StationData{RespstationName: v, RespWorkLoad: res, Error: resp_msg}
			if err := srv.Send(resp); err != nil {
				log.Printf("send error %v", err)
			}
			s.log.Info("Send workload for station : %s", v)
			//fmt.Println("Added to DB")
		}(v)
		//time.Sleep(5 * time.Second)
	}
	wg.Wait()
	return nil
}
