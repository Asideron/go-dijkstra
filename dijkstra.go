package godijkstra

type Graph struct {
	vertices map[string]*Vertex
}

func NewGraph() *Graph {
	return &Graph{
		vertices: make(map[string]*Vertex),
	}
}

func (g *Graph) AddVertex(name string) error {
	if _, ok := g.vertices[name]; ok {
		return errVertexExists
	}
	g.vertices[name] = NewVertex(name)
	return nil
}

func (g *Graph) vertexExists(name string) bool {
	_, ok := g.vertices[name]
	return ok
}

func (g *Graph) AddEdge(start, end string, price int) error {
	if !g.vertexExists(start) || !g.vertexExists(end) {
		return errNoVertexFound
	}
	g.vertices[start].AddNeighbour(end, price)
	return nil
}

type PathInfo struct {
	Path  []string
	Price int
}

func NewPathInfo() *PathInfo {
	return &PathInfo{
		Path:  make([]string, 0),
		Price: 0,
	}
}

func (g *Graph) FindShortestPath(start, end string) *PathInfo {
	// A table with shortest paths to the end destination.
	cheapestPricesTable := make(map[string]int)

	// A table with cheapest previous vertices for each one to the end.
	cheapestPreviousStopTable := make(map[string]string)

	// Globally accumulates visited cities to choose the cheapest route.
	notVisitedVertices := make(map[string]interface{})
	visitedVertices := make(map[string]interface{})

	cheapestPricesTable[start] = 0

	currentVertex := start
	for currentVertex != "" {
		visitedVertices[currentVertex] = struct{}{}
		delete(notVisitedVertices, currentVertex)

		// Watch neighbours and update cheapest stop and price tables.
		for name, price := range g.vertices[currentVertex].neighbours {
			if _, ok := visitedVertices[name]; !ok {
				notVisitedVertices[name] = struct{}{}
			}
			priceThroughCurrentVertex := cheapestPricesTable[currentVertex] + price
			if _, ok := cheapestPricesTable[name]; !ok || priceThroughCurrentVertex < cheapestPricesTable[name] {
				cheapestPricesTable[name] = priceThroughCurrentVertex
				cheapestPreviousStopTable[name] = currentVertex
			}
		}

		// Find the closest vertex among unvisited ones.
		var minPrice *int = nil
		for name := range notVisitedVertices {
			price := cheapestPricesTable[name]
			if minPrice == nil || *minPrice > price {
				minPrice = &price
				currentVertex = name
			}
		}

		// If no vertex was found - the search is over.
		if minPrice == nil {
			currentVertex = ""
		}
	}

	result := NewPathInfo()
	currentVertex = end
	for currentVertex != start {
		result.Path = append(result.Path, currentVertex)
		name := cheapestPreviousStopTable[currentVertex]
		currentVertex = name
	}
	result.Path = append(result.Path, start)
	result.Price = cheapestPricesTable[end]

	return result
}
