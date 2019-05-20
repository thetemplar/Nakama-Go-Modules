package main

import (
	"math"
	"github.com/bradfitz/slice"
)

func (v PublicMatchState_Vector2Df) rotate(degrees float32) PublicMatchState_Vector2Df {
	ca := float32(math.Cos(float64(360 - degrees) * 0.01745329251)); //0.01745329251
	sa := float32(math.Sin(float64(360 - degrees) * 0.01745329251));

	vec := PublicMatchState_Vector2Df {
		X: ca * v.X - sa * v.Y,
		Y: sa * v.X + ca * v.Y,
	}
	return vec
}

func (v PublicMatchState_Vector2Df) distance(t PublicMatchState_Vector2Df) float32 {
	return float32(math.Sqrt(math.Pow(float64(t.X - v.X), 2) + math.Pow(float64(t.Y - v.Y), 2)))
}

type NextTriangles struct{
	Distance float64
	Index 	 int
}

func (v PublicMatchState_Vector2Df) getNearest(m *Map, length int) []NextTriangles {
	res := make([]NextTriangles, len(m.Triangles))
	i := 0
	for _, triangle := range m.Triangles {
		res[i].Distance = float64(triangle.W.distance(v))
		res[i].Index = i
		i++
	}
	slice.Sort(res[:], func(i, j int) bool {
		return res[i].Distance < res[j].Distance
	})
	return res[:length]
}