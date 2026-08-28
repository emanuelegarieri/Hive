package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Measure struct {
	Index 		int
	Key   		string
	Value 		string
	GotValue 	bool	
}


func initStation() []Measure {
	station := []Measure{
		{Index: 1, Key: "airTemp"},
		{Index: 2, Key: "airPressure"},
		{Index: 7, Key: "precipitation"},
		{Index: 11, Key: "windSpeed"},
		{Index: 12, Key: "windDirection"},
		{Index: 13, Key: "humidity"},
		{Index: 14, Key: "dewPoint"},
		{Index: 15, Key: "soilMoisture"},
		{Index: 22, Key: "cloudCover"},
	}
	return station
}


func addMeasure(station []Measure, input string) []Measure {
	// Split the input in Index as int and in string for Value
	indexValue := strings.Split(input, ",") 
	index, err := strconv.Atoi(indexValue[0])
	// err manage
	if err != nil {
		fmt.Printf("Value: invalid input! Input must be string")
	}

	if indexValue[1] == "NULL" {
		for sen := range station {
			if station[sen].Index == index {
				station[sen].Value = "NULL"
				station[sen].GotValue = false
				break
			} 
		}
	} else {
		intDec := strings.Split(indexValue[1], ".")
		Dec, _ := strconv.Atoi(intDec[1])
		value := ""
		if Dec != 0 {
			value = indexValue[1]
		} else {
			value = intDec[0]
		}
		for sen := range station {
			if station[sen].Index == index {
				station[sen].Value = value
				station[sen].GotValue = true
				break
			} 
		}
	}
	return station
}


func getStation(station []Measure) []Measure {
	for sen := range station {
		if station[sen].GotValue == false {
			station[sen].Value = "NULL"
		}
	}
	return station
}


func main() {
	fmt.Println("--- Weather Station ---")
	station := initStation()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())
		switch input {
		case "exit":
			fmt.Println("Exiting...")
			return
		case "clear":
			station = initStation()
		case "get":	
			station = getStation(station)
			for i := range station {
				fmt.Printf("%s:%s\n", station[i].Key, station[i].Value)
			}
		default:
			station = addMeasure(station, input)
		}
	}
}
