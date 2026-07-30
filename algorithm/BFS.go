package algorithm

import (
	"pathfinder/data"
)

// Route pairs a route with a stable index used only for conflict lookups.
type Route struct {
	Index int
	Path  []*data.Station
}

// SAVES A LOT OF TIME TO PRECHECK SOME CASES ESPECIALLY FOR LARGE AMOUNT OF PATHS
// CHECK EXISTENCE USING MAP WITH O(1) TIME
func checkConflicts(routes []*Route) []map[int]bool {

	conflicts := make([]map[int]bool, len(routes))
	stationIncludedBy := map[*data.Station][]int{}

	for _, route := range routes {

		for _, station := range route.Path[1 : len(route.Path)-1] {
			stationIncludedBy[station] = append(stationIncludedBy[station], route.Index)
		}
	}

	for index := range routes {
		conflicts[index] = map[int]bool{}
	}

	for _, includedInPath := range stationIncludedBy {

		// Double nested in order to check all values against each other
		for _, i := range includedInPath {

			for _, k := range includedInPath {

				// I and k for ruling self out
				if i != k {
					conflicts[i][k] = true
				}
			}
		}
	}

	return conflicts
}

// The idea is to try and match combinations until reaching bestMaxLen.
// Fewer are ok, more are not desired (exponential unnecessary workload).
// first check if len is valid, then the current combination is longer than the previous and finally
func RecursiveCost(conflicts []map[int]bool, possibilities []*Route, current []*Route,
	currentLength int, bestMaxLen int, bestCombination *[]*Route, bestTotalLength *int) {

	// Rule out too long combinations !!! woohooo
	if len(current) >= bestMaxLen {
		if len(*bestCombination) == 0 || currentLength < *bestTotalLength {
			*bestCombination = append([]*Route{}, current...)
			*bestTotalLength = currentLength
		}
		return
	}

	if len(current)+len(possibilities) < bestMaxLen {
		return
	}

	for rIndex, route := range possibilities {
		remaining := make([]*Route, 0, len(possibilities))

		// Loop the next items
		for _, other := range possibilities[rIndex+1:] {

			// Check if the route has been declared in BOTH
			if !conflicts[route.Index][other.Index] {
				remaining = append(remaining, other)
			}
		}

		nextCurrent := make([]*Route, len(current), len(current)+1)
		copy(nextCurrent, current)
		nextCurrent = append(nextCurrent, route)

		RecursiveCost(conflicts, remaining, nextCurrent, currentLength+len(route.Path),
			bestMaxLen, bestCombination, bestTotalLength)
	}
}

func FindPathBFS(start, end *data.Station, trainCount, paths int) [][]*data.Station {
	var allRoutes [][]*data.Station

	type QueueItem struct {
		station *data.Station
		path    []*data.Station
		visited map[int]bool
	}

	queue := []QueueItem{{
		station: start,
		path:    []*data.Station{start},
		visited: map[int]bool{start.ID: true},
	}}

	maxPathLen := 1000000
	bestRouteLen := 0
	head := 0

	for head < len(queue) {
		current := queue[head]
		head++

		if len(current.path) > maxPathLen {
			continue
		}

		if current.station.ID == end.ID {
			if bestRouteLen == 0 {
				bestRouteLen = len(current.path)
				maxPathLen = bestRouteLen + trainCount
			}

			routeCopy := make([]*data.Station, len(current.path))
			copy(routeCopy, current.path)
			allRoutes = append(allRoutes, routeCopy)
			continue
		}

		for _, nextStation := range current.station.Connections {
			if current.visited[nextStation.ID] {
				continue
			}

			newPathLen := len(current.path) + 1
			if newPathLen <= maxPathLen {
				newVisited := make(map[int]bool, len(current.visited)+1)
				for k, v := range current.visited {
					newVisited[k] = v
				}
				newVisited[nextStation.ID] = true

				newPath := make([]*data.Station, newPathLen)
				copy(newPath, current.path)
				newPath[len(current.path)] = nextStation

				queue = append(queue, QueueItem{
					station: nextStation,
					path:    newPath,
					visited: newVisited,
				})
			}
		}
	}

	routes := make([]*Route, len(allRoutes))
	for i := range allRoutes {
		routes[i] = &Route{Index: i, Path: allRoutes[i]}
	}

	conflicts := checkConflicts(routes)
	candidates := make([]int, len(allRoutes))

	for i := range candidates {
		candidates[i] = i
	}
	var best []*Route
	var bestTotalLength int

	RecursiveCost(conflicts, routes, nil, 0, trainCount, &best, &bestTotalLength)

	bestRoutes := make([][]*data.Station, len(best))
	for i, r := range best {
		bestRoutes[i] = r.Path
	}

	return bestRoutes

}
