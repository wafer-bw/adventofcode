package vector

import (
	"fmt"
	"math"
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
