package vector

import "math"

type V2 struct{ X, Y int }

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
