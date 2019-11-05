package main

import (
	"math"
	"Nakama-Go-Modules/graphmap"
)

func (v Vector2Df) rotate(degrees float32) Vector2Df {
	ca := float32(math.Cos(float64(360 - degrees) * 0.01745329251)); //0.01745329251
	sa := float32(math.Sin(float64(360 - degrees) * 0.01745329251));

	vec := Vector2Df {
		X: ca * v.X - sa * v.Y,
		Y: sa * v.X + ca * v.Y,
	}
	return vec
}

func (v Vector2Df) length() float32 {
	return float32(math.Sqrt(math.Pow(float64(v.X), 2) + math.Pow(float64(v.Y), 2)))
}

func (v Vector2Df) distance(t *Vector2Df) float32 {
	return float32(math.Sqrt(math.Pow(float64(t.X - v.X), 2) + math.Pow(float64(t.Y - v.Y), 2)))
}

//returns -180 <> +180 degree
func (v Vector2Df) getAngleBehindTarget (t *Vector2Df, rotation float32) float64 {
	xDiff := float64(v.X - t.X)
	yDiff := float64(v.Y - t.Y)
		
	ca := math.Cos(float64(rotation) * 0.01745329251); //0.01745329251
	sa := math.Sin(float64(rotation) * 0.01745329251);
	
	x := ca * xDiff - sa * yDiff
	y := sa * xDiff + ca * yDiff

	res := math.Atan2(float64(x), float64(y)) * 57.2957795131;

	return res
}

func (v Vector2Df) isBehind(t *Vector2Df, rotation float32) bool {
	if v.distance(t) < 1 {
		return false
	}
	absAngle := math.Abs(v.getAngleBehindTarget(t, rotation))

	if(absAngle > 120){
		return true
	} 
	return false
}

func (v Vector2Df) isFacedBy(t *Vector2Df, rotation float32) bool {
	if v.distance(t) < 1 {
		return true
	}
	absAngle := math.Abs(v.getAngleBehindTarget(t, rotation))

	if(absAngle < 60){
		return true
	} 
	return false
}

func Intersection (p0, p1, p2, p3 *Vector2Df) (bool, Vector2Df) {
	s1_x := p1.X - p0.X
	s1_y := p1.Y - p0.Y
	s2_x := p3.X - p2.X
	s2_y := p3.Y - p2.Y

	s := (-s1_y * (p0.X - p2.X) + s1_x * (p0.Y - p2.Y)) / (-s2_x * s1_y + s1_x * s2_y);
	t := ( s2_x * (p0.Y - p2.Y) - s2_y * (p0.X - p2.X)) / (-s2_x * s1_y + s1_x * s2_y);
	
	if s >= 0 && s <= 1 && t >= 0 && t <= 1 {
        return true, Vector2Df {X: p0.X + (t * s1_x), Y: p0.Y + (t * s1_y)};
    }

    return false, Vector2Df {}; 
}


func IntersectingBorders (start *Vector2Df, target *Vector2Df, m *graphmap.Map) (bool) {
	for _, b := range m.Borders {
		hasIntersection, _ := Intersection(start, target, &Vector2Df{X: b.A.X, Y:b.A.Y}, &Vector2Df{X: b.B.X, Y:b.B.Y})
		if hasIntersection {
			return true
		}
	}
	return false
}