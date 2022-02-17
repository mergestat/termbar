package termbar_test

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/mergestat/termbar"
)

func TestHorizontalBarSimpleVisual(t *testing.T) {
	c := termbar.NewHorizontal([]termbar.Bar{
		{Label: "a", Value: 50},
		{Label: "b", Value: 25},
		{Label: "b", Value: 15},
	})

	var buf bytes.Buffer
	c.Print(&buf)

	t.Log("\n" + buf.String())
}

func TestHorizontalBarMaxWidth(t *testing.T) {
	maxWidth := 300
	c := termbar.NewHorizontal([]termbar.Bar{
		{Label: "a", Value: 50000},
		{Label: "b", Value: 25000},
		{Label: "b", Value: 15000},
	})

	var buf bytes.Buffer
	c.Print(&buf)

	scanner := bufio.NewScanner(&buf)
	for scanner.Scan() {
		if l := scanner.Text(); len(l) > maxWidth {
			t.Fatalf("line printed wider than maxWidth, len=%d, max=%d", len(l), maxWidth)
		}
	}
}
