package game

import (
	"fmt"
	"image/color"
	"math"

	"github.com/TaRosh/wizard_arena/gesture"
	"github.com/TaRosh/wizard_arena/mathx"
	"github.com/TaRosh/wizard_arena/spell"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const MinStrokePointDistance = 8

type Game struct {
	Width  int
	Height int

	Input InputState

	// raw points
	Stroke MouseStroke
	// resampled points
	ResampledStroke []mathx.Vec2

	Spells      []spell.Spell
	Normalize   []mathx.Vec2
	Comparison  float64
	CastedSpell string
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, fmt.Sprintf("Comparison is %f %q", g.Comparison, g.CastedSpell))
	// draw raw points
	for _, p := range g.Stroke.Points {
		vector.StrokeCircle(screen, float32(p.X), float32(p.Y), 2, 2, color.White, false)
	}
	// draw resampled points
	for _, point := range g.ResampledStroke {
		vector.StrokeCircle(screen, float32(point.X), float32(point.Y), 4, 4, color.RGBA{0xff, 0, 0, 0xff}, false)
	}
	for i := 0; i < len(g.ResampledStroke)-1; i++ {
		a := g.ResampledStroke[i]
		b := g.ResampledStroke[i+1]
		vector.StrokeLine(screen, float32(a.X), float32(a.Y), float32(b.X), float32(b.Y), 2, color.RGBA{0xff, 0, 0, 128}, false)
	}
}

// Layout implements ebiten.Game.
func (g *Game) Layout(outsideWidth int, outsideHeight int) (screenWidth int, screenHeight int) {
	return g.Width, g.Height
}

func (g *Game) readInput() {
	/*

		type InputState struct {
			MousePos mathx.Vec2

			Buttons Input

			Movement mathx.Vec2
		}
	*/
	var inputState InputState
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		inputState.Buttons |= InputUp
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		inputState.Buttons |= InputDown
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		inputState.Buttons |= InputLeft
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		inputState.Buttons |= InputRight
	}

	mx, my := ebiten.CursorPosition()
	inputState.MousePos = mathx.Vec2{float64(mx), float64(my)}
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		inputState.Buttons |= InputLeftPressed
	}
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		inputState.Buttons |= InputLeftJustPressed
	}
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		inputState.Buttons |= InputLeftJustReleased
	}
	g.Input = inputState
}

func (g *Game) updateStroke() {
	if g.Input.Buttons.IsLeftPressed() {
		if len(g.Stroke.Points) == 0 {
			g.Stroke.Points = append(g.Stroke.Points, mathx.Vec2{g.Input.MousePos.X, g.Input.MousePos.Y})
		} else {
			last := g.Stroke.Points[len(g.Stroke.Points)-1]
			current := g.Input.MousePos
			if last.Distance(current) > MinStrokePointDistance {
				g.Stroke.Points = append(g.Stroke.Points, mathx.Vec2{g.Input.MousePos.X, g.Input.MousePos.Y})
			}
		}
	}
	if g.Input.Buttons.IsLeftJustReleased() && len(g.Stroke.Points) > 1 {
		g.ResampledStroke = gesture.Resample(g.Stroke.Points, 32)
		g.Normalize = gesture.Normalize(g.ResampledStroke)
		// iterate over all spells
		// pick closest
		guess := math.MaxFloat64
		for _, spell := range spell.Spells {
			g.Comparison = gesture.CompareGestures(g.Normalize, spell.Points)
			if g.Comparison < guess {
				guess = g.Comparison
				g.CastedSpell = spell.Name
			}
		}
		// g.Comparison = gesture.CompareGestures(g.Normalize, spell.GestureCircle.Points)
		// guess := g.Comparison
		// g.CastedSpell = spell.GestureCircle.Name
		// g.Comparison = gesture.CompareGestures(g.Normalize, spell.GestureY.Points)
		// if g.Comparison < guess {
		// 	guess = g.Comparison
		// 	g.CastedSpell = spell.GestureY.Name
		// }
		// g.Comparison = gesture.CompareGestures(g.Normalize, spell.GestureFireball.Points)
		// if g.Comparison < guess {
		// 	guess = g.Comparison
		// 	g.CastedSpell = spell.GestureFireball.Name
		// }
		// g.Comparison = gesture.CompareGestures(g.Normalize, spell.GestureWind.Points)
		// if g.Comparison < guess {
		// 	guess = g.Comparison
		// 	g.CastedSpell = spell.GestureWind.Name
		// }
		g.Stroke.Points = g.Stroke.Points[:0]
		g.ResampledStroke = g.ResampledStroke[:0]
		return
	}
}

// Update implements ebiten.Game.
func (g *Game) Update() error {
	// Read input
	g.readInput()
	// Update Simulation
	g.updateStroke()
	// g.updateSpells
	// Resolve Collisions
	// Cleanup

	return nil
}

func NewGame(w, h int) *Game {
	g := Game{
		Width:  w,
		Height: h,
	}

	return &g
}

func (g *Game) Run() error {
	return ebiten.RunGame(g)
}
