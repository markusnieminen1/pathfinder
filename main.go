package main

import (
	//"fmt"
	"log"
	"pathfinder/algorithm"
	"pathfinder/data"
	"pathfinder/grid"
)

func main() {
	err := grid.InitGrid("test_files/ai_generated_example.map")

	if err != nil {
		log.Fatalln(err)
	}

	current_path := []string{}
	so_far_best_path := []string{}
	found_routes := [][]string{}
	BFS_path := []string{}
	first_path := []string{}
	second_path := []string{}

	shortest := 100_000

	algorithm.FindPathDFS(data.StationsMap["Pine_Top"], data.StationsMap["Windy_Point"], &current_path, &shortest, &so_far_best_path, &found_routes)

	algorithm.BreadthFirstSearchStations(data.StationsMap["Pine_Top"], data.StationsMap["Windy_Point"], &BFS_path)

	algorithm.MaxPaths(&found_routes, &first_path, &second_path)

	//fmt.Println(so_far_best_path)
	//fmt.Println(BFS_path)
}
