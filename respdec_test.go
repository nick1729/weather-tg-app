package main

import (
	"encoding/json"
	"os"
	"testing"
)

// Tests decoding json response func
func TestLoadWeatherData(t *testing.T) {

	var w tWeather

	file, errF := os.Open("./tests/respdec_test.json")
	if errF != nil {
		t.Error("Expected:", nil, "got:", errF)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	w = tWeather{}
	errDec := decoder.Decode(&w)
	if errDec != nil {
		t.Error("Expected:", nil, "got:", errDec)
	}
}
