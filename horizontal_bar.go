package termbar

import (
	"fmt"
	"io"
	"math"
	"strings"
)

// HorizontalOption is an option for a horizontal bar graph
type HorizontalOption interface {
	setHorizontalOption() func(*Horizontal)
}

// Horizontal is a horizontal bar chart
type Horizontal struct {
	Chart
	*baseChart
	chars          string
	Bars           Bars
	labelSeparator string
	valueFormatter func(bar Bar) string
}

type setHorizontalBarChars struct{ chars string }

func (o setHorizontalBarChars) setHorizontalOption() func(*Horizontal) {
	return func(c *Horizontal) {
		c.chars = o.chars
	}
}

// WithBarChars sets the characters to use in the bar
func WithBarChars(s string) HorizontalOption {
	return &setHorizontalBarChars{chars: s}
}

type setHorizontalLabelSeparator struct{ labelSeparator string }

func (o setHorizontalLabelSeparator) setHorizontalOption() func(*Horizontal) {
	return func(c *Horizontal) {
		c.labelSeparator = o.labelSeparator
	}
}

// WithLabelSeparator sets the separator to use between the label and bar
func WithLabelSeparator(s string) HorizontalOption {
	return &setHorizontalLabelSeparator{labelSeparator: s}
}

type setValueFormatter struct{ valueFormatter func(Bar) string }

func (o setValueFormatter) setHorizontalOption() func(*Horizontal) {
	return func(c *Horizontal) {
		c.valueFormatter = o.valueFormatter
	}
}

// WithValueFormatter sets a formatter function for displaying the value after the bar
func WithValueFormatter(f func(Bar) string) HorizontalOption {
	return &setValueFormatter{valueFormatter: f}
}

// NewHorizontal returns a new horizontal bar chart.
// By default, it will produce a set of bars that look like:
//
// a: ▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇ 50
//
// b: ▇▇▇▇▇▇▇▇▇ 25
//
// c: ▇▇▇ 15
//
// d: ▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇ 78
func NewHorizontal(values Bars, options ...HorizontalOption) *Horizontal {
	// default values
	// TODO(patrickdevivo) move this out into a central location?
	c := &Horizontal{
		baseChart:      newDefaultBaseChart(),
		chars:          "▇",
		Bars:           values,
		labelSeparator: ": ",
		valueFormatter: func(bar Bar) string { return fmt.Sprintf(" %0.f", bar.Value) },
	}

	for _, opt := range options {
		switch v := opt.(type) {
		case ChartOption:
			v.set()(c.baseChart)
		case HorizontalOption:
			v.setHorizontalOption()(c)
		}
	}

	return c
}

// Print the bar chart to a writer
func (c *Horizontal) Print(w io.Writer) {
	maxVal := c.Bars.MaxValue()
	if c.maxVal != 0 {
		maxVal = c.maxVal
	}
	maxLab := c.Bars.MaxLabelWidth()

	for _, bar := range c.Bars {
		s := c.valueFormatter(bar)
		v := int(math.Round((bar.Value/maxVal)*float64(c.maxWidth))) - maxLab - len(c.chars) - len(s)
		if v < 1 {
			v = 1
		}
		chars := c.chars
		if bar.Chars != "" {
			chars = bar.Chars
		}
		b := strings.Repeat(chars, v)

		fmt.Fprintf(w, "%*s%s%s%s\n", maxLab, bar.Label, c.labelSeparator, b, s)
	}
}
