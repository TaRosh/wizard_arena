package templating

import (
	"image"
	"image/color"

	"github.com/quasilyte/gmath"
)

type Point struct {
	Visits    int
	MaxVisits int
	Pos       gmath.Vec
}

func (p Point) Dot(other Point) float64 {
	return p.Pos.Dot(other.Pos)
}

func (p Point) Coord() (x int, y int) {
	return int(p.Pos.X), int(p.Pos.Y)
}

func IsBlack(c color.Color) bool {
	r1, g1, b1, _ := c.RGBA()
	r2, g2, b2, _ := color.Black.RGBA()
	return r1 == r2 && g1 == g2 && b1 == b2
}

type Tree struct {
	Width  int
	Height int
	img    image.Image
	points [][]Point
}

func (t Tree) getNeighbors(x, y int, dirs []struct{ dx, dy int }) []Point {
	var neighbors []Point
	for _, d := range dirs {
		nx, ny := x+d.dx, y+d.dy
		if nx >= 0 && nx < t.Width &&
			ny >= 0 && ny < t.Height {
			if IsBlack(t.img.At(nx, ny)) &&
				t.points[nx][ny].Visits < t.points[nx][ny].MaxVisits {
				neighbors = append(neighbors, t.points[nx][ny])
			}
		}
	}
	return neighbors
}
