package gesture

import (
	"math"

	"github.com/TaRosh/wizard_arena/mathx"
)

func CompareGestures(input, template []mathx.Vec2) float64 {
	if len(input) != len(template) {
		return math.MaxFloat64
	}
	var sum float64
	// calculate distance for each point
	// if a[i] on distance for b[i] is 0
	// so they the same
	for i := range input {
		sum += input[i].Distance(template[i])
	}
	// WHY?
	return sum / float64(len(input))
}
