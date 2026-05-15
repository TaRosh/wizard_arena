package spell

import (
	"github.com/TaRosh/wizard_arena/gesture"
	"github.com/TaRosh/wizard_arena/mathx"
)

var GestureFireball = GestureTemplate{
	Name: "Fireball",
	// square
	// [-1,-1]-----[1,-1]
	// |             |
	// |             |
	// [-1,1]----[1,1]
	Points: gesture.Normalize(gesture.Resample([]mathx.Vec2{
		// top
		{-1, -1},
		{-0.5, -1},
		{0, -1},
		{0.5, -1},
		{1, -1},
		// right
		{1, -0.5},
		{1, 0},
		{1, 0.5},
		{1, 1},
		// bottom
		{0.5, 1},
		{0, 1},
		{-0.5, 1},
		{-1, 1},
		// left
		{-1, 0.5},
		{-1, 0},
		{-1, -0.5},
		{-1, -1},
	}, 32)),
}
