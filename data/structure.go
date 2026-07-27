package data

// Basic structure of the node
type Station struct {
	Connections []*Station
	Coordinates [2]int
	Name        string
	Visited     bool
	ID          int
}

type BfsQue struct {
	Station *Station
	Path    []*Station
}

var StationsMap map[string]*Station = map[string]*Station{} // Saves pointer to a Station by Station name
var CoordsMap map[[2]int]*Station = map[[2]int]*Station{}   // Saves pointer to a Station by Coordinates
var MAX_X_COORDINATE, MAX_Y_COORDINATE, MIN_X_COORDINATE, MIN_Y_COORDINATE int

type VisualisingData struct {
	Graph             *map[string]*Station
	ViewBox           string
	Start             *Station
	End               *Station
	Fastest_route_ids *[]Station
}

type SearchEvent struct {
	Station_Id int  `json:"i"`
	Visited    bool `json:"v"`
}

var Events []SearchEvent

// Function that can be swapped
var RecordEvent = func(id int, visited bool) {
	Events = append(Events, SearchEvent{Station_Id: id, Visited: visited})
}

func noopRecordEvent(id int, visited bool) {}

func SetLoggingEnabled(enabled bool) {

	if !enabled {
		RecordEvent = noopRecordEvent
	}
}

type LinkedList struct {
	NodeGrid []Station
}

// for ComboTrains, should be sorted by shortest path in turns first, and then max amount of paths first
type GroupedPaths struct {
	Paths              [][]string
	AmountOfPaths      int
	LongestPathInTurns int
}

type GroupsOfPaths struct {
	Groups []GroupedPaths
}
