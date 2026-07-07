package grid

// Basic structure of the node
type Station struct {
	Connections []*Station
	Coordinates [2]int
	Name        string
}

type LinkedList struct {
	NodeGrid []Station
}
