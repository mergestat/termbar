package termbar

import (
	"fmt"
	"io"
	"math"
	"strings"

	"golang.org/x/term"
)

// HorizontalOption is an option for a horizontal bar graph
type HorizontalOption func(*horizontal)

type horizontal struct {
	maxWidth       int
	chars          string
	bars           Bars
	labelSeparator string
	valueFormatter func(bar Bar) string
}

// WithMaxWidth sets the maximum width
func WithMaxWidth(w int) HorizontalOption {
	return func(c *horizontal) {
		c.maxWidth = w
	}
}

// WithBarChars sets the characters to use in the bar
func WithBarChars(s string) HorizontalOption {
	return func(c *horizontal) {
		c.chars = s
	}
}

// WithLabelSeparator sets the separator to use between the label and bar
func WithLabelSeparator(s string) HorizontalOption {
	return func(c *horizontal) {
		c.labelSeparator = s
	}
}

// WithValueFormatter sets a formatter function for displaying the value after the bar
func WithValueFormatter(f func(Bar) string) HorizontalOption {
	return func(c *horizontal) {
		c.valueFormatter = f
	}
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
func NewHorizontal(values Bars, options ...HorizontalOption) *horizontal {
	w, _, _ := term.GetSize(0)
	if w < 0 {
		w = 40
	}

	c := &horizontal{
		maxWidth:       w,
		chars:          "▇",
		bars:           values,
		labelSeparator: ": ",
		valueFormatter: func(bar Bar) string { return fmt.Sprintf(" %0.f", bar.Value) },
	}

	for _, opt := range options {
		opt(c)
	}

	return c
}

// Print the bar chart to a writer
func (c *horizontal) Print(w io.Writer) {
	maxVal := c.bars.MaxValue()
	maxLab := c.bars.MaxLabelWidth()

	for _, bar := range c.bars {
		s := c.valueFormatter(bar)
		v := int(math.RoundToEven((bar.Value/maxVal)*float64(c.maxWidth))) - maxLab - len(c.chars) - len(s)
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
