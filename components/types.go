package components

import "github.com/TomatoMr/boomboom/rounds"

type Point struct {
	X         int
	Y         int
	LastPoint *Point
}

func NewPoint(round int, lastPoint *Point, direction string) *Point {
	if lastPoint == nil {
		return &Point{
			X:         0,
			Y:         0,
			LastPoint: nil,
		}
	}
	roundParams := rounds.GetRoundParams(round)
	x := lastPoint.X
	y := lastPoint.Y
	switch direction {
	case "up":
		if lastPoint.Y > 0 {
			y -= 1
		}
	case "down":
		if lastPoint.Y < roundParams.Height-1 {
			y += 1
		}
	case "left":
		if lastPoint.X > 0 {
			x -= 1
		}
	case "right":
		if lastPoint.X < roundParams.Width-1 {
			x += 1
		}
	default:

	}
	if x == lastPoint.X && y == lastPoint.Y {
		return lastPoint
	} else if lastPoint.LastPoint != nil && x == lastPoint.LastPoint.X && y == lastPoint.LastPoint.Y {
		return lastPoint.LastPoint
	} else {
		return &Point{
			X:         x,
			Y:         y,
			LastPoint: lastPoint,
		}
	}
}
