package utils

import (
	"math/rand"
	"time"
)

func AddNoise(lat, lng float64) (float64, float64) {
	rand.Seed(time.Now().UnixNano())
	noiseLat := rand.Float64()*0.02 - 0.01 // [-0.01, 0.01]
	noiseLng := rand.Float64()*0.02 - 0.01 // [-0.01, 0.01]
	return lat + noiseLat, lng + noiseLng
}
