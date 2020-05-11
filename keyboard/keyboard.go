package keyboard

import (
	"fmt"
	"github.com/TomatoMr/boomboom/components"
	"github.com/TomatoMr/boomboom/render"
)
import "github.com/nsf/termbox-go"

func KeyEvent(round int) {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	fmt.Println("press enter to start the game.")
	startPoint := components.NewPoint(round, nil, "")
	currentPoint := startPoint
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc, termbox.KeyCtrlQ:
				fmt.Println("bye")
				break
			case termbox.KeyArrowUp:
				currentPoint = components.NewPoint(round, currentPoint, "up")
				render.Render(round, currentPoint)
			case termbox.KeyArrowDown:
				currentPoint = components.NewPoint(round, currentPoint, "down")
				render.Render(round, currentPoint)
			case termbox.KeyArrowLeft:
				currentPoint = components.NewPoint(round, currentPoint, "left")
				render.Render(round, currentPoint)
			case termbox.KeyArrowRight:
				currentPoint = components.NewPoint(round, currentPoint, "right")
				render.Render(round, currentPoint)
			case termbox.KeyEnter:
				render.Render(round, startPoint)
			default:

			}
		}
	}

}

