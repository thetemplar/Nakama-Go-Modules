package main

import (
	"math/rand"
)

func randomInt(min, max int32) int32 {
    return min + rand.Int31n(max-min)
}
func randomPercentage() float32 {
    return rand.Float32()
}