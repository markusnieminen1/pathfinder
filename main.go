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

	shortest := 10_000

	grid.FindPath(grid.StationsMap["Golden_Pillar"], grid.StationsMap["Grand_Cargo_Yard"], &current_path, &shortest, &so_far_best_path, &found_routes)

	log.Println(so_far_best_path)
	log.Println(found_routes)
	// grid.ExportTopologyHTML("real_example.html")

}
