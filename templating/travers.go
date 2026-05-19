package templating

import (
	"image"
	"image/color"

	"github.com/quasilyte/gmath"
)

func travers(img image.Image) []Point {
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
			points[x][y] = Point{Visits: 0, MaxVisits: 2, Pos: gmath.Vec{X: float64(x), Y: float64(y)}}
			r1, g1, b1, _ := img.At(x, y).RGBA()
			r2, g2, b2, _ := color.Black.RGBA()
			if r1 == r2 && g1 == g2 && b1 == b2 && startingPoint == nil {
				startingPoint = &points[x][y]
			}
		}
	}
	tree := Tree{
		Width:  width,
		Height: height,
		img:    img,
		points: points,
	}
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
	for x := range width {
		for y := range height {
			if IsBlack(img.At(x, y)) {
				degree := len(tree.getNeighbors(x, y, dirs))
				switch degree {
				case 1, 2:
					tree.points[x][y].MaxVisits = 1
				default:
					tree.points[x][y].MaxVisits = 2
				}
				// count neighbors
			}
		}
	}
	// points created

	// 1. starting point
	path := make([]Point, 0)
	currentPoint := points[int(startingPoint.Pos.X)][int(startingPoint.Pos.Y)]
	previousPoint := currentPoint
	stack := []Point{currentPoint}
	for len(stack) > 0 {
		currentPoint = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		path = append(path, currentPoint)
		points[int(currentPoint.Pos.X)][int(currentPoint.Pos.Y)].Visits++
		// find neighbors !visited && black
		//   w   b    w
		//[0,0][1,0][2,0]
		//   w   x    w
		// [0,1][1,1][1,2]
		//   w   b    w
		// [0,2][1,2][2,2]

		neighbors := make([]Point, 0)
		x, y := currentPoint.Coord()
		neighbors = append(neighbors, tree.getNeighbors(x, y, dirs)...)

		if len(neighbors) <= 0 {
			continue
		}

		var dir float64
		var next int
		prevDirection := currentPoint.Pos.Sub(previousPoint.Pos).Normalized()
		// pick best
		for n := range neighbors {
			dirGuess := neighbors[n].Pos.Sub(currentPoint.Pos).Normalized()
			dot := prevDirection.Dot(dirGuess)
			if dot > dir && neighbors[n].Visits == 0 {
				dir = dot
				next = n
			}
		}
		// set prev
		previousPoint = currentPoint
		// add new point to travers
		stack = append(stack, neighbors[next])

		// if no neightbors -> end
		// if meltiple pick closest (dot product ) to direction
		// set current to chosen neighbor
	}
	return path
}
