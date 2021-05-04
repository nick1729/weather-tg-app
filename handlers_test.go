package main

import (
	"errors"
	"fmt"
	"testing"
)

// Tests printWeather func
func TestPrintWeather(t *testing.T) {

	var (
		c         tConfig
		w         tWeather
		err       error
		expd, got string
	)

	// testing print with °F
	c.Units = "imperial"
	w.Name = "NY"
	w.Main.Temp = 9.2
	w.Main.TempMin = 8.7
	w.Main.TempMax = 9.6
	err = nil
	expd = "Selected region: NY\nTemperature: +9.2°F\nMinimum: +8.7°F\nMaximum: +9.6°F"

	got = printWeather(c, w, err)
	if got != expd {
		t.Error("Expected:", expd, "got:", got)
	}

	// testing print with °C
	c.Units = "metric"
	w.Name = "LA"
	w.Main.Temp = -4.2
	w.Main.TempMin = -3.6
	w.Main.TempMax = -4.8
	err = nil
	expd = "Selected region: LA\nTemperature: -4.2°C\nMinimum: -3.6°C\nMaximum: -4.8°C"

	got = printWeather(c, w, err)
	if got != expd {
		t.Error("Expected:", expd, "got:", got)
	}

	// testing print with error #1
	c = tConfig{}
	w = tWeather{}
	err = errors.New("Incorrect city name or unable to load weather data")
	expd = "Error! Incorrect city name or unable to load weather data"

	got = printWeather(c, w, err)
	if got != expd {
		t.Error("Expected:", expd, "got:", got)
	}

	// testing print with error #2
	c = tConfig{}
	w = tWeather{}
	err = errors.New("Unable to read response")
	expd = "Error! Unable to read response"

	got = printWeather(c, w, err)
	if got != expd {
		t.Error("Expected:", expd, "got:", got)
	}

	// testing print with error #3
	c = tConfig{}
	w = tWeather{}
	err = errors.New("Unable to decode response")
	expd = "Error! Unable to decode response"

	got = printWeather(c, w, err)
	if got != expd {
		t.Error("Expected:", expd, "got:", got)
	}
}

// Tests setCoordinates func
func TestSetCoordinates(t *testing.T) {

	var (
		c, gotC, expdC        tConfig
		args, expdMsg, gotMsg string
		expdCity, gotCity     bool
	)

	// testing correct coordinates #1
	args = "17.62, 37.79"
	expdC.Coord.Lon = 17.62
	expdC.Coord.Lat = 37.79
	expdCity = false
	expdMsg = "Installed coordinates:\nLongitude: 17.62\nLatitude: 37.79"

	gotC, gotCity, gotMsg = setCoordinates(c, args)
	if gotC != expdC || gotCity != expdCity || gotMsg != expdMsg {
		t.Error("Expected:", expdC, gotCity, expdMsg, "got:", gotC, gotCity, gotMsg)
	}

	// testing correct coordinates #2
	args = "179.96, -89.72"
	expdC.Coord.Lon = 179.96
	expdC.Coord.Lat = -89.72
	expdCity = false
	expdMsg = "Installed coordinates:\nLongitude: 179.96\nLatitude: -89.72"

	gotC, gotCity, gotMsg = setCoordinates(c, args)
	if gotC != expdC || gotCity != expdCity || gotMsg != expdMsg {
		t.Error("Expected:", expdC, gotCity, expdMsg, "got:", gotC, gotCity, gotMsg)
	}

	// testing incorrect coordinates #1
	args = "-310.79, -32.17"
	expdC.Coord.Lon = 0.0
	expdC.Coord.Lat = 0.0
	expdCity = true
	expdMsg = "Incorrect longitude!"

	gotC, gotCity, gotMsg = setCoordinates(c, args)
	if gotC != expdC || gotCity != expdCity || gotMsg != expdMsg {
		t.Error("Expected:", expdC, gotCity, expdMsg, "got:", gotC, gotCity, gotMsg)
	}

	// testing incorrect coordinates #2
	args = "hw, -82.27"
	expdC.Coord.Lon = 0.0
	expdC.Coord.Lat = 0.0
	expdCity = true
	expdMsg = "Incorrect longitude!"

	gotC, gotCity, gotMsg = setCoordinates(c, args)
	if gotC != expdC || gotCity != expdCity || gotMsg != expdMsg {
		t.Error("Expected:", expdC, gotCity, expdMsg, "got:", gotC, gotCity, gotMsg)
	}

	// testing incorrect coordinates #3
	args = "-12.67, 192.32"
	expdC.Coord.Lon = 0.0
	expdC.Coord.Lat = 0.0
	expdCity = true
	expdMsg = "Incorrect latitude!"

	gotC, gotCity, gotMsg = setCoordinates(c, args)
	if gotC != expdC || gotCity != expdCity || gotMsg != expdMsg {
		t.Error("Expected:", expdC, gotCity, expdMsg, "got:", gotC, gotCity, gotMsg)
	}

	// testing incorrect coordinates #4
	args = "-79.12, qwerty"
	expdC.Coord.Lon = 0.0
	expdC.Coord.Lat = 0.0
	expdCity = true
	expdMsg = "Incorrect latitude!"

	gotC, gotCity, gotMsg = setCoordinates(c, args)
	if gotC != expdC || gotCity != expdCity || gotMsg != expdMsg {
		t.Error("Expected:", expdC, gotCity, expdMsg, "got:", gotC, gotCity, gotMsg)
	}
}

