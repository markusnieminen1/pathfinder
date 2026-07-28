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

// RunScheduler simulates turn-by-turn movements.
// Enforces that no intermediate station is occupied by >1 train at any turn.
func RunScheduler(pathSet [][]string, totalTrains int) [][]string {
	if len(pathSet) == 0 || totalTrains <= 0 {
		return nil
	}

	assignments := DistributeTrains(pathSet, totalTrains)

	var activeTrains []*ActiveTrain
	trainPointer := 0

	// Track next available turn when a path can accept a new train
	pathNextAvailableTurn := make([]int, len(pathSet))

	var turnHistory [][]string
	startStation := pathSet[0][0]
	endStation := pathSet[0][len(pathSet[0])-1]

	for turn := 1; ; turn++ {
		var turnMoves []string
		// Tracks station occupancy for the current turn (StationName -> TrainID)
		stationOccupancy := make(map[string]int)

		// 1. Advance active trains forward
		var remainingActive []*ActiveTrain
		for _, t := range activeTrains {
			t.StepIndex++
			if t.StepIndex < len(t.Path) {
				stationName := t.Path[t.StepIndex]

				// Enforce strict single occupancy for intermediate stations
				if stationName != startStation && stationName != endStation {
					if occupyingTrain, occupied := stationOccupancy[stationName]; occupied {
						panic(fmt.Sprintf("Occupancy Violation at Turn %d: Station '%s' occupied by T%d and T%d", turn, stationName, occupyingTrain, t.ID))
					}
					stationOccupancy[stationName] = t.ID
				}

				turnMoves = append(turnMoves, fmt.Sprintf("T%d-%s", t.ID, stationName))

				if t.StepIndex < len(t.Path)-1 {
					remainingActive = append(remainingActive, t)
				}
			}
		}
		activeTrains = remainingActive

		// 2. Dispatch waiting trains onto paths
		for trainPointer < len(assignments) {
			assign := assignments[trainPointer]
			pIdx := assign.PathIndex

			if pathNextAvailableTurn[pIdx] <= turn {
				firstStation := assign.Path[1]

				// Check station occupancy before dispatching
				if firstStation != startStation && firstStation != endStation {
					if _, occupied := stationOccupancy[firstStation]; occupied {
						break // Station blocked this turn; wait for next turn
					}
					stationOccupancy[firstStation] = assign.TrainID
				}

				t := &ActiveTrain{
					ID:        assign.TrainID,
					Path:      assign.Path,
					StepIndex: 1,
				}
				turnMoves = append(turnMoves, fmt.Sprintf("T%d-%s", t.ID, firstStation))

				if len(t.Path) > 2 {
					activeTrains = append(activeTrains, t)
				}
				pathNextAvailableTurn[pIdx] = turn + 1
				trainPointer++
			} else {
				break
			}
		}

		// Termination check
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
