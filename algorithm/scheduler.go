package algorithm

import (
	"fmt"
	"strings"
)

// PathAssignment maps a train ID to its assigned station path
type PathAssignment struct {
	TrainID int
	Path    []string
}

// DistributeTrains assigns each of the N trains to the optimal path in pathSet.
// Paths in pathSet should be vertex-disjoint (except for start and end).
func DistributeTrains(pathSet [][]string, totalTrains int) []PathAssignment {
	pathCounts := make([]int, len(pathSet))
	assignments := make([]PathAssignment, totalTrains)

	for trainID := 1; trainID <= totalTrains; trainID++ {
		bestPathIdx := 0
		minCost := len(pathSet[0]) + pathCounts[0]

		for i := 1; i < len(pathSet); i++ {
			cost := len(pathSet[i]) + pathCounts[i]
			if cost < minCost {
				minCost = cost
				bestPathIdx = i
			}
		}

		pathCounts[bestPathIdx]++
		assignments[trainID-1] = PathAssignment{
			TrainID: trainID,
			Path:    pathSet[bestPathIdx],
		}
	}

	return assignments
}

// RunScheduler simulates train movements turn-by-turn and outputs "L1-station L2-station" per line.
func RunScheduler(pathSet [][]string, totalTrains int) {
	if len(pathSet) == 0 || totalTrains <= 0 {
		return
	}

	assignments := DistributeTrains(pathSet, totalTrains)

	type ActiveTrain struct {
		ID        int
		Path      []string
		StepIndex int
	}

	var activeTrains []*ActiveTrain
	trainPointer := 0

	// Track next available turn when a path can accept a new train
	pathNextAvailableTurn := make([]int, len(pathSet))

	// Map path start->end identifier to index
	pathIndexMap := make(map[string]int)
	for idx, p := range pathSet {
		key := p[0] + "->" + p[len(p)-1] + fmt.Sprintf("-%d", idx)
		pathIndexMap[key] = idx
	}

	for turn := 1; ; turn++ {
		var turnMoves []string

		// 1. Advance active trains forward
		var remainingActive []*ActiveTrain
		for _, t := range activeTrains {
			t.StepIndex++
			if t.StepIndex < len(t.Path) {
				turnMoves = append(turnMoves, fmt.Sprintf("T%d-%s", t.ID, t.Path[t.StepIndex]))
				if t.StepIndex < len(t.Path)-1 {
					remainingActive = append(remainingActive, t)
				}
			}
		}
		activeTrains = remainingActive

		// 2. Dispatch waiting trains if path is available this turn
		for trainPointer < len(assignments) {
			assign := assignments[trainPointer]

			// Find matching path index
			pIdx := 0
			for i, p := range pathSet {
				if len(p) == len(assign.Path) && p[0] == assign.Path[0] && p[len(p)-1] == assign.Path[len(assign.Path)-1] {
					pIdx = i
					break
				}
			}

			if pathNextAvailableTurn[pIdx] <= turn {
				t := &ActiveTrain{
					ID:        assign.TrainID,
					Path:      assign.Path,
					StepIndex: 1, // Advance to first station after start
				}
				turnMoves = append(turnMoves, fmt.Sprintf("T%d-%s", t.ID, t.Path[1]))
				if len(t.Path) > 2 {
					activeTrains = append(activeTrains, t)
				}
				pathNextAvailableTurn[pIdx] = turn + 1
				trainPointer++
			} else {
				break
			}
		}

		// Termination condition
		if len(turnMoves) == 0 && trainPointer >= len(assignments) && len(activeTrains) == 0 {
			break
		}

		if len(turnMoves) > 0 {
			fmt.Println(strings.Join(turnMoves, " "))
		}
	}
}
