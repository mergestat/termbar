package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/harmonica"
	"github.com/mergestat/termbar"
)

func main() {
	p := tea.NewProgram(newModel(3))

	if p.Start() != nil {
		fmt.Println("could not start program")
		os.Exit(1)
	}
}

type barWithVelocity struct {
	termbar.Bar
	Velocity float64
}

type model struct {
	chartOpts []termbar.HorizontalOption
	current   []barWithVelocity
	target    termbar.Bars
	spring    harmonica.Spring
}

type frameMsg time.Time
type updateMsg time.Time

func newModel(barCount int) *model {
	current := make([]barWithVelocity, barCount)
	target := make(termbar.Bars, barCount)
	for i := 0; i < barCount; i++ {
		bar := termbar.Bar{
			Label: fmt.Sprintf("%d", i),
			Value: float64(rand.Intn(99) + 1),
		}
		current[i] = barWithVelocity{
			Bar:      bar,
			Velocity: 0,
		}
		target[i] = bar
	}

	m := model{
		spring:  harmonica.NewSpring(harmonica.FPS(60), 12.0, .7),
		current: current,
		target:  target,
		chartOpts: []termbar.HorizontalOption{
			termbar.WithMaxVal(100.0),
			termbar.WithValueFormatter(func(b termbar.Bar) string { return "" }),
			termbar.WithLabelSeparator(" "),
		},
	}

	return &m
}

func (m *model) Init() tea.Cmd {
	m.updateTarget()
	return tea.Batch(
		animate(),
		updateTargetTick(),
	)
}

func animate() tea.Cmd {
	return tea.Tick(time.Second/60, func(t time.Time) tea.Msg {
		return frameMsg(t)
	})
}

func updateTargetTick() tea.Cmd {
	return tea.Tick(2*time.Second, func(t time.Time) tea.Msg {
		return updateMsg(t)
	})
}

func (m *model) updateTarget() {
	newBars := make(termbar.Bars, len(m.target))
	for b, bar := range m.target {
		newBars[b] = termbar.Bar{
			Label: bar.Label,
			Value: float64(rand.Intn(99) + 1),
		}
	}
	m.target = newBars
}

func (m *model) updateCurrent() {
	newBars := make([]barWithVelocity, len(m.target))
	for b, bv := range m.current {
		n, nv := m.spring.Update(bv.Value, bv.Velocity, m.target[b].Value)
		newBars[b] = barWithVelocity{
			Bar: termbar.Bar{
				Label: bv.Bar.Label,
				Value: n,
			},
			Velocity: nv,
		}
	}
	m.current = newBars
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case tea.KeyMsg:
		return m, tea.Quit
	case updateMsg:
		m.updateTarget()
		return m, updateTargetTick()
	case frameMsg:
		m.updateCurrent()
		return m, animate()
	default:
		return m, nil
	}
}

func (m *model) View() string {
	var buf bytes.Buffer
	bars := make(termbar.Bars, len(m.current))
	for b, bv := range m.current {
		bars[b] = bv.Bar
	}
	termbar.NewHorizontal(bars, m.chartOpts...).Print(&buf)
	return buf.String()
}
