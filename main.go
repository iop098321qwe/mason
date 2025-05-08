package main

import (
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/bubbles/progress"
)

// ---------------- CONFIG ----------------

const (
	workDuration      = 25 * time.Minute
	shortBreakDuration = 5 * time.Minute
	longBreakDuration  = 15 * time.Minute
	cyclesBeforeLongBreak = 4
	tickInterval       = time.Second
)

// --------------- MODEL ------------------

type mode int

const (
	modeWork mode = iota
	modeShortBreak
	modeLongBreak
)

type model struct {
	startTime   time.Time
	elapsed     time.Duration
	duration    time.Duration
	state       string
	currentMode mode
	cycleCount  int
	paused      bool
	progressBar progress.Model
}

type tickMsg time.Time
type errMsg error

// -------------- INITIALIZE --------------

func initialModel() model {
	p := progress.New(progress.WithDefaultGradient())
	return model{
		startTime:   time.Now(),
		elapsed:     0,
		duration:    workDuration,
		state:       "Pomodoro Running (Work)",
		currentMode: modeWork,
		cycleCount:  0,
		paused:      false,
		progressBar: p,
	}
}

func (m model) Init() tea.Cmd {
	return tick()
}

// --------------- UPDATE -----------------

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		case "p":
			m.paused = !m.paused
			if m.paused {
				m.state = "Paused"
			} else {
				m.startTime = time.Now().Add(-m.elapsed)
				m.state = "Resumed"
			}
			return m, tick()
		}

	case tickMsg:
		if m.paused {
			return m, tick()
		}

		m.elapsed = time.Since(m.startTime)

		if m.elapsed >= m.duration {
			// Transition to next mode
			return nextMode(m), tick()
		}

		return m, tick()
	}

	return m, nil
}

// ---------- MODE TRANSITIONS ------------

func nextMode(m model) model {
	m.elapsed = 0

	switch m.currentMode {
	case modeWork:
		m.cycleCount++
		if m.cycleCount%cyclesBeforeLongBreak == 0 {
			m.currentMode = modeLongBreak
			m.duration = longBreakDuration
			m.state = "Long Break"
		} else {
			m.currentMode = modeShortBreak
			m.duration = shortBreakDuration
			m.state = "Short Break"
		}
	case modeShortBreak, modeLongBreak:
		m.currentMode = modeWork
		m.duration = workDuration
		m.state = "Work"
	}

	m.startTime = time.Now()
	return m
}

// --------------- VIEW -------------------

func (m model) View() string {
	progressPercent := float64(m.elapsed.Seconds()) / float64(m.duration.Seconds())
	bar := m.progressBar.ViewAs(progressPercent)

	return fmt.Sprintf(
		"🍅 Pomodoro Timer\n\nState: %s\nTime Remaining: %s\n\n%s\n\n[ p: pause/resume | q: quit ]",
		m.state,
		formatDuration(m.duration - m.elapsed),
		bar,
	)
}

// ---------- TICKER FUNCTION -------------

func tick() tea.Cmd {
	return tea.Tick(tickInterval, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

// --------- FORMAT TIME OUTPUT -----------

func formatDuration(d time.Duration) string {
	if d < 0 {
		return "00:00"
	}
	mins := int(d.Minutes())
	secs := int(d.Seconds()) % 60
	return fmt.Sprintf("%02d:%02d", mins, secs)
}

// --------------- MAIN -------------------

func main() {
	p := tea.NewProgram(initialModel())
	if err := p.Start(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