// Tests setConvUnits func
func TestSetConvUnits(t *testing.T) {

	var (
		c, gotC, expdC     tConfig
		w, gotW, expdW     tWeather
		s, expdMsg, gotMsg string
	)

	// testing metric to imperial convertation
	c.Units = "metric"
	w.Main.Temp = 7.7
	s = "imperial"
	expdW.Main.Temp = 45.9
	expdC.Units = "imperial"
	expdMsg = "Installed imperial measurement units"

	gotC, gotW, gotMsg = setConvUnits(c, w, s)
	if gotC != expdC || fmt.Sprintf("%+.1f", gotW.Main.Temp) != fmt.Sprintf("%+.1f", expdW.Main.Temp) || gotMsg != expdMsg {
		t.Error("Expected:", expdC, fmt.Sprintf("%+.1f", expdW.Main.Temp), expdMsg,
			"got:", gotC, fmt.Sprintf("%+.1f", gotW.Main.Temp), gotMsg)
	}

	// testing imperial to metric convertation
	c.Units = "imperial"
	w.Main.Temp = 45.9
	s = "metric"
	expdW.Main.Temp = 7.7
	expdC.Units = "metric"
	expdMsg = "Installed metric measurement units"

	gotC, gotW, gotMsg = setConvUnits(c, w, s)
	if gotC != expdC || fmt.Sprintf("%+.1f", gotW.Main.Temp) != fmt.Sprintf("%+.1f", expdW.Main.Temp) || gotMsg != expdMsg {
		t.Error("Expected:", expdC, fmt.Sprintf("%+.1f", expdW.Main.Temp), expdMsg,
			"got:", gotC, fmt.Sprintf("%+.1f", gotW.Main.Temp), gotMsg)
	}

	// testing incorrect argument #1
	s = "hello world"
	expdW.Main.Temp = w.Main.Temp
	expdC.Units = c.Units
	expdMsg = "Wrong arguments. Must be [metric] or [imperial]"

	gotC, gotW, gotMsg = setConvUnits(c, w, s)
	if gotC != expdC || fmt.Sprintf("%+.1f", gotW.Main.Temp) != fmt.Sprintf("%+.1f", expdW.Main.Temp) || gotMsg != expdMsg {
		t.Error("Expected:", expdC, fmt.Sprintf("%+.1f", expdW.Main.Temp), expdMsg,
			"got:", gotC, fmt.Sprintf("%+.1f", gotW.Main.Temp), gotMsg)
	}

	// testing incorrect argument #2
	s = ""
	expdW.Main.Temp = w.Main.Temp
	expdC.Units = c.Units
	expdMsg = "Wrong arguments. Must be [metric] or [imperial]"

	gotC, gotW, gotMsg = setConvUnits(c, w, s)
	if gotC != expdC || fmt.Sprintf("%+.1f", gotW.Main.Temp) != fmt.Sprintf("%+.1f", expdW.Main.Temp) || gotMsg != expdMsg {
		t.Error("Expected:", expdC, fmt.Sprintf("%+.1f", expdW.Main.Temp), expdMsg,
			"got:", gotC, fmt.Sprintf("%+.1f", gotW.Main.Temp), gotMsg)
	}
}
