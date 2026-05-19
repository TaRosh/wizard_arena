package spell

import (
	"github.com/TaRosh/wizard_arena/gesture"
	"github.com/TaRosh/wizard_arena/mathx"
	"github.com/TaRosh/wizard_arena/templating"
)

// image name
func GetTemplate(name string, filename string) GestureTemplate {
	gestureTemplate := GestureTemplate{
		Name:   name,
		Points: []mathx.Vec2{},
	}
	path, err := templating.GetTemplate(filename)
	// TODO: think about error handling
	if err != nil {
		panic(err)
	}
	gestureTemplate.Points = gesture.Normalize(gesture.Resample(path, 32))
	return gestureTemplate
}
