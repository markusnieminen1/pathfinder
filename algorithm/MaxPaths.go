package algorithm

import "pathfinder/data"
import "fmt"
import "slices"
import "cmp"



func MaxPaths(found_routes *[][]string, firstpath *[]string, secondpath *[]string) {

slices.SortFunc((*found_routes), func(a, b []string) int {
	return cmp.Compare(len(a), len(b))
})

	//dereferencing a pointer with ()
starts := len(data.StationsMap[(*found_routes)[0][0]].Connections)
ends := len(data.StationsMap[(*found_routes)[0][(len((*found_routes)[0])-1)]].Connections)
MaxAmountOfPaths := min(ends, starts)

	//shortest one path "group" always happens, assuming the program stops before if it is invalid, using it to initialize struct
PathGroup := data.GroupedPaths{Paths: [][]string{}, AmountOfPaths: 0, LongestPathInTurns: 0}
PathGroup = AddPathToGroup(PathGroup, (*found_routes)[0])
fmt.Println("Baseline Shortest Single Path ", PathGroup)

TopGroupOfGroupedPaths := data.GroupsOfPaths{Groups: []data.GroupedPaths{}}
TopGroupOfGroupedPaths.Groups = append(TopGroupOfGroupedPaths.Groups, PathGroup)

fmt.Println(MaxAmountOfPaths, "max amount of paths")

if MaxAmountOfPaths > 1{
		// find two and more paths
FindPathCombo(MaxAmountOfPaths, found_routes, &TopGroupOfGroupedPaths)
}

fmt.Println(MaxAmountOfPaths, "max vs valid amount found", len(TopGroupOfGroupedPaths.Groups))
		
fmt.Println("this is the TOP pathgroup ",TopGroupOfGroupedPaths)

}

func FindPathCombo(MaximumPaths int, found_routes *[][]string, TopGroup *data.GroupsOfPaths) {

PathGroupCombo := data.GroupedPaths{Paths: [][]string{}, AmountOfPaths: 0, LongestPathInTurns: 0}

OverlapBool := false

for _, v := range (*found_routes) {

	fmt.Println(v)

	OverlapBool = false
	PathCandidate1 := v
	PathCandidate2 := v

		for _, w := range (*found_routes){

			PathCandidate2 = w

			fmt.Println(w)

			for _, x := range v[1:len(v)-1]{ 	//choose paths to compare against each other

				for _, y := range w[1:len(w)-1]{	//check station by station for overlap

					if x == y{

						OverlapBool = true
					}

				}

			}

			if OverlapBool == false{
				break
			}


		}

		if OverlapBool == false{

			//two path group
			fmt.Println("found non-overlap - OverlapBool = ", OverlapBool)
			fmt.Println(PathCandidate1)
			fmt.Println(PathCandidate2)
			PathGroupCombo = AddPathToGroup(PathGroupCombo, PathCandidate1)
			PathGroupCombo = AddPathToGroup(PathGroupCombo, PathCandidate2)
			fmt.Println(PathGroupCombo)
			TopGroup.Groups = append(TopGroup.Groups, PathGroupCombo) //FIX!!!!!!!make TopGroup global address in main
			
			// if MaximumPaths > 2{
			// 	CurrentPaths := 2
			// 	PathCandidate1 = append(PathCandidate1[1:len(PathCandidate1)-1], PathCandidate2[1:len(PathCandidate2)-1]...)
			// 	fmt.Println("\n BLACKLIST \n",PathCandidate1)
			// 	var Blacklist [][]string{}

			// 	copy_of_route := make([]string, len(PathCandidate1))
			// 	copy(copy_of_route, PathCandidate1)
			// 	Blacklist = append(Blacklist, copy_of_route)

			// 	copy_of_route = make([]string, len(PathCandidate2))
			// 	copy(copy_of_route, PathCandidate2)
			// 	Blacklist = append(Blacklist, copy_of_route)

			// 	var DidWeFindMore bool
			// 	var FoundAdditionalPaths [][]string{}
			// 	DidWeFindMore = FindNonOverlappingPaths(Blacklist, CurrentPaths, MaximumPaths, *found_routes)
			// }

			break
			}	

} //FOR LOOP ENDS

//counter := 1	

fmt.Println("returning THIS SHOULD HAPPEN IF THERE ARE MAX PATHS or LESS THAN MAX DUE TO ONLY OVERLAPPING PATHS, OverlapBool = ", OverlapBool)
return

}

func FindNonOverlappingPaths(Blacklist [][]string, CurrentPaths int, MaximumPaths int, found_routes *[][]string) bool, [][]string{

NonOverlap := true

	for _, v := range (Blacklist) {

		NonOverlap = true

		for j, w := range (*found_routes){

			for _, x := range v[1:len(v)-1]{ 	//choose paths to compare against each other

				for _, y := range w[1:len(w)-1]{	//check station by station for overlap

					if x == y{

						NonOverlap = false
				
					}
				}
			}

			//made it thru
			if NonOverlap {
			
				copy_of_route := make([]string, len((*found_routes)[j]))
				copy(copy_of_route, (*found_routes)[j])
				Blacklist = append(Blacklist, copy_of_route)
				
				if CurrentPaths < MaximumPaths{
				found, finalBlacklist := FindNonOverlappingPaths(Blacklist, (CurrentPaths+1), MaximumPaths, (*found_routes)[(j+1):])
				}
			}
		}
	}

return true

}




func CompareBlacklistToPath(start, end *data.Station, path *[]string, shortest_route_len *int, best_route *[]string, found_routes *[][]string) {

	// Ignore loops
	if start.Visited || len(*path) >= *shortest_route_len {
		return
	}

	start.Visited = true

	*path = append(*path, start.Name)

	// Check if the node we are looking is in the next nodes
	if slices.Contains(start.Connections, end) {

		*path = append(*path, end.Name) // Add end station for the prints etc

		// Should you update the path
		if len(*path) < *shortest_route_len {
			*best_route = make([]string, len(*path))
			copy(*best_route, *path) // CREATE COPY OF THE ITEM, THE PATH WILL CHANGE SO HAVING SAME POINTER WILL CAUSE THE DATA TO BE LOST
			*shortest_route_len = len(*best_route)
		}

		copy_of_route := make([]string, len(*path))
		copy(copy_of_route, *path)

		*found_routes = append(*found_routes, copy_of_route)

		*path = (*path)[:len(*path)-2]
		start.Visited = false

		return
	}

	// Search deeper
	for _, nodeptr := range start.Connections {
		FindPathDFS(nodeptr, end, path, shortest_route_len, best_route, found_routes)
	}

	// Pop last item out and return
	*path = (*path)[:len(*path)-1]
	start.Visited = false

	// return

}


func AddPathToGroup(pathgroup data.GroupedPaths, path []string) data.GroupedPaths {

		pathgroup.AmountOfPaths++

		if len(path)-1 > pathgroup.LongestPathInTurns{
		pathgroup.LongestPathInTurns = len(path)-1
		}

		copy_of_route := make([]string, len(path))
		copy(copy_of_route, path)

		pathgroup.Paths = append(pathgroup.Paths, copy_of_route)

		return pathgroup

}






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


// func NewGroupInit() *data.GroupedPaths {

// return &data.GroupedPaths{Paths: [][]*string{}, AmountOfPaths: 0, LongestPathInTurns: 0}

// }