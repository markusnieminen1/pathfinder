package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"pathfinder/algorithm"
	"pathfinder/data"
	"pathfinder/grid"
	"pathfinder/visualising"
	"syscall"
)

func main() {
	err := grid.InitGrid("test_files/ai_generated2.map")

	if err != nil {
		log.Fatalln(err)
	}
	start_station := data.StationsMap["Crystal_Crossing"]
	end_station := data.StationsMap["Stone_Row"]
	/*
		current_path := []string{}
		so_far_best_path := []string{}
		found_routes := [][]string{}

		shortest := 100_000

		algorithm.FindPathDFS(start_station, end_station, &current_path, &shortest, &so_far_best_path, &found_routes)
	*/
	BFS_path := []string{}
	algorithm.BreadthFirstSearchStations(start_station, end_station, &BFS_path)
	/*
		BFS_path := []string{}

		fmt.Println(so_far_best_path)
		fmt.Println(BFS_path)
	*/

	// RUN WEBSERVER
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()
	visualising.InitWeb(ctx, start_station, end_station)

}
