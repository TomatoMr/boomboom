package keyboard

import (
	"fmt"
	"github.com/TomatoMr/boomboom/components"
	"github.com/TomatoMr/boomboom/render"
	"github.com/TomatoMr/boomboom/rounds"
)
import "github.com/nsf/termbox-go"

func KeyEvent(round int) {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	fmt.Println("press enter to start the game.")
	rp := rounds.GetRoundParams(round)
	startPoint, isStop := components.NewPoint(rp, nil, "")
	if isStop {
		fmt.Println("you cannot go to there.")
	}
	currentPoint := startPoint
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc, termbox.KeyCtrlQ:
				fmt.Println("bye")
				break
			case termbox.KeyArrowUp:
				if checkTrap(round, currentPoint, "up") {
					fmt.Println("you cannot go to a trap.")
					continue
				}
				currentPoint, isStop = components.NewPoint(rp, currentPoint, "up")
				if isStop {
					fmt.Println("you cannot go to there.")
				}
				render.Render(round, currentPoint)
			case termbox.KeyArrowDown:
				if checkTrap(round, currentPoint, "down") {
					fmt.Println("you cannot go to a trap.")
					continue
				}
				currentPoint, isStop = components.NewPoint(rp, currentPoint, "down")
				if isStop {
					fmt.Println("you cannot go to there.")
				}
				render.Render(round, currentPoint)
			case termbox.KeyArrowLeft:
				if checkTrap(round, currentPoint, "left") {
					fmt.Println("you cannot go to a trap.")
					continue
				}
				currentPoint, isStop = components.NewPoint(rp, currentPoint, "left")
				if isStop {
					fmt.Println("you cannot go to there.")
				}
				render.Render(round, currentPoint)
			case termbox.KeyArrowRight:
				if checkTrap(round, currentPoint, "right") {
					fmt.Println("you cannot go to a trap.")
					continue
				}
				currentPoint, isStop = components.NewPoint(rp, currentPoint, "right")
				if isStop {
					fmt.Println("you cannot go to there.")
				}
				render.Render(round, currentPoint)
			case termbox.KeyCtrlA:
				fmt.Println("give u a road.")
				render.BestPath(rp, startPoint)
				pm := render.GetPM()
				length := 0
				for l, _ := range pm.Points {
					length = l
					break
				}
				for l, _ := range pm.Points {
					if length > l {
						length = l
					}
				}
				render.Render(round, pm.Points[length][0])
			case termbox.KeyEnter:
				currentPoint = startPoint
				render.Render(round, currentPoint)
			default:

			}
		}
	}
}

func checkTrap(round int, point *components.Point, direction string) bool {
	rp := rounds.GetRoundParams(round)
	x := point.X
	y := point.Y
	switch direction {
	case "up":
		y -= 1
	case "down":
		y += 1
	case "left":
		x -= 1
	case "right":
		x += 1
	}
	for _, v := range rp.TrapPosition {
		if x == v[0] && y == v[1] {
			return true
		}
	}
	return false
}
