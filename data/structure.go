package data

// Basic structure of the node
type Station struct {
	Connections []*Station
	Coordinates [2]int
	Name        string
	Visited     bool
	ID          int
}

var StationsMap map[string]*Station = map[string]*Station{} // Saves pointer to a Station by Station name
var CoordsMap map[[2]int]*Station = map[[2]int]*Station{}   // Saves pointer to a Station by Coordinates
var MAX_X_COORDINATE, MAX_Y_COORDINATE, MIN_X_COORDINATE, MIN_Y_COORDINATE int

type VisualisingData struct {
	Graph   *map[string]*Station
	ViewBox string
	Start   *Station
	End     *Station
}

type SearchEvent struct {
	Station_Id int  `json:"i"`
	Visited    bool `json:"v"`
}

var Events []SearchEvent

type LinkedList struct {
	NodeGrid []Station
}
