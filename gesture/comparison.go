package gesture

import (
	"math"

	"github.com/TaRosh/wizard_arena/mathx"
)

//	func CompareGestures(input, template []mathx.Vec2) float64 {
//		if len(input) != len(template) {
//			return math.MaxFloat64
//		}
//		var sum float64
//		// calculate distance for each point
//		// if a[i] on distance for b[i] is 0
//		// so they the same
//		for i := range input {
//			sum += input[i].Distance(template[i])
//		}
//		// WHY?
//		return sum / float64(len(input))
//	}
func CompareGestures(input, template []mathx.Vec2) float64 {
	chamferScore := 0.0

	for _, iPoint := range input {
		minDist := math.MaxFloat64
		for _, tPoint := range template {
			d := iPoint.Distance(tPoint)
			if d < minDist {
				minDist = d
			}
		}
		chamferScore += minDist
	}

	for _, tPoint := range template {
		minDist := math.MaxFloat64
		for _, iPoint := range input {
			d := tPoint.Distance(iPoint)
			if d < minDist {
				minDist = d
			}
		}
		chamferScore += minDist
	}

	return chamferScore / (float64(len(input) + len(template)))
}

// 3 score
// 1. Chamfer - distance
// 2. Angle between points
// 3 Corner count 3- triangle 4- square, circle
// func CompareGestures(input, template []mathx.Vec2) float64 {
// 	chamferScore := 0.0
// 	for _, iPoint := range input {
// 		minDist := math.MaxFloat64
// 		for _, tPoint := range template {
// 			d := iPoint.Distance(tPoint)
// 			if d < minDist {
// 				minDist = d
// 			}
// 		}
// 		chamferScore += minDist
// 	}
// 	for _, tPoint := range template {
// 		minDist := math.MaxFloat64
// 		for _, iPoint := range input {
// 			d := tPoint.Distance(iPoint)
// 			if d < minDist {
// 				minDist = d
// 			}
// 		}
// 		chamferScore += minDist
// 	}
// 	chamferScore /= float64(len(input) + len(template))
// 	// 2.Angle differences
// 	angleScore := 0.0
// 	for i := 1; i < len(input)-2; i++ {
// 		prevDir := input[i].Sub(input[i-1])
// 		nextDir := input[i+1].Sub(input[i])
// 		angleInput := prevDir.AngleBetween(nextDir)
//
// 		prevDirTemplate := template[i].Sub(template[i-1])
// 		nextDirTemplate := template[i+1].Sub(template[i])
// 		angleTemplate := prevDirTemplate.AngleBetween(nextDirTemplate)
//
// 		angleScore += math.Abs(angleInput - angleTemplate)
//
// 	}
// 	// 3. Corner count
// 	cornersInput := CountCorners(input)
// 	cornersTemplate := CountCorners(template)
// 	cornerScore := math.Abs(float64(cornersInput) - float64(cornersTemplate))
// 	// Weight
// 	wChamfer := 0.6
// 	wAngle := 0.3
// 	wCorner := 0.1
// 	finalScore := chamferScore*wChamfer + angleScore*wAngle + cornerScore*wCorner
// 	return finalScore
// }

// func CountCorners(points []mathx.Vec2) int {
// 	thresholdAngle := 120.0
// 	var count int
// 	for i := 1; i < len(points)-2; i++ {
// 		v1 := points[i].Sub(points[i-1])
// 		v2 := points[i+1].Sub(points[i])
// 		angle := v1.AngleBetween(v2)
// 		if angle < thresholdAngle {
// 			count += 1
// 		}
// 	}
// 	return count
// }
