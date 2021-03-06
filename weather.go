package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Loads weather data from url
func downloadWeatherData(url string, c tConfig) (tWeather, error) {

	var w tWeather

	client := http.Client{}
	resp, errResp := client.Get(url)
	if resp.StatusCode != 200 {
		log.Println(errResp)
		return w, errors.New("Incorrect city name or unable to load weather data")
	}
	defer resp.Body.Close()

	rBody, errRead := ioutil.ReadAll(resp.Body)
	if errRead != nil {
		log.Println(errRead)
		return w, errors.New("Unable to read response")
	}

	errDec := json.Unmarshal(rBody, &w)
	if errDec != nil {
		log.Println(errDec)
		return w, errors.New("Unable to decode response")
	}

	//if location doesn't have city name
	if w.Name == "" {
		w.Name = fmt.Sprintf("%.2f, %.2f", c.Coord.Lon, c.Coord.Lat)
	}

	return w, nil
}
