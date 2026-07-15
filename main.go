package main

import (
	"container/list"
	"fmt"
	"strings"
)

// type Node struct {
// 	Value     string
// 	Neighbors []string
// }

// type Station struct {
// 	Connections []*Station
// 	Name        string
// }

func main() {

	// tree := map[string]Node{
	// 	"John":     {Value: "John", Neighbors: []string{"George", "Sam", "Edward"}},
	// 	"George":   {Value: "George", Neighbors: []string{"Richard"}},
	// 	"Sam":      {Value: "Sam", Neighbors: []string{"Richard", "Briana"}},
	// 	"Edward":   {Value: "Edward", Neighbors: []string{"Anett", "Shaun"}},
	// 	"Richard":  {Value: "Richard", Neighbors: []string{"Franklin"}},
	// 	"Briana":   {Value: "Briana", Neighbors: []string{"Lynsey", "Karen"}},
	// 	"Anett":    {Value: "Anett", Neighbors: []string{"Wilson"}},
	// 	"Shaun":    {Value: "Shaun", Neighbors: []string{}},
	// 	"Franklin": {Value: "Franklin", Neighbors: []string{}},
	// 	"Lynsey":   {Value: "Lynsey", Neighbors: []string{}},
	// 	"Karen":    {Value: "Karen", Neighbors: []string{}},
	// 	"Wilson":   {Value: "Wilson", Neighbors: []string{}},
	// }

	// var StationsMap map[string]*Station = map[string]*Station{} // Saves pointer to a Station by Station name

	// waterloo := NewStation("waterloo", []*Station{victoria, euston})
	// victoria := NewStation("victoria", []*Station{waterloo, st_pancras})
	// st_pancras := NewStation("st_pancras", []*Station{victoria, euston})
	// euston := NewStation("euston", []*Station{waterloo, st_pancras})

	// StationsMap[waterloo.Name] = waterloo
	// StationsMap[victoria.Name] = victoria
	// StationsMap[st_pancras.Name] = st_pancras
	// StationsMap[euston.Name] = euston

	// tree1 := map[string]*Station{
	//  	"waterloo":     {Name: "waterloo", Connections: []*Station{StationsMap["victoria"], StationsMap["euston"]}},
	// 	"victoria":     {Name: "victoria", Connections: []*Station{StationsMap["waterloo"], StationsMap["st_pancras"]}},
	// 	"st_pancras":     {Name: "st_pancras", Connections: []*Station{StationsMap["victoria"], StationsMap["euston"]}},
	// 	"euston":     {Name: "euston", Connections: []*Station{StationsMap["st_pancras"], StationsMap["waterloo"]}},
	// }
	// jungle := NewStation("jungle")

	// tree2 := map[string]*Station{
	// 	"jungle":     {Name: "jungle", Connections: []*Station{StationsMap["grasslands"], StationsMap["farms"], StationsMap["green_belt"]}},
	// 	"green_belt": {Name: "green_belt", Connections: []*Station{StationsMap["jungle"], StationsMap["village"]}},
	// 	"village":    {Name: "village", Connections: []*Station{StationsMap["green_belt"], StationsMap["mountain"]}},
	// 	"mountain":   {Name: "mountain", Connections: []*Station{StationsMap["farms"], StationsMap["village"], StationsMap["treetop"], StationsMap["wetlands"]}},
	// 	"treetop":    {Name: "treetop", Connections: []*Station{StationsMap["mountain"], StationsMap["desert"]}},
	// 	"grasslands": {Name: "grasslands", Connections: []*Station{StationsMap["jungle"], StationsMap["suburbs"]}},
	// 	"suburbs":    {Name: "suburbs", Connections: []*Station{StationsMap["grasslands"], StationsMap["clouds"]}},
	// 	"clouds":     {Name: "clouds", Connections: []*Station{StationsMap["suburbs"], StationsMap["wetlands"]}},
	// 	"wetlands":   {Name: "wetlands", Connections: []*Station{StationsMap["desert"], StationsMap["mountain"], StationsMap["clouds"]}},
	// 	"farms":      {Name: "farms",	Connections: []*Station{StationsMap["jungle"], StationsMap["downtown"], StationsMap["mountain"]}},
	// 	"downtown":   {Name: "downtown", Connections: []*Station{StationsMap["farms"], StationsMap["metropolis"]}},
	// 	"metropolis": {Name: "metropolis", Connections: []*Station{StationsMap["downtown"], StationsMap["industrial"]}},
	// 	"industrial": {Name: "industrial", Connections: []*Station{StationsMap["desert"], StationsMap["metropolis"]}},
	// 	"desert":     {Name: "desert", Connections: []*Station{StationsMap["wetlands"], StationsMap["treetop"], StationsMap["industrial"]}},
	// }

	//fmt.Println(BreadthFirstSearch(tree, "John", "Anett"))

	//fmt.Println(BreadthFirstSearch(tree1, "jungle", "desert"))

	testing := CreateTestGraph()

	fmt.Println("testing Neighbors function")
	fmt.Println(testing.Neighbors("euston"))

	fmt.Println(BreadthFirstSearchStations(testing, "victoria", "euston"))

}

// func NewStation(data string) *Station {
// 	return &Station{
// 		Connections: nil,
// 		Name:        data,
// 	}

// }

//var graphTest map[string]*GraphNode

type Graph map[string]GraphNode

