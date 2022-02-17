package termbar

type Bar struct {
	Label string
	Value float64
	Chars string
}

type Bars []Bar

// MaxValue returns the maxiumum Value from a list of Bars
func (bars Bars) MaxValue() float64 {
	var max float64
	for _, bar := range bars {
		if v := bar.Value; v > max {
			max = v
		}
	}
	return max
}

// MaxLabelWidth returns the max length of the Label from a list of Bars
func (bars Bars) MaxLabelWidth() int {
	var max int
	for _, bar := range bars {
		if l := len(bar.Label); l > max {
			max = l
		}
	}
	return max
}
