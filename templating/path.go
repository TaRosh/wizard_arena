package templating

import (
	"image"
	"os"

	"github.com/TaRosh/wizard_arena/mathx"
)

func GetTemplate(name string) ([]mathx.Vec2, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}
	path := travers(img)
	out := make([]mathx.Vec2, len(path))
	for i := range path {
		out[i] = mathx.Vec2{X: path[i].Pos.X, Y: path[i].Pos.Y}
	}
	return out, nil
}
