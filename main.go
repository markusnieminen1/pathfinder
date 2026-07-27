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
		fmt.Fprintln(os.Stderr, "Error")
		os.Exit(1)
	}

	err = grid.InitGrid(abs_filepath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error")
		os.Exit(1)
	}

	start_station, data_exists := data.StationsMap[start_station_name]
	if !data_exists {
		fmt.Fprintln(os.Stderr, "Error")
		os.Exit(1)
	}

	end_station, data_exists := data.StationsMap[end_station_name]
	if !data_exists {
		fmt.Fprintln(os.Stderr, "Error")
		os.Exit(1)
	}

	data.SetLoggingEnabled(arguments.Visualising)

	// Search for path:
	current_path_DFS := []data.Station{}
	so_far_best_path_DFS := []data.Station{}
	found_routes_DFS := [][]data.Station{}
	shortest := 10_000

	algorithm.FindPathDFS(start_station, end_station, &current_path_DFS, &shortest, &so_far_best_path_DFS, &found_routes_DFS)

	// Format best path for scheduler
	var pathSet [][]string
	if len(so_far_best_path_DFS) > 0 {
		var singlePath []string
		for _, st := range so_far_best_path_DFS {
			singlePath = append(singlePath, st.Name)
		}
		pathSet = append(pathSet, singlePath)
	}

	// Run turn-by-turn scheduler
	if len(pathSet) > 0 && train_count > 0 {
		algorithm.RunScheduler(pathSet, train_count)
	}

	if arguments.Visualising {
		ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
		defer cancel()
		visualising.InitWeb(ctx, start_station, end_station, &so_far_best_path_DFS)
	}
}
