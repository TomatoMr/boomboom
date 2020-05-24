package rounds

import (
	"math/rand"
	"sync"
	"time"
)

type RoundParams struct {
	Height       int     //画布高度
	Width        int     //画布宽度
	BoomNum      int     //炸弹数量
	BoomPosition [][]int //炸弹位置，坐标数组
	TrapPosition [][]int //陷阱位置，坐标数组
	StartPoint   []int   //起点坐标
	EndPoint     []int   //终点坐标
}

var roundParams RoundParams
var once = &sync.Once{}

func GetRoundParams(round int) RoundParams {
	once.Do(func() {
		roundParams = RoundParams{
			Height:  10,
			Width:   10,
			BoomNum: 1,
			BoomPosition: [][]int{
				{5, 5},
			},
		}
		switch round {
		case 1:
			roundParams = RoundParams{
				Height:     10,
				Width:      10,
				StartPoint: []int{0, 0},
				EndPoint:   []int{9, 9},
			}
			roundParams.BoomNum = minBooms(roundParams.Height, roundParams.Width)
			bp := make([][]int, 0)
			tp := make([][]int, 0)
			rand.Seed(time.Now().UnixNano())
			for i := 0; i < roundParams.BoomNum-1; i++ {
				x := rand.Intn(roundParams.Width-3) + 2
				y := rand.Intn(roundParams.Height-3) + 2
				bp = append(bp, []int{x, y})
				tp = append(tp, []int{x - 1, y}, []int{x + 1, y}, []int{x, y - 1}, []int{x, y + 1})
			}
			roundParams.BoomPosition = bp
			roundParams.TrapPosition = tp
		}
	})

	return roundParams
}

func minBooms(h, w int) int {
	min := h
	if h-w > 0 {
		min = h
	}
	boomNum := min / 3
	if boomNum%3 != 0 {
		boomNum++
	}
	return boomNum
}
