package graphmap

import (
	"math"
	"fmt"
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
	Neighbors 	map[int32]float32 // maps the neighbor node to the weight of the connection to it
}


func (m *Map) GetTriangle(x, y float32) (bool, int64) {
	for i, triangle := range m.Triangles {
		isItIn, _, _, _ := triangle.IsInTriangle(x, y)
		if isItIn {
			return true, int64(i);
		}
	}
	return false, -1
}

func (v Vector) Distance(t *Vector) float32 {
	return float32(math.Sqrt(math.Pow(float64(t.X - v.X), 2) + math.Pow(float64(t.Y - v.Y), 2)))
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

	w1 = float32(math.Round(float64(w1)*100000)/100000)
	w2 = float32(math.Round(float64(w2)*100000)/100000)
	sum := float32(math.Round(float64(w1+w2)*100000)/100000)

	//fmt.Printf("\tIsInTriangle? (%v %v @ %v %v %v) :  %v  %v  %v = %v\n", x, y, t.A, t.B, t.C, w1, w2, sum, (w1 >= 0) && (w2 >= 0) && (w1 + w2 < 1))

	return (w1 >= 0) && (w2 >= 0) && (sum <= 1), w1, w2, sum
}

func Intersection (p0, p1, p2, p3 *Vector) (bool, Vector) {
	s1_x := p1.X - p0.X
	s1_y := p1.Y - p0.Y
	s2_x := p3.X - p2.X
	s2_y := p3.Y - p2.Y

	s := float64((-s1_y * (p0.X - p2.X) + s1_x * (p0.Y - p2.Y)) / (-s2_x * s1_y + s1_x * s2_y));
	t := float64(( s2_x * (p0.Y - p2.Y) - s2_y * (p0.X - p2.X)) / (-s2_x * s1_y + s1_x * s2_y));
	
	s = math.Round(s*100000)/100000
	t = math.Round(t*100000)/100000
	
	//fmt.Printf("Intersection: %v %v  - %v %v - found: %v %v : %v \n", p0, p1, p2, p3, s, t, (s >= 0 && s <= 1 && t >= 0 && t <= 1))

	if s >= 0 && s <= 1 && t >= 0 && t <= 1 {
        return true, Vector {X: p0.X + (float32(t) * s1_x), Y: p0.Y + (float32(t) * s1_y)};
    }

    return false, Vector {}; 
}

func (t Triangle) GetEdgeTowardsPoint(curX, curY, targetX, targetY float32) (bool, Vector, Edge) {
	if curX == targetX && curY == targetY {
		fmt.Printf("No difference between Points!!! \n")
	}

	found1, point1 := Intersection(&t.A, &t.B, &Vector{X: curX, Y:curY}, &Vector{X: targetX, Y:targetY})
	if found1 {
		return true, point1, Edge {A: t.A, B: t.B}
	}
	found2, point2 := Intersection(&t.A, &t.C, &Vector{X: curX, Y:curY}, &Vector{X: targetX, Y:targetY})
	if found2 {
		return true, point2, Edge {A: t.A, B: t.C}
	}
	found3, point3 := Intersection(&t.B, &t.C, &Vector{X: curX, Y:curY}, &Vector{X: targetX, Y:targetY})
	if found3 {
		return true, point3, Edge {A: t.B, B: t.C}
	}
    return false, Vector {}, Edge {}; 
}

func (e Edge) MinDistance(pointX, pointY float32) (bool, Vector) {	
    point := Vector{
		X: pointX,
		Y: pointY,
	}
	
    AE := Vector{
		X: point.X - e.A.X,
		Y: point.Y - e.A.Y,
	}
	
    AB := Vector{
		X: e.B.X - e.A.X,
		Y: e.B.Y - e.A.Y,
	}
	
    BE := Vector{
		X: point.X - e.B.X,
		Y: point.Y - e.B.Y,
	}

	//Finding the squared magnitude of AB
	ATB2 := float32(math.Pow(float64(AB.X), 2) + math.Pow(float64(AB.Y), 2))

    // Calculating the dot product 
    AB_BE := (AB.X * BE.X + AB.Y * BE.Y)
	AB_AE := (AB.X * AE.X + AB.Y * AE.Y)
	
	t := AB_AE / ATB2

	//fmt.Printf("MinDistance : %v %v  :  %v %v \n", AB_BE, AB_AE,  AB_AE / ATB2, Vector { X: e.A.X + AB.X * t, Y: e.A.Y + AB.Y * t })
  
    // Case 1 
    if (AB_BE > 0) { 
		return true, e.B
    } else if (AB_AE < 0) { 
		return true, e.A
    } else { 
		return false, Vector { X: e.A.X + AB.X * t, Y: e.A.Y + AB.Y * t }
	} 
} 
