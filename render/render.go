package render

import (
	"fmt"
	"github.com/TomatoMr/boomboom/components"
	"github.com/TomatoMr/boomboom/rounds"
	"os"
	"os/exec"
	"runtime"
	"sync"
)

func Render(round int, point *components.Point) {
	clearScreen()
	roundParams := rounds.GetRoundParams(round)
	road := analyzePoint(point)
	for i := 0; i < roundParams.Height; i++ {
		for j := 0; j < roundParams.Width; j++ {
			isStep := false
			for _, r := range road {
				if j == r[0] && i == r[1] {
					isStep = true
					break
				}
			}
			isBoom := false
			for _, b := range roundParams.BoomPosition {
				if j == b[0] && i == b[1] {
					isBoom = true
					break
				}
			}
			if isStep {
				fmt.Print("□")
			} else if isBoom {
				fmt.Print("⊙")
			} else {
				fmt.Print("◎")
			}
			if j == roundParams.Width-1 {
				fmt.Print("\n")
			}
		}
	}
	if point.X == roundParams.EndPoint[0] && point.Y == roundParams.EndPoint[1] {
		fmt.Print("bingo")
	}

}

func analyzePoint(point *components.Point) [][]int {
	result := make([][]int, 0)
	if point == nil {
		return result
	}
	return getPoints(point, result)
}

func getPoints(point *components.Point, result [][]int) [][]int {
	if point == nil {
		return nil
	}
	p := []int{point.X, point.Y}
	result = append(result, p)
	if point.LastPoint == nil {
		return result
	} else {
		result = getPoints(point.LastPoint, result)
	}
	return result
}

func clearScreen() {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "linux":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

type PointsMap struct {
	Points map[int][]*components.Point
	mux    *sync.Mutex
}

var pm = &PointsMap{
	Points: make(map[int][]*components.Point, 0),
	mux:    &sync.Mutex{},
}

func BestPath(rp rounds.RoundParams, point *components.Point) {
	if rp.EndPoint[0] != point.X || rp.EndPoint[1] != point.Y {
		pointR, isStopR := components.NewPoint(rp, point, "right")
		if !isStopR {
			BestPath(rp, pointR)
		}
		//若允许路径往后或往上走，则可以去掉注释
		//pointL, isStopL := components.NewPoint(rp, point, "left")
		//if !isStopL {
		//	BestPath(rp, pointL)
		//}
		pointD, isStopD := components.NewPoint(rp, point, "down")
		if !isStopD {
			BestPath(rp, pointD)
		}
		//pointU, isStopU := components.NewPoint(rp, point, "up")
		//if !isStopU {
		//	BestPath(rp, pointU)
		//}
	} else {
		pm.mux.Lock()
		pm.Points[point.Length] = append(pm.Points[point.Length], point)
		pm.mux.Unlock()
	}
}

func GetPM() *PointsMap {
	return pm
}
