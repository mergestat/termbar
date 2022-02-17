package termbar

import (
	"fmt"
	"io"
	"math"
)

var sparks = []string{"▁", "▂", "▃", "▅", "▇"}

type Sparkline struct {
	Chart
	*baseChart
	Bars Bars
}

// SparklineOption is an option for a sparkline
type SparklineOption interface {
	setSparklineOption() func(*Sparkline)
}

// NewSparkline creates a Sparkline chart from a set of bars
func NewSparkline(bars Bars, options ...SparklineOption) *Sparkline {
	c := &Sparkline{
		Bars:      bars,
		baseChart: &baseChart{},
	}

	for _, opt := range options {
		switch v := opt.(type) {
		case ChartOption:
			v.set()(c.baseChart)
		case SparklineOption:
			v.setSparklineOption()(c)
		}
	}
	return c
}

func (c *Sparkline) Print(w io.Writer) {
	maxVal := c.Bars.MaxValue()
	if c.maxVal != 0 {
		maxVal = c.maxVal
	}

	for _, bar := range c.Bars {
		v := math.Max(math.Round((float64(len(sparks))-1)*((bar.Value-1)/(maxVal-1))), 0)
		fmt.Fprintf(w, "%s", sparks[int(v)])
	}
	fmt.Fprintln(w)
}
