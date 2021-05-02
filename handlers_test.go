package main

import (
	"errors"
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
