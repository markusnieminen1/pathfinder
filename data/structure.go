package data

// Basic structure of the node
type Station struct {
	Connections []*Station
	Coordinates [2]int
	Name        string
	Visited     bool
}

var StationsMap map[string]*Station = map[string]*Station{} // Saves pointer to a Station by Station name
var CoordsMap map[[2]int]*Station = map[[2]int]*Station{}   // Saves pointer to a Station by Coordinates

type LinkedList struct {
	NodeGrid []Station
}

//for ComboTrains, should be sorted by shortest path in turns first, and then max amount of paths first
type GroupedPaths struct {
	Paths [][]string
	AmountOfPaths int
	LongestPathInTurns int
}

type GroupsOfPaths struct {
	Groups []GroupedPaths
}