type GraphNode struct {
	Name    string
	Edges   map[string]GraphNode
	X       int
	Y       int
	Visited bool
}

func NewGraph() Graph {
	return Graph{}
}

func (g Graph) AddNode(name string) {
	g[name] = GraphNode{
		Name:  name,
		Edges: make(map[string]GraphNode),
	}
}

func AddNodeToMainGraph(name string, variable Graph) {
	variable[name] = GraphNode{
		Name: name,
	}
}

// AddEdge : adds a directional edge
func (g Graph) AddEdge(StartStation string, ConnectStation string) {
	g[StartStation].Edges[ConnectStation] = g[ConnectStation]
}

func (g Graph) Neighbors(nameGiven string) []string {
	neighbors := []string{}
	for _, theGraphNode := range g {
		for edge := range theGraphNode.Edges {
			// from nameGiven station connecting to other stations
			if theGraphNode.Name == nameGiven {
				neighbors = append(neighbors, edge)
			}

			// from other stations connecting to nameGiven station
			// if edge == nameGiven {
			// 	neighbors = append(neighbors, theGraphNode.Name)
			// }
		}
	}
	return neighbors
}

// Nodes : returns a list of node IDs
// func (g *Graph) Nodes() []int {
//     nodes := make([]int, len(g.nodes))
//     for i := range g.nodes {
//         nodes[i] = i
//     }
//     return nodes
// }

func CreateTestGraph() Graph {
	TestGraph := NewGraph()
	TestGraph.AddNode("waterloo")
	TestGraph.AddNode("victoria")
	TestGraph.AddNode("st_pancras")
	TestGraph.AddNode("euston")
	TestGraph.AddEdge("waterloo", "victoria")
	TestGraph.AddEdge("waterloo", "euston")
	TestGraph.AddEdge("victoria", "waterloo")
	TestGraph.AddEdge("victoria", "st_pancras")
	TestGraph.AddEdge("st_pancras", "victoria")
	TestGraph.AddEdge("st_pancras", "euston")
	TestGraph.AddEdge("euston", "waterloo")
	TestGraph.AddEdge("euston", "st_pancras")
	return TestGraph
}

func BreadthFirstSearchStations(tree map[string]GraphNode, root, target string) string {
	const notFound = "not_found" // default value

	// check if root and target exist in the tree
	rootNode, rootExists := tree[root]
	_, targetExists := tree[target]
	if !rootExists || !targetExists {
		return notFound
	}

	// initialize the queue and push the root node
	queue := list.New()
	queue.PushBack(rootNode)

	// create a parent map to save the interactions and recreate the path
	parents := make(map[string]string) // initialize queue
	parents[root] = ""                 // initialize root without any parents

	// while queue has elements, keep iterating
	for queue.Len() > 0 {
		currentNode := queue.Front().Value.(GraphNode) // get first element
		queue.Remove(queue.Front())                    // remove first element from queue

		// compare if node is equals to target
		if strings.EqualFold(currentNode.Name, target) {
			// the target has been looked
			// reconstructing the path
			var route []string
			for len(currentNode.Name) > 0 {
				// recreating route
				route = append([]string{currentNode.Name}, route...)
				// changing pointer
				currentNode.Name = parents[currentNode.Name]
			}

			// returning path result
			return strings.Join(route, "->")
		}

		// iterate neighbors
		for edge := range currentNode.Edges {
			// check if the neighbor has not already been visited
			if _, visited := parents[edge]; !visited {
				parents[edge] = currentNode.Name // add neighbor to parents map associated to current node value
				queue.PushBack(tree[edge])       // add neighbor to the end of the queue
			}
		}
	}

	return notFound
}

// func BreadthFirstSearch(tree map[string]Node, root, target string) string {
// 	const notFound = "not_found" // default value

// 	// check if root and target exist in the tree
// 	rootNode, rootExists := tree[root]
// 	_, targetExists := tree[target]
// 	if !rootExists || !targetExists {
// 		return notFound
// 	}

// 	// initialize the queue and push the root node
// 	q := list.New()
// 	q.PushBack(rootNode)

// 	// create a parent map to save the interactions and recreate the path
// 	parents := make(map[string]string) // initialize queue
// 	parents[root] = ""                 // initialize root without any parents

// 	// while queue has elements, keep iterating
// 	for q.Len() > 0 {
// 		currentNode := q.Front().Value.(Node) // get first element
// 		q.Remove(q.Front())                   // remove first element from queue

// 		// compare if node is equals to target
// 		if strings.EqualFold(currentNode.Value, target) {
// 			// the target has been looked
// 			// reconstructing the path
// 			var route []string
// 			for len(currentNode.Value) > 0 {
// 				// recreating route
// 				route = append([]string{currentNode.Value}, route...)
// 				// changing pointer
// 				currentNode.Value = parents[currentNode.Value]
// 			}

// 			// returning path result
// 			return strings.Join(route, "->")
// 		}

// 		// iterate neighbors
// 		for _, neighbor := range currentNode.Neighbors {
// 			// check if the neighbor has not already been visited
// 			if _, visited := parents[neighbor]; !visited {
// 				parents[neighbor] = currentNode.Value // add neighbor to parents map associated to current node value
// 				q.PushBack(tree[neighbor])            // add neighbor to the end of the queue
// 			}
// 		}
// 	}

// 	return notFound
// }
