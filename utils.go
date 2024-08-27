package main

import "math"

// MaxDriveTime is the maximum length of a driver's shift
const MaxDriveTime = 720

// Location is a struct defining the x and y coordinates of a Location
type Location struct {
	x float64
	y float64
}

// BaseLocation is the start and end location for all drivers
var BaseLocation = Location{
	x: 0,
	y: 0,
}

// Load defines the structure read from the input file
type Load struct {
	ID      int
	Pickup  Location
	DropOff Location
}

// Driver is a struct to store the list of loads that this driver will complete, along with an attribute to define the total time driven to complete these jobs
type Driver struct {
	listOfLoadIDs   []int
	totalTimeDriven float64
}

// getTimeToTravelBetweenTwoLocations is the time (in minutes) it takes for a driver to travel from a given source to a
// given destination
func getTimeToTravelBetweenTwoLocations(source Location, destination Location) float64 {
	return math.Sqrt(math.Pow(destination.x-source.x, 2) + math.Pow(destination.y-source.y, 2))
}
