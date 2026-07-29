package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"pathfinder/algorithm"
	"pathfinder/arguments"
	"pathfinder/data"
	"pathfinder/grid"
	"pathfinder/visualising"
	"syscall"
)

func main() {
	abs_filepath, start_station_name, end_station_name, train_count, err := arguments.ReadArgs()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: "+err.Error())
		os.Exit(1)
	}

	err = grid.InitGrid(abs_filepath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: "+err.Error())
		os.Exit(1)
	}

	start_station, data_exists := data.StationsMap[start_station_name]
	if !data_exists {
		fmt.Fprintln(os.Stderr, "Error: Given start station does not exist.")
		os.Exit(1)
	}

	end_station, data_exists := data.StationsMap[end_station_name]
	if !data_exists {
		fmt.Fprintln(os.Stderr, "Error: Given end station does not exist.")
		os.Exit(1)
	}

	var possible_paths int = len(start_station.Connections)
	if len(start_station.Connections) > len(end_station.Connections) {
		possible_paths = len(end_station.Connections)
	}

	if possible_paths < 1 {
		fmt.Fprintln(os.Stderr, "Error: No route between given start&end stations")
		os.Exit(1)
	}

	data.SetLoggingEnabled(arguments.Visualising)

	Paths := algorithm.FindPathBFS(start_station, end_station, train_count, possible_paths)

	var pathSet [][]string
	if len(Paths) > 0 {
		for _, pathPtr := range Paths {
			var singlePath []string
			for _, st := range pathPtr {
				singlePath = append(singlePath, st.Name)
			}
			pathSet = append(pathSet, singlePath)
		}
	} else {
		fmt.Fprintln(os.Stderr, "Error: No paths found.")
		os.Exit(1)
	}

	var turns [][]string
	if len(pathSet) > 0 && train_count > 0 {
		turns = algorithm.RunScheduler(pathSet, train_count)
	}

	if arguments.Visualising {
		ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
		defer cancel()

		var fastest_path []data.Station
		for _, stPtr := range Paths[0] {
			if stPtr != nil {
				fastest_path = append(fastest_path, *stPtr)
			}
		}

		visualising.InitWeb(ctx, start_station, end_station, fastest_path, turns)
	}
}
