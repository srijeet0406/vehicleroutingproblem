package main

import (
	"testing"
)

func TestEndToEnd(t *testing.T) {
	loads, err := readLoadsFromFile("resources/inputFile.txt")
	if err != nil {
		t.Fatalf(err.Error())
	}
	cost := route(loads)

	if cost != 1950.7673237494123 {
		t.Errorf("Wrong cost for inputFile.txt. Expected 1950.7673237494123, got %f", cost)
	}
}

func TestGetTimeToTravelBetweenTwoLocations(t *testing.T) {
	pickup := Location{
		x: -50.1,
		y: 80.0,
	}
	dropoff := Location{
		x: 90.1,
		y: 12.2,
	}
	load := Load{
		ID:      1,
		Pickup:  pickup,
		DropOff: dropoff,
	}
	timeFromBaseToPickup := getTimeToTravelBetweenTwoLocations(BaseLocation, load.Pickup)
	timeFromPickupToDropOff := getTimeToTravelBetweenTwoLocations(load.Pickup, load.DropOff)
	timeFromDropOffToBase := getTimeToTravelBetweenTwoLocations(load.DropOff, BaseLocation)
	cost := 1*500 + (timeFromBaseToPickup + timeFromPickupToDropOff + timeFromDropOffToBase)

	loads, err := readLoadsFromFile("resources/inputFileWithOneLoad.txt")
	if err != nil {
		t.Fatalf(err.Error())
	}
	costFromParsingFile := route(loads)

	if cost != costFromParsingFile {
		t.Errorf("Expected the cost calculated by hand (cost1) and the cost calculated by reading input file (cost2) to be the same, but got cost1: %f, cost2: %f", cost, costFromParsingFile)
	}
}
