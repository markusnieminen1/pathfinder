package algorithm

import (
	"fmt"
	"strings"
)

// PathAssignment links a train ID directly to its assigned path index and route.
type PathAssignment struct {
	TrainID   int
	PathIndex int
	Path      []string
}

// DistributeTrains assigns totalTrains to paths in pathSet to minimize total turns.
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
			TrainID:   trainID,
			PathIndex: bestPathIdx,
			Path:      pathSet[bestPathIdx],
		}
	}

	return assignments
}

// ActiveTrain tracks a train traversing its assigned path.
type ActiveTrain struct {
	ID        int
	Path      []string
	StepIndex int
}

// RunScheduler simulates turn-by-turn train movements.
func RunScheduler(pathSet [][]string, totalTrains int) [][]string {
	if len(pathSet) == 0 || totalTrains <= 0 {
		return nil
	}

	assignments := DistributeTrains(pathSet, totalTrains)

	var activeTrains []*ActiveTrain
	trainPointer := 0

	// Tracks the next turn when each path can accept a new train dispatch
	pathNextAvailableTurn := make([]int, len(pathSet))

	var turnHistory [][]string

	for turn := 1; ; turn++ {
		var turnMoves []string

		// 1. Advance active trains forward by 1 step
		var remainingActive []*ActiveTrain
		for _, t := range activeTrains {
			t.StepIndex++
			if t.StepIndex < len(t.Path) {
				stationName := t.Path[t.StepIndex]
				turnMoves = append(turnMoves, fmt.Sprintf("T%d-%s", t.ID, stationName))

				// Keep train active until it reaches the destination (last index)
				if t.StepIndex < len(t.Path)-1 {
					remainingActive = append(remainingActive, t)
				}
			}
		}
		activeTrains = remainingActive

		// 2. Dispatch waiting trains onto their assigned paths
		for trainPointer < len(assignments) {
			assign := assignments[trainPointer]
			pIdx := assign.PathIndex

			if pathNextAvailableTurn[pIdx] <= turn {
				firstStation := assign.Path[1]
				turnMoves = append(turnMoves, fmt.Sprintf("T%d-%s", assign.TrainID, firstStation))

				// If path has intermediate stations before the end, track as active train
				if len(assign.Path) > 2 {
					activeTrains = append(activeTrains, &ActiveTrain{
						ID:        assign.TrainID,
						Path:      assign.Path,
						StepIndex: 1,
					})
				}

				pathNextAvailableTurn[pIdx] = turn + 1
				trainPointer++
			} else {
				break
			}
		}

		// Termination check: all trains dispatched and no active trains on tracks
		if len(turnMoves) == 0 && trainPointer >= len(assignments) && len(activeTrains) == 0 {
			break
		}

		if len(turnMoves) > 0 {
			line := strings.Join(turnMoves, " ")
			fmt.Println(line)
			turnHistory = append(turnHistory, turnMoves)
		}
	}

	return turnHistory
}
