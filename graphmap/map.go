package graphmap

import (
	"github.com/bradfitz/slice"
	"math"
)

type Vector struct {
	X 			float32
	Y 			float32
}

type Map struct {
	Borders		[]Edge
	Triangles	[]Triangle
}

type Edge struct {
	A 			Vector
	B 			Vector
}

type Triangle struct {
	A 			Vector
	B 			Vector
	C 			Vector
	W 			Vector
	
	Id       	int32
	Neighbors 	map[*Triangle]int // maps the neighbor node to the weight of the connection to it
}


func (v Vector) distance(t *Vector) float32 {
	return float32(math.Sqrt(math.Pow(float64(t.X - v.X), 2) + math.Pow(float64(t.Y - v.Y), 2)))
}

// Returns the map of neighbors.
func (n *Triangle) GetNeighbors() map[*Triangle]int {
	if n == nil {
		return nil
	}
	
	Neighbors := n.Neighbors

	return Neighbors
}

// Returns the Nodees key.
func (n *Triangle) Key() int32 {
	if n == nil {
		return -1
	}

	key := n.Id
	
	return key
}


type NextTriangles struct{
	Distance float64
	Index 	 int
}

func (m *Map) getNearest(v Vector, length int) []NextTriangles {
	res := make([]NextTriangles, len(m.Triangles))
	i := 0
	for _, triangle := range m.Triangles {
		res[i].Distance = float64(triangle.W.distance(&v))
		res[i].Index = i
		i++
	}
	slice.Sort(res[:], func(i, j int) bool {
		return res[i].Distance < res[j].Distance
	})
	return res[:length]
}

func (t Triangle) IsInTriangle(x, y float32) (bool, float32, float32, float32) {
	// Compute vectors        
	v0_x := t.C.X - t.A.X  
	v0_y := t.C.Y - t.A.Y
	v1_x := t.B.X - t.A.X
	v1_y := t.B.Y - t.A.Y
	v2_x := x     - t.A.X
	v2_y := y     - t.A.Y

	// Compute dot products
	dot00 := v0_x * v0_x + v0_y * v0_y
	dot01 := v0_x * v1_x + v0_y * v1_y
	dot02 := v0_x * v2_x + v0_y * v2_y
	dot11 := v1_x * v1_x + v1_y * v1_y
	dot12 := v1_x * v2_x + v1_y * v2_y

	// Compute barycentric coordinates
	invDenom := 1 / (dot00 * dot11 - dot01 * dot01)
	w1 := (dot11 * dot02 - dot01 * dot12) * invDenom
	w2 := (dot00 * dot12 - dot01 * dot02) * invDenom

	return (w1 >= 0) && (w2 >= 0) && (w1 + w2 < 1), w1, w2, w1+w2
}