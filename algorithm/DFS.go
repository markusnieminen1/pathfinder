package algorithm

import (
	"pathfinder/data"
)

// Depth-First-Search. The function is not optimised so looping will go through all possible combinations.
// It would be better to rule out certain paths e.g. known deadends and maybe try calculate some best halfway nodes, so that
// the possibilities decrease a lot.
func FindPathDFS(start, end *data.Station, path *[]data.Station,
	shortest_route_len *int, best_route *[]data.Station, found_routes *[][]data.Station) {

	if start.Visited || len(*path) >= *shortest_route_len {
		return
	}

	start.Visited = true
	data.RecordEvent(start.ID, true)
	*path = append(*path, *start)

	if start.ID == end.ID {
		if len(*path) < *shortest_route_len {
			*best_route = make([]data.Station, len(*path))
			copy(*best_route, *path)
			*shortest_route_len = len(*best_route)
		}
		copy_of_route := make([]data.Station, len(*path))
		copy(copy_of_route, *path)
		*found_routes = append(*found_routes, copy_of_route)

	} else {
		for _, nodeptr := range start.Connections {
			FindPathDFS(nodeptr, end, path, shortest_route_len, best_route, found_routes)
		}
	}

	*path = (*path)[:len(*path)-1]
	start.Visited = false
	data.RecordEvent(start.ID, false)
}
