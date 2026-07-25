package algorithm

import (
	"pathfinder/data"
	"slices"
)

// Depth-First-Search. The function is not optimised so looping will go through all possible combinations.
// It would be better to rule out certain paths e.g. known deadends and maybe try calculate some best halfway nodes, so that
// the possibilities decrease a lot.
func FindPathDFS(start, end *data.Station, path *[]data.Station,
	shortest_route_len *int, best_route *[]data.Station, found_routes *[][]data.Station) {

	// Ignore loops
	if start.Visited || len(*path) >= *shortest_route_len {
		return
	}

	start.Visited = true
	data.Events = append(data.Events, data.SearchEvent{Station_Id: start.ID, Visited: true})
	*path = append(*path, *start)

	// Check if the node we are looking is in the next nodes
	if slices.Contains(start.Connections, end) {

		*path = append(*path, *end) // Add end station for the prints etc

		// Should you update the path
		if len(*path) < *shortest_route_len {
			*best_route = make([]data.Station, len(*path))
			copy(*best_route, *path) // CREATE COPY OF THE ITEM, THE PATH WILL CHANGE SO HAVING SAME POINTER WILL CAUSE THE DATA TO BE LOST
			*shortest_route_len = len(*best_route)
		}

		copy_of_route := make([]data.Station, len(*path))
		copy(copy_of_route, *path)

		*found_routes = append(*found_routes, copy_of_route)

		*path = (*path)[:len(*path)-2]
		start.Visited = false
		data.Events = append(data.Events, data.SearchEvent{Station_Id: start.ID, Visited: false})

		return
	}

	// Search deeper
	for _, nodeptr := range start.Connections {
		FindPathDFS(nodeptr, end, path, shortest_route_len, best_route, found_routes)
	}

	// Pop last item out and return
	*path = (*path)[:len(*path)-1]
	start.Visited = false
	data.Events = append(data.Events, data.SearchEvent{Station_Id: start.ID, Visited: false})

	// return

}
