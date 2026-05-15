package mathx

import "math"

type Vec2 struct {
	X float64
	Y float64
}

func (v Vec2) Distance(other Vec2) float64 {
	a := (v.X - other.X)
	b := (v.Y - other.Y)
	return math.Sqrt(a*a + b*b)
}

func Lerp(a, b Vec2, t float64) Vec2 {
	return a.Add(b.Sub(a).Mulf(t))
}

func Map(x, fromMin, fromMax, toMin, toMax float64) float64 {
	return (x-fromMin)*(toMax-toMin)/(fromMax-fromMin) + toMin
}

func (v Vec2) Mulf(value float64) Vec2 {
	return Vec2{
		X: v.X * value,
		Y: v.Y * value,
	}
}

func (v Vec2) Add(other Vec2) Vec2 {
	return Vec2{
		X: v.X + other.X,
		Y: v.Y + other.Y,
	}
}

func (v Vec2) Sub(other Vec2) Vec2 {
	return v.Add(Vec2{
		X: -other.X,
		Y: -other.Y,
	})
}
