package vector

import (
	"fmt"
	"math"
)

var (
	Cardinal2North V2 = V2{X: 0, Y: -1}
	Cardinal2South V2 = V2{X: 0, Y: 1}
	Cardinal2East  V2 = V2{X: 1, Y: 0}
	Cardinal2West  V2 = V2{X: -1, Y: 0}
)

type V2 struct{ X, Y int }

func (a V2) String() string {
	return fmt.Sprintf("(%d, %d)", a.X, a.Y)
}

func (a V2) OrthoDistance(b V2) V2 {
	dx := a.X - b.X
	dy := a.Y - b.Y
	return V2{dx, dy}
}

// Not sure if "orthogonal slope" is a real term but the [V2] returned here is:
// {X: run, Y: rise}
func (a V2) OrthoSlope(b V2) V2 {
	rise := b.Y - a.Y
	run := b.X - a.X
	return V2{X: run, Y: rise}
}

func (a V2) Distance(b V2) float64 {
	dx := a.X - b.X
	dy := a.Y - b.Y
	return math.Sqrt(float64(dx*dx) + float64(dy*dy))
}

func (a V2) Add(b V2) V2 {
	return V2{a.X + b.X, a.Y + b.Y}
}

func (a V2) Sub(b V2) V2 {
	return V2{a.X - b.X, a.Y - b.Y}
}

func (a V2) Neg() V2 {
	return V2{-a.X, -a.Y}
}

func (a V2) ToDir() string {
	switch {
	case a.X == 0 && a.Y == -1:
		return "up"
	case a.X == 0 && a.Y == 1:
		return "down"
	case a.X == -1 && a.Y == 0:
		return "left"
	case a.X == 1 && a.Y == 0:
		return "right"
	default:
		return "unknown"
	}
}

func (a V2) RotateRight() V2 {
	cardinalTurnMap := map[V2]V2{
		Cardinal2North: Cardinal2East,
		Cardinal2East:  Cardinal2South,
		Cardinal2South: Cardinal2West,
		Cardinal2West:  Cardinal2North,
	}
	return cardinalTurnMap[a]
}

func (a V2) Translate(n int, dir V2) V2 {
	return V2{a.X + n*dir.X, a.Y + n*dir.Y}
}
