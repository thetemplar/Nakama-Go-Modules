package main

import (
	"math/rand"
)

func randomInt(min, max int32) int32 {
    return min + rand.Int31n(max-min)
}
func randomFloat(min, max float32) float32 {
    difference := max - min
    return rand.Float32() * difference + min
}
func randomFloatInt(min, max int32) float32 {
    return randomFloat(float32(min), float32(max))
}
func randomPercentage() float32 {
    return rand.Float32()
}