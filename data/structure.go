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
