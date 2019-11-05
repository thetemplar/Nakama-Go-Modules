package graphmap

import (
	"errors"
)

type Graph struct {
	Triangles map[int32]*Triangle // A map of all the Triangles in this graph, indexed by their key.
}

// Initializes a new graph.
func New() *Graph {
	return &Graph{Triangles: map[int32]*Triangle{}}
}

// Returns the amount of Triangles contained in the graph.
func (g *Graph) Len() int {
	return len(g.Triangles)
}

// Creates a new Triangle. Returns true if the Triangle was created, false if the key is already in use.
func (g *Graph) Add(key int32) bool {
	if g.get(key) != nil {
		return false
	}

	// create new Triangle and add it to the graph
	g.Triangles[key] = &Triangle{
		Id: key,
		Neighbors: make(map[*Triangle]int),
	}

	return true
}

// Deletes the Triangle with the specified key. Returns false if key is invalid.
func (g *Graph) Delete(key int32) bool {
	// get Triangle in question
	n := g.get(key)
	if n == nil {
		return false
	}

	// iterate over neighbors, remove edges from neighboring Triangles
	for neighbor, _ := range n.Neighbors {
		delete(neighbor.Neighbors, n)
	}

	// delete Triangle
	delete(g.Triangles, key)

	return true
}

// Returns a slice containing all Triangles. The slice is empty if the graph contains no Triangles.
func (g *Graph) GetAll() (all []*Triangle) {
	for _, n := range g.Triangles {
		all = append(all, n)
	}
	return
}

// Returns the Triangle with this key, or nil and an error if there is no Triangle with this key.
func (g *Graph) Get(key int32) (n *Triangle, err error) {
	n = g.get(key)

	if n == nil {
		err = errors.New("graph: invalid key")
	}

	return
}

// Internal function, does NOT lock the graph, should only be used in between RLock() and RUnlock() (or Lock() and Unlock()).
func (g *Graph) get(key int32) *Triangle {
	return g.Triangles[key]
}

// Creates an edge between the Triangles specified by the keys. Returns false if one or both of the keys are invalid or if they are the same.
// If there already is a connection, it is overwritten with the new edge weight.
func (g *Graph) Connect(keyA int32, keyB int32, weight int) bool {
	// reflective edges are forbidden
	if keyA == keyB {
		return false
	}

	// get Triangles and check for validity of keys
	TriangleA := g.get(keyA)
	TriangleB := g.get(keyB)

	if TriangleA == nil || TriangleB == nil {
		return false
	}

	TriangleA.Neighbors[TriangleB] = weight
	TriangleB.Neighbors[TriangleA] = weight

	// success
	return true
}

// Removes an edge connecting the two Triangles. Returns false if one or both of the keys are invalid or if they are the same.
func (g *Graph) Disconnect(keyA int32, keyB int32) bool {
	// recursive edges are forbidden
	if keyA == keyB {
		return false
	}

	// get Triangles and check for validity of keys
	TriangleA := g.get(keyA)
	TriangleB := g.get(keyB)

	if TriangleA == nil || TriangleB == nil {
		return false
	}

	delete(TriangleA.Neighbors, TriangleB)
	delete(TriangleB.Neighbors, TriangleA)

	return true
}

// Returns true and the edge weight if there is an edge between the Triangles specified by their keys. Returns false and 0 if one or both keys are invalid, if they are the same, or if there is no edge between the Triangles.
func (g *Graph) Adjacent(keyA int32, keyB int32) (exists bool, weight int) {
	// sanity check
	if keyA == keyB {
		return
	}

	TriangleA := g.get(keyA)
	if TriangleA == nil {
		return
	}

	TriangleB := g.get(keyB)
	if TriangleB == nil {
		return
	}

	// choose Triangle with less edges (easier to find 1 in 10 than to find 1 in 100)
	if len(TriangleA.Neighbors) < len(TriangleB.Neighbors) {
		// iterate over it's map of edges; when the right Triangle is found, return
		for neighbor, weight := range TriangleA.Neighbors {
			if neighbor == TriangleB {
				return true, weight
			}
		}
	} else {
		// iterate over it's map of edges; when the right Triangle is found, return
		for neighbor, weight := range TriangleB.Neighbors {
			if neighbor == TriangleA {
				return true, weight
			}
		}
	}

	return
}
