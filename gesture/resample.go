package gesture

import (
	"github.com/TaRosh/wizard_arena/mathx"
)

func Resample(points []mathx.Vec2, pointsToResample int) []mathx.Vec2 {
	if len(points) == 0 {
		return nil
	}
	if pointsToResample == 1 {
		return []mathx.Vec2{points[0]}
	}
	// len(points) / pointsToResample - 1
	// -1 becouse we divide by space between points
	// A.. B ... C.. D
	// 4 point 3 space between
	spacing := PathLength(points) / (float64(pointsToResample) - 1)
	// spacing = px gaps between
	// A----P----P----P----P----B
	//   ^ spacing px
	out := []mathx.Vec2{points[0]}
	var distance float64
	for i := 1; i < len(points); i++ {
		prev := points[i-1]
		next := points[i]
		lineSegment := prev.Distance(next)
		for distance+lineSegment >= spacing {
			partSegment := spacing - distance
			t := partSegment / lineSegment
			newPoint := mathx.Lerp(prev, next, t)
			out = append(out, newPoint)
			prev = newPoint
			lineSegment -= partSegment
			distance = 0
		}
		distance += lineSegment
	}
	if len(out) < pointsToResample {
		out = append(out, points[len(points)-1])
	}
	return out
}

func PathLength(points []mathx.Vec2) float64 {
	if len(points) < 2 {
		return 0
	}
	var length float64
	for i := range len(points) - 1 {
		distance := points[i].Distance(points[i+1])
		length += distance
	}
	return length
}
