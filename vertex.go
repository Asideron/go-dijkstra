package godijkstra

type Vertex struct {
	name       string
	neighbours map[string]int
}

func NewVertex(name string) *Vertex {
	return &Vertex{
		name:       name,
		neighbours: make(map[string]int),
	}
}

func (v *Vertex) AddNeighbour(name string, price int) {
	v.neighbours[name] = price
}
