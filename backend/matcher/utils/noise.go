package utils

import (
	"math/rand/v2"
)

func AddNoise(lat, lng float64) (float64, float64) { //nolint:gosec
	noiseLat := rand.Float64()*0.02 - 0.01 //nolint:gosec
	noiseLng := rand.Float64()*0.02 - 0.01 //nolint:gosec
	return lat + noiseLat, lng + noiseLng
}
