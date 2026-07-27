package algorithm

import "pathfinder/data"

func SelectOptimalCombination(allRoutes [][]*data.Station, maxPaths int) [][]*data.Station {

	var bestCombination [][]*data.Station
	var maxCount int
	var minTotalLen int = 1000000

	hasConflict := func(route []*data.Station, selected [][]*data.Station) bool {
		used := make(map[int]bool)
		for _, r := range selected {
			for i := 1; i < len(r)-1; i++ {
				used[r[i].ID] = true
			}
		}
		for i := 1; i < len(route)-1; i++ {
			if used[route[i].ID] {
				return true
			}
		}
		return false
	}

	var backtrack func(startIdx int, current [][]*data.Station)
	backtrack = func(startIdx int, current [][]*data.Station) {
		currentCount := len(current)
		currentTotalLen := 0
		for _, r := range current {
			currentTotalLen += len(r)
		}

		if currentCount > maxCount || (currentCount == maxCount && currentTotalLen < minTotalLen) {
			maxCount = currentCount
			minTotalLen = currentTotalLen

			bestCombination = make([][]*data.Station, len(current))
			copy(bestCombination, current)
		}

		if len(current) >= maxPaths {
			return
		}

		for i := startIdx; i < len(allRoutes); i++ {
			if !hasConflict(allRoutes[i], current) {
				backtrack(i+1, append(current, allRoutes[i]))
			}
		}
	}

	backtrack(0, [][]*data.Station{})
	return bestCombination
}

func FindPathBFS(start, end *data.Station, extraSteps, paths int) [][]*data.Station {
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
				maxPathLen = bestRouteLen + extraSteps
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

	return SelectOptimalCombination(allRoutes, paths)
}
