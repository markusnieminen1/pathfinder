package algorithm

import (
	"container/list"
	"pathfinder/grid"
	"strings"
)

// func TestBFS() {

// 	testing := CreateTestGraph()

// 	fmt.Println("testing Neighbors function")
// 	fmt.Println(testing.Neighbors("euston"))
// 	fmt.Println("testing BFS function")
// 	fmt.Println(BreadthFirstSearchStations(testing, "victoria", "euston"))

// }

// type Graph map[string]GraphNode

// type GraphNode struct {
// 	Name    string
// 	Edges   map[string]GraphNode
// 	X       int
// 	Y       int
// 	Visited bool
// }

// func NewGraph() Graph {
// 	return Graph{}
// }

// func (g Graph) AddNode(name string) {
// 	g[name] = GraphNode{
// 		Name:  name,
// 		Edges: make(map[string]GraphNode),
// 	}
// }

// // AddEdge : adds a directional edge
// func (g Graph) AddEdge(StartStation string, ConnectStation string) {
// 	g[StartStation].Edges[ConnectStation] = g[ConnectStation]
// }

// // Neighbors : prints all neighbors for station
// func (g Graph) Neighbors(nameGiven string) []string {
// 	neighbors := []string{}
// 	for _, theGraphNode := range g {
// 		for edge := range theGraphNode.Edges {
// 			// from nameGiven station connecting to other stations
// 			if theGraphNode.Name == nameGiven {
// 				neighbors = append(neighbors, edge)
// 			}

// 			// from other stations connecting to nameGiven station
// 			// if edge == nameGiven {
// 			// 	neighbors = append(neighbors, theGraphNode.Name)
// 			// }
// 		}
// 	}
// 	return neighbors
// }

// Nodes : returns a list of node IDs
// func (g *Graph) Nodes() []int {
//     nodes := make([]int, len(g.nodes))
//     for i := range g.nodes {
//         nodes[i] = i
//     }
//     return nodes
// }

// func CreateTestGraph() Graph {
// 	TestGraph := NewGraph()
// 	TestGraph.AddNode("waterloo")
// 	TestGraph.AddNode("victoria")
// 	TestGraph.AddNode("st_pancras")
// 	TestGraph.AddNode("euston")
// 	TestGraph.AddEdge("waterloo", "victoria")
// 	TestGraph.AddEdge("waterloo", "euston")
// 	TestGraph.AddEdge("victoria", "waterloo")
// 	TestGraph.AddEdge("victoria", "st_pancras")
// 	TestGraph.AddEdge("st_pancras", "victoria")
// 	TestGraph.AddEdge("st_pancras", "euston")
// 	TestGraph.AddEdge("euston", "waterloo")
// 	TestGraph.AddEdge("euston", "st_pancras")
// 	return TestGraph
// }

// grid.StationsMap["Pine_Top"]

func BreadthFirstSearchStations(root, target *grid.Station, path *[]string) {
	//const notFound = "not_found" // default value

	// check if root and target exist in the tree

	// initialize the queue and push the root node
	queue := list.New()
	queue.PushBack(root)

	// create a parent map to save the interactions and recreate the path
	parents := make(map[string]string) // initialize queue
	parents[root.Name] = ""            // initialize root without any parents

	// while queue has elements, keep iterating
	for queue.Len() > 0 {
		currentNode := queue.Front().Value.(*grid.Station) // get first element
		queue.Remove(queue.Front())                        // remove first element from queue

		// compare if node is equals to target
		if strings.EqualFold(currentNode.Name, target.Name) {
			// the target has been looked
			// reconstructing the path

			for len(currentNode.Name) > 0 {
				// recreating route
				*path = append([]string{currentNode.Name}, *path...)
				// changing pointer
				currentNode.Name = parents[currentNode.Name]
			}

			// returning path result
			// strings.Join(route, "->")
		}

		// iterate neighbors
		for _, edge := range currentNode.Connections {
			// check if the neighbor has not already been visited
			if _, visited := parents[edge.Name]; !visited {
				parents[edge.Name] = currentNode.Name // add neighbor to parents map associated to current node value
				queue.PushBack(edge)                  // add neighbor to the end of the queue
			}
		}
	}

	// return notFound
}

// func BreadthFirstSearchStations(tree map[string]GraphNode, root, target string) string {
// 	const notFound = "not_found" // default value

// 	// check if root and target exist in the tree
// 	rootNode, rootExists := tree[root]
// 	_, targetExists := tree[target]
// 	if !rootExists || !targetExists {
// 		return notFound
// 	}

// 	// initialize the queue and push the root node
// 	queue := list.New()
// 	queue.PushBack(rootNode)

// 	// create a parent map to save the interactions and recreate the path
// 	parents := make(map[string]string) // initialize queue
// 	parents[root] = ""                 // initialize root without any parents

// 	// while queue has elements, keep iterating
// 	for queue.Len() > 0 {
// 		currentNode := queue.Front().Value.(GraphNode) // get first element
// 		queue.Remove(queue.Front())                    // remove first element from queue

// 		// compare if node is equals to target
// 		if strings.EqualFold(currentNode.Name, target) {
// 			// the target has been looked
// 			// reconstructing the path
// 			var route []string
// 			for len(currentNode.Name) > 0 {
// 				// recreating route
// 				route = append([]string{currentNode.Name}, route...)
// 				// changing pointer
// 				currentNode.Name = parents[currentNode.Name]
// 			}

// 			// returning path result
// 			return strings.Join(route, "->")
// 		}

// 		// iterate neighbors
// 		for edge := range currentNode.Edges {
// 			// check if the neighbor has not already been visited
// 			if _, visited := parents[edge]; !visited {
// 				parents[edge] = currentNode.Name // add neighbor to parents map associated to current node value
// 				queue.PushBack(tree[edge])       // add neighbor to the end of the queue
// 			}
// 		}
// 	}

// 	return notFound
// }

// import (
// 	"slices"
// )

// // grid.StationsMap["Pine_Top"]

// // Depth-First-Search
// func ChangeIntoThisFormFindPath(start, end *Station, path *[]string, shortest_route_len *int, best_route *[]string, found_routes *[][]string) {

// 	// Ignore loops
// 	if start.Visited || len(*path) >= *shortest_route_len {
// 		return
// 	}

// 	start.Visited = true

// 	*path = append(*path, start.Name)

// 	// Check if the node we are looking is in the next nodes
// 	if slices.Contains(start.Connections, end) {

// 		*path = append(*path, end.Name) // Add end station for the prints etc

// 		// Should you update the path
// 		if len(*path) < *shortest_route_len {
// 			*best_route = make([]string, len(*path))
// 			copy(*best_route, *path) // CREATE COPY OF THE ITEM, THE PATH WILL CHANGE SO HAVING SAME POINTER WILL CAUSE THE DATA TO BE LOST
// 			*shortest_route_len = len(*best_route)
// 		}

// 		copy_of_route := make([]string, len(*path))
// 		copy(copy_of_route, *path)

// 		*found_routes = append(*found_routes, copy_of_route)

// 		*path = (*path)[:len(*path)-2]
// 		start.Visited = false

// 		return
// 	}

// 	// Search deeper
// 	for _, nodeptr := range start.Connections {
// 		FindPath(nodeptr, end, path, shortest_route_len, best_route, found_routes)
// 	}

// 	// Pop last item out and return
// 	*path = (*path)[:len(*path)-1]
// 	start.Visited = false

// 	// return

// }
