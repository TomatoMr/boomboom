package render

import (
	"github.com/TomatoMr/boomboom/components"
	"testing"
)

const ROUND = 1

func TestRender(t *testing.T) {
	Render(ROUND, &components.Point{
		X:         0,
		Y:         0,
		LastPoint: &components.Point{
			X:         1,
			Y:         0,
			LastPoint: &components.Point{
				X:         2,
				Y:         0,
				LastPoint: &components.Point{
					X:         2,
					Y:         1,
					LastPoint: nil,
				},
			},
		},
	})
}
