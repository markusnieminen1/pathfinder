package algorithm

import (
	"pathfinder/data"
)

// SelectOptimalCombination calculates turns for a path set given train count.
func calculateTurns(paths [][]*data.Station, trainCount int) int {
	if len(paths) == 0 {
		return 1000000
	}
	pathCounts := make([]int, len(paths))
	maxTurns := 0
	for t := 0; t < trainCount; t++ {
		bestIdx := 0
		minCost := (len(paths[0]) - 1) + pathCounts[0]
		for i := 1; i < len(paths); i++ {
			cost := (len(paths[i]) - 1) + pathCounts[i]
			if cost < minCost {
				minCost = cost
				bestIdx = i
			}
		}
		pathCounts[bestIdx]++
		if minCost > maxTurns {
			maxTurns = minCost
		}
	}
	return maxTurns
}

// FindPathBFS finds the optimal set of node-disjoint paths using
// Edmonds-Karp flow augmentation on a node-split residual graph.
func FindPathBFS(start, end *data.Station, trainCount, maxPaths int) [][]*data.Station {
	if start == nil || end == nil || start.ID == end.ID {
		return nil
	}

	// Map stations by ID for fast lookup
	stationByID := make(map[int]*data.Station)
	for _, st := range data.StationsMap {
		stationByID[st.ID] = st
	}

	inNode := func(id int) int { return 2 * id }
	outNode := func(id int) int { return 2*id + 1 }

	startOut := outNode(start.ID)
	endIn := inNode(end.ID)

	adj := make(map[int][]int)
	capacity := make(map[int]map[int]int)
	initialCapacity := make(map[int]map[int]int)

	addEdge := func(u, v, cap int) {
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)

		if capacity[u] == nil {
			capacity[u] = make(map[int]int)
			initialCapacity[u] = make(map[int]int)
		}
		if capacity[v] == nil {
			capacity[v] = make(map[int]int)
			initialCapacity[v] = make(map[int]int)
		}

		capacity[u][v] = cap
		capacity[v][u] = 0
		initialCapacity[u][v] = cap
		initialCapacity[v][u] = 0
	}

	// 1. Build Node-Split Graph
	for _, st := range data.StationsMap {
		if st.ID != start.ID && st.ID != end.ID {
			addEdge(inNode(st.ID), outNode(st.ID), 1)
		}
	}

	// 2. Build Connections
	for _, st := range data.StationsMap {
		for _, conn := range st.Connections {
			if st.ID < conn.ID { // Avoid duplicate edge addition
				addEdge(outNode(st.ID), inNode(conn.ID), 1)
				addEdge(outNode(conn.ID), inNode(st.ID), 1)
			}
		}
	}

	var bestPathSet [][]*data.Station
	minTurns := 1000000

	// 3. Iterative Augmenting Paths Loop
	for k := 1; k <= maxPaths; k++ {
		parent := make(map[int]int)
		visited := make(map[int]bool)

		queue := []int{startOut}
		visited[startOut] = true
		found := false

		for len(queue) > 0 {
			curr := queue[0]
			queue = queue[1:]

			if curr == endIn {
				found = true
				break
			}

			for _, nextNode := range adj[curr] {
				if !visited[nextNode] && capacity[curr][nextNode] > 0 {
					visited[nextNode] = true
					parent[nextNode] = curr
					queue = append(queue, nextNode)
				}
			}
		}

		if !found {
			break // No more disjoint paths available
		}

		// Augment flow along shortest path
		curr := endIn
		for curr != startOut {
			p := parent[curr]
			capacity[p][curr] -= 1
			capacity[curr][p] += 1
			curr = p
		}

		// Extract active disjoint paths
		currentPaths := extractPaths(start, end, capacity, initialCapacity, stationByID)
		turns := calculateTurns(currentPaths, trainCount)

		if turns < minTurns {
			minTurns = turns
			bestPathSet = currentPaths
		} else if turns > minTurns+5 {
			// Stop if adding extra paths starts degrading performance significantly
			break
		}
	}

	return bestPathSet
}

// extractPaths traces active flow edges to reconstruct vertex-disjoint paths.
func extractPaths(start, end *data.Station, capacity, initialCapacity map[int]map[int]int, stationByID map[int]*data.Station) [][]*data.Station {
	var paths [][]*data.Station
	startOut := 2*start.ID + 1
	endIn := 2 * end.ID

	// Find all active outgoing edges from start
	for v, initCap := range initialCapacity[startOut] {
		if initCap == 1 && capacity[startOut][v] == 0 {
			// Follow path from v
			path := []*data.Station{start}
			currNode := v

			for currNode != endIn {
				stID := currNode / 2
				st := stationByID[stID]
				path = append(path, st)

				outN := 2*stID + 1
				// Step to next station's in-node
				nextIn := -1
				for nextCandidate, ic := range initialCapacity[outN] {
					if ic == 1 && capacity[outN][nextCandidate] == 0 {
						nextIn = nextCandidate
						break
					}
				}
				if nextIn == -1 {
					break
				}
				currNode = nextIn
			}

			path = append(path, end)
			paths = append(paths, path)
		}
	}

	return paths
}
