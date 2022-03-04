package termbar

import (
	"io"

	"golang.org/x/term"
)

type Chart interface {
	Print(io.Writer)
}

type baseChart struct {
	maxWidth int
	maxVal   float64
}

func newDefaultBaseChart() *baseChart {
	w, _, _ := term.GetSize(0)
	if w < 0 {
		w = 40
	}

	return &baseChart{
		maxWidth: w,
	}
}

type ChartOption interface {
	set() func(*baseChart)
	HorizontalOption
	SparklineOption
}

type setMaxWidth struct{ maxWidth int }

func (o setMaxWidth) set() func(c *baseChart) {
	return func(c *baseChart) {
		c.maxWidth = o.maxWidth
	}
}
func (o setMaxWidth) setHorizontalOption() func(*Horizontal) { return nil }
func (o setMaxWidth) setSparklineOption() func(*Sparkline)   { return nil }

// WithMaxWidth sets the maximum character width of the chart
func WithMaxWidth(w int) ChartOption {
	return &setMaxWidth{maxWidth: w}
}

type setMaxVal struct{ maxVal float64 }

func (o setMaxVal) set() func(*baseChart) {
	return func(c *baseChart) {
		c.maxVal = o.maxVal
	}
}

func (o setMaxVal) setHorizontalOption() func(*Horizontal) { return nil }
func (o setMaxVal) setSparklineOption() func(*Sparkline)   { return nil }

// WithMaxVal sets the max value bars will be made relative to.
// By default, the bar with the highest value will take up the full width
// and all other bars will be sized relative to it.
// However, for charts where an absolute max value is desired, such as a percentage,
// you can use this option to make all bars sized relative to a specific value (i.e. 100.0 for % bars)
func WithMaxVal(m float64) ChartOption {
	return &setMaxVal{maxVal: m}
}

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
