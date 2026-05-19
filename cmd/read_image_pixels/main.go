package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/png"
	"os"

	"github.com/quasilyte/gmath"
)

type Point struct {
	Visited bool
	Pos     gmath.Vec
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

func main() {
	f, err := os.Open("a.png")
	if err != nil {
		panic(err)
	}
	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}
	leftUp := img.Bounds().Min
	rightBottom := img.Bounds().Max

	width := rightBottom.X - leftUp.X
	height := rightBottom.Y - leftUp.Y
	points := make([][]Point, width)
	for i := range width {
		points[i] = make([]Point, height)
	}
	// empyt matrix height * width

	var startingPoint *Point

	for y := range height {
		for x := range width {
			points[x][y] = Point{Visited: false, Pos: gmath.Vec{X: float64(x), Y: float64(y)}}
			r1, g1, b1, _ := img.At(x, y).RGBA()
			r2, g2, b2, _ := color.Black.RGBA()
			if r1 == r2 && g1 == g2 && b1 == b2 && startingPoint == nil {
				startingPoint = &points[x][y]
			}
		}
	}
	for x := range width {
		for y := range height {
			fmt.Println("Pos", points[x][y])
		}
	}
	// points created
	fmt.Println("Starting point", startingPoint)

	// 1. starting point
	path := make([]Point, 0)
	currentPoint := points[int(startingPoint.Pos.X)][int(startingPoint.Pos.Y)]
	previousPoint := currentPoint
	dirs := []struct{ dx, dy int }{
		{-1, -1},
		{0, -1},
		{1, -1},
		{-1, 0},
		{1, 0},
		{-1, 1},
		{0, 1},
		{1, 1},
	}
	for {
		fmt.Println("Current", currentPoint)
		path = append(path, currentPoint)
		points[int(currentPoint.Pos.X)][int(currentPoint.Pos.Y)].Visited = true
		// find neighbors !visited && black
		//   w   b    w
		//[0,0][1,0][2,0]
		//   w   x    w
		// [0,1][1,1][1,2]
		//   w   b    w
		// [0,2][1,2][2,2]

		neighbors := make([]Point, 0)
		x, y := currentPoint.Coord()
		for _, d := range dirs {
			nx, ny := x+d.dx, y+d.dy
			if nx >= 0 && nx < width &&
				ny >= 0 && ny < height {
				if IsBlack(img.At(nx, ny)) &&
					!points[nx][ny].Visited {
					neighbors = append(neighbors, points[nx][ny])
				}
			}
		}
		if len(neighbors) == 0 {
			break
		}
		var dir float64
		var next int
		prevDirection := currentPoint.Pos.Sub(previousPoint.Pos)
		for n := range neighbors {
			dirGuess := neighbors[n].Pos.Sub(currentPoint.Pos)
			dot := prevDirection.Dot(dirGuess)
			if dot > dir {
				dir = dot
				next = n
			}
		}
		currentPoint = neighbors[next]

		// if no neightbors -> end
		// if meltiple pick closest (dot product ) to direction
		// set current to chosen neighbor
	}
	fmt.Println("PATH", path)
}
