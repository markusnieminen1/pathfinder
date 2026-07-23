package algorithm

import "pathfinder/data"
import "fmt"
import "slices"
import "cmp"

// // // Depth-First-Search 
// func AllPathsThenMaxPaths(start, end *Station, path *[]string, shortest_route_len *int, best_route *[]string, found_routes *[][]string) {

// 	// Ignore loops
// 	if start.Visited || len(*path) >= *shortest_route_len {
// 		return
// 	}

// 	start.Visited = true

// 	*path = append(*path, start.Name)

// 	// Check if the node we are looking is in the next nodes
// 	if slices.Contains(start.Connections, end) {

// 		*path = append(*path, end.Name) // Add end station for the prints etc

// 		// Should you update the path
// 		if len(*path) < *shortest_route_len {
// 			*best_route = make([]string, len(*path))
// 			copy(*best_route, *path) // CREATE COPY OF THE ITEM, THE PATH WILL CHANGE SO HAVING SAME POINTER WILL CAUSE THE DATA TO BE LOST
// 			*shortest_route_len = len(*best_route)
// 		}

// 		copy_of_route := make([]string, len(*path))
// 		copy(copy_of_route, *path)

// 		*found_routes = append(*found_routes, copy_of_route)

// 		*path = (*path)[:len(*path)-2]
// 		start.Visited = false

// 		return
// 	}

// 	// Search deeper
// 	for _, nodeptr := range start.Connections {
// 		AllPathsThenMaxPaths(nodeptr, end, path, shortest_route_len, best_route, found_routes)
// 	}

// 	// Pop last item out and return
// 	*path = (*path)[:len(*path)-1]
// 	start.Visited = false

// 	// return

// }

func MaxPaths(found_routes *[][]string, firstpath *[]string, secondpath *[]string) {

slices.SortFunc((*found_routes), func(a, b []string) int {
	return cmp.Compare(len(a), len(b))
})

for _, v := range (*found_routes) {

	fmt.Println(v)

}

//(len((*found_routes)[0]))

	//dereferencing a pointer with ()
starts := len(data.StationsMap[(*found_routes)[0][0]].Connections)
ends := len(data.StationsMap[(*found_routes)[0][(len((*found_routes)[0])-1)]].Connections)
MaxAmountOfPaths := min(ends, starts) 
fmt.Println(MaxAmountOfPaths)

//shortest one path "group" always happens, using it to initialize struct
PathGroup := data.GroupedPaths{Paths: [][]string{}, AmountOfPaths: 0, LongestPathInTurns: 0}
PathGroup = AddPathToGroup(PathGroup, (*found_routes)[0])
fmt.Println("this is the pathgroup ", PathGroup)

TopGroupOfGroupedPaths := data.GroupsOfPaths{Groups: []data.GroupedPaths{}}
TopGroupOfGroupedPaths.Groups = append(TopGroupOfGroupedPaths.Groups, PathGroup)
fmt.Println("this is the TOP pathgroup ",TopGroupOfGroupedPaths)
}

// func NewGroupInit() *data.GroupedPaths {

// return &data.GroupedPaths{Paths: [][]*string{}, AmountOfPaths: 0, LongestPathInTurns: 0}

// }

func AddPathToGroup(pathgroup data.GroupedPaths, path []string) data.GroupedPaths {

		pathgroup.AmountOfPaths++
		pathgroup.LongestPathInTurns = len(path)-1

		copy_of_route := make([]string, len(path))
		copy(copy_of_route, path)

		pathgroup.Paths = append(pathgroup.Paths, copy_of_route)

		return pathgroup

}