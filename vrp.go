package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: ./vorto <input data fileName>")
	}
	fileName := os.Args[1]
	loads, err := readLoadsFromFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	_ = route(loads)
}

func route(loads []Load) float64 {
	var finalCost float64
	var numDrivers int
	// to keep track of which loads have been processed already
	visited := make([]bool, len(loads))
	numLoadsInitial := len(loads)
	numLoads := numLoadsInitial
	closestLoadIndex := numLoadsInitial // set the initial value to something that's greater than the total no of loads
	closestLoadTime := float64(13 * 60) // since 12 * 60 is the max length of a driver's shift

	for numLoads > 0 {
		listOfLoadIDs := make([]int, 0)
		driver := Driver{
			listOfLoadIDs: listOfLoadIDs,
		}
		var currentTimeDrivenByDriver float64
		// start from the base location
		currentLocation := BaseLocation

		for {
			// reset the index and load time
			closestLoadIndex = numLoadsInitial
			closestLoadTime = 13 * 60

			for index, l := range loads {
				if visited[index] {
					continue
				}
				loadTime := getTimeToTravelBetweenTwoLocations(currentLocation, l.Pickup) + getTimeToTravelBetweenTwoLocations(l.Pickup, l.DropOff)
				timeToGetBackToBase := getTimeToTravelBetweenTwoLocations(l.DropOff, BaseLocation)
				totalTimeFromCurrentLocationToLoadToBase := loadTime + timeToGetBackToBase
				if totalTimeFromCurrentLocationToLoadToBase+currentTimeDrivenByDriver <= MaxDriveTime && loadTime < closestLoadTime {
					// can take current load
					closestLoadIndex = index
					closestLoadTime = loadTime
				}
			}

			// closest job could not be figured out
			if closestLoadIndex == numLoadsInitial {
				break
			}

			// mark the closest load as visited
			visited[closestLoadIndex] = true
			numLoads--
			driver.listOfLoadIDs = append(driver.listOfLoadIDs, closestLoadIndex+1)
			currentTimeDrivenByDriver = currentTimeDrivenByDriver + closestLoadTime
			driver.totalTimeDriven = currentTimeDrivenByDriver
			currentLocation = loads[closestLoadIndex].DropOff
		}

		currentTimeDrivenByDriver = currentTimeDrivenByDriver + getTimeToTravelBetweenTwoLocations(currentLocation, BaseLocation)
		ans := "["
		for _, loadID := range driver.listOfLoadIDs {
			ans += strconv.Itoa(loadID) + ","
		}
		fmt.Println(ans[:len(ans)-1] + "]")
		numDrivers++
		finalCost += currentTimeDrivenByDriver
	}
	return float64(numDrivers*500) + finalCost
}
