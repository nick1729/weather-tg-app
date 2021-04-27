package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

// Creates url by city name or via coordinates and calls downloadWeatherData
func getWeather(c tConfig, byCity bool) (tWeather, error) {

	var (
		w   tWeather
		url string
		err error
	)

	url = fmt.Sprintf("%s/data/2.5/weather?lat=%f&lon=%f&units=%s&lang=%s&appid=%s",
		c.ApiURL, c.Coord.Lat, c.Coord.Lon, c.Units, c.Lang, c.ApiKEY)

	if byCity {
		url = fmt.Sprintf("%s/data/2.5/weather?q=%s&units=%s&lang=%s&appid=%s",
			c.ApiURL, c.City, c.Units, c.Lang, c.ApiKEY)
	}

	w, err = downloadWeatherData(url, c)

	return w, err
}

// Creates output message of the received weather data
func printWeather(c tConfig, w tWeather, err error) string {

	var sign, msg string

	if err != nil {
		return fmt.Sprintf("Error! %s", err.Error())
	}

	if c.Units == "imperial" {
		sign = "°F"
	} else {
		sign = "°C"
	}

	msg = fmt.Sprintf("Selected region: %s\nTemperature: %+.1f%s\nMinimum: %+.1f%s\nMaximum: %+.1f%s",
		w.Name, w.Main.Temp, sign, w.Main.TempMin, sign, w.Main.TempMax, sign)

	return msg
}

// Prints help command
func printHelp() string {

	var msg string

	msg = `Available commands:
/weather - show current weather
/city [city name] - set name of the city
/coordinates [37.62, 55.75] - set city by nearest coordinates (longitude -180..180, latitude -90..90)
/units [metric, imperial] - set measurement units
/lang [ar, cz, de, en, fr, it, ja, kr, nl, pt, ru, sp, tr, ua] - set language`

	return msg
}

// Checking and setting coordinates
func setCoordinates(c tConfig, s string) (tConfig, bool, string) {

	var (
		args []string
		msg  string
	)

	args = strings.Split(s, ", ")
	if len(args) == 2 {
		Lon, errLon := strconv.ParseFloat(args[0], 64)
		if errLon != nil || Lon < -180.0 || Lon > 180.0 {
			log.Println("errLon")
			return c, true, "Incorrect longitude!"
		}
		Lat, errLat := strconv.ParseFloat(args[1], 64)
		if errLat != nil || Lat < -90.0 || Lat > 90.0 {
			log.Println("errLat")
			return c, true, "Incorrect latitude!"
		}
		c.Coord.Lon = Lon
		c.Coord.Lat = Lat

	}

	msg = fmt.Sprint("Installed coordinates:",
		"\nLongitude: ", c.Coord.Lon,
		"\nLatitude: ", c.Coord.Lat)

	return c, false, msg
}

// Setting up and translating units of measurement
func setConvUnits(c tConfig, w tWeather, s string) (tConfig, tWeather, string) {

	var msg string

	if s == "metric" || s == "imperial" {
		switch {
		case c.Units == "metric" && s == "imperial":
			w.Main.Temp = 1.8*w.Main.Temp + 32
			w.Main.TempMin = 1.8*w.Main.TempMin + 32
			w.Main.TempMax = 1.8*w.Main.TempMax + 32
		case c.Units == "imperial" && s == "metric":
			w.Main.Temp = (w.Main.Temp - 32) / 1.8
			w.Main.TempMin = (w.Main.TempMin - 32) / 1.8
			w.Main.TempMax = (w.Main.TempMax - 32) / 1.8
		}
		c.Units = s
		msg = fmt.Sprintf("Installed %s measurement units", c.Units)
	} else {
		msg = "Wrong arguments. Must be [metric] or [imperial]"
	}

	return c, w, msg
}
