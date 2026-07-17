package main

import (
	"log"
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

	shortest := 100_000

	grid.FindPath(grid.StationsMap["Pine_Top"], grid.StationsMap["Windy_Point"], &current_path, &shortest, &so_far_best_path, &found_routes)

}
