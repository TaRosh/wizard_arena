package spell

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Spell interface {
	Update()
	Draw(screen *ebiten.Image)
	Dead() bool
}
