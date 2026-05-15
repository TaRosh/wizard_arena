package spell

import (
	"github.com/TaRosh/wizard_arena/mathx"
)

type GestureTemplate struct {
	Name string
	// resempled and normalized
	Points []mathx.Vec2
}
