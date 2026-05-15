package game

import "github.com/TaRosh/wizard_arena/mathx"

type (
	Input uint8
)

type InputState struct {
	MousePos mathx.Vec2

	Buttons Input

	Movement mathx.Vec2
}

type MouseStroke struct {
	Points []mathx.Vec2
}

const (
	InputUp Input = 1 << iota
	InputDown
	InputLeft
	InputRight
	InputLeftPressed
	InputLeftJustPressed
	InputLeftJustReleased
)

func (i Input) IsLeftPressed() bool {
	return i&InputLeftPressed != 0
}

func (i Input) IsLeftJustPressed() bool {
	return i&InputLeftJustPressed != 0
}

func (i Input) IsLeftJustReleased() bool {
	return i&InputLeftJustReleased != 0
}

func (i Input) IsUp() bool {
	return i&InputUp != 0
}

func (i Input) IsDown() bool {
	return i&InputDown != 0
}

func (i Input) IsLeft() bool {
	return i&InputLeft != 0
}

func (i Input) IsRight() bool {
	return i&InputRight != 0
}
