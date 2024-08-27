package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

// readLoadsFromFile reads the list of loads from an input file
func readLoadsFromFile(fileName string) ([]Load, error) {
	nodes := make([]Load, 0)
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
		}
	}(file)

	scanner := bufio.NewScanner(file)
	scanner.Scan() // skip the first header line
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, " ")
		if len(tokens) != 3 {
			log.Fatalf("Invalid line in file %s", fileName)
		}
		id, err := strconv.Atoi(tokens[0])
		if err != nil {
			return nil, err
		}
		source := tokens[1][1 : len(tokens[1])-1]
		sourceCoordinates := strings.Split(source, ",")
		destination := tokens[2][1 : len(tokens[2])-1]
		destinationCoordinates := strings.Split(destination, ",")
		if len(sourceCoordinates) != 2 || len(destinationCoordinates) != 2 {
			log.Fatalf("Invalid format for source or destination in file %s", fileName)
		}
		sourceX, err := strconv.ParseFloat(sourceCoordinates[0], 64)
		if err != nil {
			log.Fatalf("Invalid value for source coordinate in file %s, error: %v", fileName, err)
		}
		sourceY, err := strconv.ParseFloat(sourceCoordinates[1], 64)
		if err != nil {
			log.Fatalf("Invalid value for source coordinate in file %s, error: %v", fileName, err)
		}
		destinationX, err := strconv.ParseFloat(destinationCoordinates[0], 64)
		if err != nil {
			log.Fatalf("Invalid value for destination coordinate in file %s, error: %v", fileName, err)
		}
		destinationY, err := strconv.ParseFloat(destinationCoordinates[1], 64)
		if err != nil {
			log.Fatalf("Invalid value for destination coordinate in file %s, error: %v", fileName, err)
		}
		sourceLocation := Location{
			x: sourceX,
			y: sourceY,
		}
		destinationLocation := Location{
			x: destinationX,
			y: destinationY,
		}
		node := Load{
			ID:      id,
			Pickup:  sourceLocation,
			DropOff: destinationLocation,
		}
		nodes = append(nodes, node)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return nodes, nil
}
