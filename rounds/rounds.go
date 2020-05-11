package rounds

type Canvas struct {
	Height int
	Width int
}

func GetCanvasParams(round int) Canvas {
	params := Canvas{
		Height: 10,
		Width:  10,
	}
	switch round {
	case 1:
		params = Canvas{
			Height: 10,
			Width:  10,
		}
	}
	return params
}
