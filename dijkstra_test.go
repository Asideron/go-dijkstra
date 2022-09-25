package godijkstra_test

import (
	"testing"

	dijkstra "github.com/Asideron/go-dijkstra"
)

func TestShortesPath(t *testing.T) {
	g := dijkstra.NewGraph()

	g.AddVertex("Atlanta")
	g.AddVertex("Boston")
	g.AddVertex("Chicago")
	g.AddVertex("Denver")
	g.AddVertex("El Paso")

	g.AddEdge("Atlanta", "Boston", 100)
	g.AddEdge("Atlanta", "Denver", 160)
	g.AddEdge("Boston", "Chicago", 120)
	g.AddEdge("Boston", "Denver", 180)
	g.AddEdge("Chicago", "El Paso", 80)
	g.AddEdge("Denver", "Chicago", 40)
	g.AddEdge("Denver", "El Paso", 140)

	pathInfo := g.FindShortestPath("Atlanta", "El Paso")

	expectedPrice := 280
	if pathInfo.Price != expectedPrice {
		t.Errorf("Wrong price. Expected %d. Got %d.", expectedPrice, pathInfo.Price)
	}

	expectedPath := []string{"Atlanta", "Denver", "Chicago", "El Paso"}
	if len(expectedPath) != len(pathInfo.Path) {
		t.Errorf("Wrong path length. Expected %d. Got %d.", len(expectedPath), len(pathInfo.Path))
	}
	for i := range pathInfo.Path {
		if pathInfo.Path[i] != expectedPath[len(expectedPath)-1-i] {
			t.Errorf("Wrong path. Expected %v. Got %v.", expectedPath, pathInfo.Path)
		}
	}
}
