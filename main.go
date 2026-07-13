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

	log.Println(grid.StationsMap["Timberline"])

}
