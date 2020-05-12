package rounds

import (
	"math/rand"
	"sync"
)

type RoundParams struct {
	Height       int
	Width        int
	BoomNum      int
	BoomPosition [][]int
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
				Height:  10,
				Width:   10,
				BoomNum: 1,
			}
			x := rand.Intn(roundParams.Width-3) + 2
			y := rand.Intn(roundParams.Height-3) + 2
			bp := [][]int{{x, y}}
			roundParams.BoomPosition = bp
		}
	})

	return roundParams
}
