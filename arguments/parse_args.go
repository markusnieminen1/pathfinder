package arguments

import (
	"errors"
	"os"
	"pathfinder/grid"
	"strconv"
)

func GetHelp() string {

	return ("USAGE: \n" +
		"go run . [path to file containing network map] [start station] [end station] [number of trains] \n" +
		"Example: go run . network.map waterloo st_pancras 4")
}

func ReadArgs() (filename, start_station, end_station string, count int, err error) {

	// network.map waterloo st_pancras 4

	// ARG index 1 = .map file where all routes and stations are declared
	// ARG index 2 = start station
	// ARG index 3 = end station
	// ARG index 4 = train count

	// flags starting from index 5

	if len(os.Args) < 5 {
		return "", "", "", 0, errors.New("Missing arguments.")
	}

	// Check the file is ok
	filename, err = grid.GetAbsPath(os.Args[1])

	if err != nil {
		return "", "", "", 0, errors.Join(err, errors.New("Invalid file: "+os.Args[1]))
	}

	f, err := os.Open(filename)

	if err != nil {
		return "", "", "", 0, errors.Join(err, errors.New("Cannot open input file: "+os.Args[1]))
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
