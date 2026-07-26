package main

import (
	"context"
	"fmt"
	"log"
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

	abs_filepath, start_station_name, end_station_name, _, err := arguments.ReadArgs()

	if err != nil {
		// fmt.Println("Invalid arguments. " + err.Error())
		fmt.Print("ERROR: ")
		fmt.Println(err.Error())
		arguments.PrintHelp()
	}

	found_routes := [][]string{}

	current_path_DFS := []data.Station{}
	so_far_best_path_DFS := []data.Station{}
	found_routes_DFS := [][]data.Station{}

	BFS_path := []string{}
	first_path := []string{}
	second_path := []string{}
	err = grid.InitGrid(abs_filepath)

	if err != nil {
		log.Fatalln("Failed to Initialise the grid: " + err.Error())
	}

	start_station, data_exists := data.StationsMap[start_station_name]

	if !data_exists {
		log.Fatalln("Start station does not exist in the map! ('" + start_station_name + "')")
	}

	end_station, data_exists := data.StationsMap[end_station_name]

	if !data_exists {
		log.Fatalln("End station does not exist in the map! ('" + end_station_name + "')")
	}

	algorithm.MaxPaths(&found_routes, &first_path, &second_path)
	data.SetLoggingEnabled(arguments.Visualising)

	// Do not look further than this amount of nodes
	shortest := 10_000

	algorithm.FindPathDFS(start_station, end_station, &current_path_DFS, &shortest, &so_far_best_path_DFS, &found_routes_DFS)

	algorithm.BreadthFirstSearchStations(start_station, end_station, &BFS_path)

	// RUN WEBSERVER
	if arguments.Visualising {
		ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
		defer cancel()
		visualising.InitWeb(ctx, start_station, end_station, &so_far_best_path_DFS)
	}
}
