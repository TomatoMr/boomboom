package components

import (
	"github.com/TomatoMr/boomboom/rounds"
)

type Point struct {
	X         int
	Y         int
	LastPoint *Point
	Length    int
}

func NewPoint(rp rounds.RoundParams, lastPoint *Point, direction string) (*Point, bool) {
	if lastPoint == nil {
		return &Point{
			X:         0,
			Y:         0,
			LastPoint: nil,
			Length:    1,
		}, false
	}
	x := lastPoint.X
	y := lastPoint.Y
	switch direction {
	case "up":
		if lastPoint.Y > 0 {
			y -= 1
		} else {
			return lastPoint, true
		}
	case "down":
		if lastPoint.Y < rp.Height-1 {
			y += 1
		} else {
			return lastPoint, true
		}
	case "left":
		if lastPoint.X > 0 {
			x -= 1
		} else {
			return lastPoint, true
		}
	case "right":
		if lastPoint.X < rp.Width-1 {
			x += 1
		} else {
			return lastPoint, true
		}
	default:

	}
	if x == lastPoint.X && y == lastPoint.Y {
		return lastPoint, true
	}

	for _, v := range rp.BoomPosition {
		if x == v[0] && y == v[1] {
			return lastPoint, true
		}
	}
	for _, v := range rp.TrapPosition {
		if x == v[0] && y == v[1] {
			return lastPoint, true
		}
	}

	point := &Point{
		X:         x,
		Y:         y,
		LastPoint: lastPoint,
		Length:    lastPoint.Length + 1,
	}
	if HasDuplicatedPoint(point) {
		return lastPoint, true
	} else {
		return point, false
	}
}

func HasDuplicatedPoint(p *Point) bool {
	_p := p
	x, y := _p.X, _p.Y
	for _p.LastPoint != nil {
		if _p.LastPoint.X == x && _p.LastPoint.Y == y {
			return true
		} else {
			_p = _p.LastPoint
		}
	}
	return false
}
