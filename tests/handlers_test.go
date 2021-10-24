package tests

import (
	"RailwayStationsWorkload/handlers"
	"os"
	"testing"
)

func TestReadCsvFile(t *testing.T) {
	url_file := os.Getenv("STATIONSURLS")
	for _, tt := range []struct {
		filepath string
		station  string
		expect   string
	}{
		{filepath: url_file, station: "Новолисино", expect: "https://maps.google.com/?cid=2426643917469029704"},
		{filepath: url_file, station: "новолисино", expect: "Incorrect station name or this station is not supported for now :("},
		{filepath: "WrongPath", station: "Новолисино", expect: "Cannot Open Urls File"},
	} {
		if tmp, _ := handlers.ReadCsvFile(tt.filepath, tt.station); tmp != tt.expect {
			t.Fatalf("expected url of %v to be %s, got %s\n", tt.station, tt.expect, tmp)
		}
	}

}
