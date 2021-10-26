package handlers

//request examples for testing by hand
//grpcurl --plaintext -d '{"stationName":"Мшинская","isUpdateDB":true}' localhost:9092 MyService.GetStationWorkload
//grpcurl --plaintext -d '{"stationName":"Мшинская"}' localhost:9092 MyService.GetStationWorkloadFromDB
import (
	"RailwayStationsWorkload/protobuff"
	"encoding/csv"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type MyError struct {
	error_msg string
}

func (m MyError) Error() string {
	return m.error_msg
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
func GetMap(uurl string, wait time.Duration) (map[string]*protobuff.DayWork, error) {
	client := &http.Client{}
	result := make(map[string]*protobuff.DayWork)
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
	body, err1 := ioutil.ReadAll(resp.Body)
	if err1 != nil {
		return result, err1
	}
	str_body := string(body)
	//Now we have the body of our station page and starting to retrieve data

	//fmt.Println(str_body)
	split_body := strings.Split(str_body, ",")
	//fmt.Println(split_body)

	//Finding the begging of workload info
	re := regexp.MustCompile(`\[\[\[7`)
	start_ind := FindIndex(re, 0, split_body)
	if start_ind == 0 {
		return result, MyError{error_msg: "Cannot find workload on google map page"}
	}
	//Finding the ending of workload info
	re1 := regexp.MustCompile(`0\]\]`)
	end_ind := FindIndex(re1, start_ind, split_body)
	if end_ind == 0 {
		return result, MyError{error_msg: "Cannot find workload on google map page"}
	}
	cnt := 0
	days := []string{"Sunday", "Monday", "Thuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
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
				result[tmp] = &protobuff.DayWork{DayWorkload: tmp_map}
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
