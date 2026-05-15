package gesture

import (
	"math"

	"github.com/TaRosh/wizard_arena/mathx"
)

func Normalize(points []mathx.Vec2) []mathx.Vec2 {
	if len(points) == 0 {
		return nil
	}
	// compute centroid
	var sumX, sumY float64
	for _, point := range points {
		sumX += point.X
		sumY += point.Y
	}
	cX := sumX / float64(len(points))
	cY := sumY / float64(len(points))
	// Translate points to origin (0,0)
	// substraction make centroid suppose (4, 3)
	// make it to zero (4, 3) - (4, 3) = (0, 0)
	translated := make([]mathx.Vec2, len(points))
	for idx, point := range points {
		translated[idx] = mathx.Vec2{X: point.X - cX, Y: point.Y - cY}
	}
	// Compute max distance for scaling
	// compute max value of one of the axis
	// then divide to make it in range [-1, 1]
	var maxDist float64
	for _, point := range translated {
		dist := math.Abs(point.X)
		if math.Abs(point.Y) > dist {
			dist = math.Abs(point.Y)
		}
		if dist > maxDist {
			maxDist = dist
		}
	}
	if maxDist == 0 {
		return translated
	}
	// scale: divide value by max value
	// got part of it 0.5 or full 1
	for i, point := range translated {
		translated[i] = mathx.Vec2{X: point.X / maxDist, Y: point.Y / maxDist}
	}
	return translated
}
