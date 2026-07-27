package arguments

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func PrintHelp() {
	fmt.Println("\nUSAGE:")
	fmt.Println("go run . [path to file containing network map] [start station] [end station] [number of trains]")
	fmt.Println("\nExample: go run . test_files/network_example.map waterloo st_pancras 4\n")
	flag.PrintDefaults()
	os.Exit(0)
}

func ReadFlags() {
	flag.BoolVar(&Visualising, "v", false, "Enable result visualisation")
	flag.BoolVar(&AllowInvalidData, "aid", false, "Allow invalid data")

	if len(os.Args) > 5 {
		flag.CommandLine.Parse(os.Args[5:])
	}
}

func ReadArgs() (filename, start_station, end_station string, count int, err error) {

	// network.map waterloo st_pancras 4

	// ARG index 1 = .map file where all routes and stations are declared
	// ARG index 2 = start station
	// ARG index 3 = end station
	// ARG index 4 = train count

	// flags starting from index 5

	ReadFlags()

	if len(os.Args) > 1 && os.Args[1] == "-h" {
		PrintHelp()

	}

	if len(os.Args) < 5 {
		return "", "", "", 0, errors.New("Missing arguments.")
	}

	// Check the file is ok
	filename, err = filepath.Abs(os.Args[1])

	if err != nil {
		return "", "", "", 0, errors.New("Invalid file path: " + os.Args[1])
	}

	f, err := os.Open(filename)

	if err != nil {
		return "", "", "", 0, errors.New("Cannot open input file: " + filename + ". The file does not exist in given location or the permissions are not ok for reading the file.")
	}
	f.Close()

	// start and end existing on the map are validated after the node graph has been initialised.
	start_station = os.Args[2]
	end_station = os.Args[3]

	if len(start_station) < 1 {
		return "", "", "", 0, errors.New("Invalid start station:" + start_station)
	}

	if len(end_station) < 1 {
		return "", "", "", 0, errors.New("Invalid end station" + end_station)
	}

	count, err = strconv.Atoi(os.Args[4])

	if err != nil {
		return "", "", "", 0, errors.New("Given train count is not a valid integer number: " + os.Args[4])
	}

	return
}
