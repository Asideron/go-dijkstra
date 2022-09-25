package godijkstra

import "errors"

var (
	errVertexExists  = errors.New("vertex already exists")
	errNoVertexFound = errors.New("no vertex found")
)
