package main

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Constants for timer durations
const (
	DefaultWorkDuration       = 25 * time.Minute
	DefaultShortBreakDuration = 5 * time.Minute
	DefaultLongBreakDuration  = 15 * time.Minute
)

// SessionState defines the current state of the timer
type SessionState int

const (
	StateWork SessionState = iota
	StateShortBreak
	StateLongBreak
)

// String returns the string representation of the session state
func (s SessionState) String() string {
	switch s {
	case StateWork:
		return "Work Session"
	case StateShortBreak:
		return "Short Break"
	case StateLongBreak:
		return "Long Break"
	default:
		return "Unknown"
	}
}

// Styles
var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(lipgloss.Color("#7D56F4")).
			Padding(0, 2).
			MarginBottom(1)

	timerStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#04B575")).
			Bold(true).
			Height(3).
			Width(25).
			Align(lipgloss.Center, lipgloss.Center).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#7D56F4")).
			MarginBottom(1)

	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#626262")).
			MarginTop(1)
	
	progressStyle = lipgloss.NewStyle().
			PaddingLeft(2).
			PaddingRight(2)
)

type tickMsg time.Time

type model struct {
	duration     time.Duration
	timeLeft     time.Duration
	running      bool
	state        SessionState
	sessionCount int
	progress     progress.Model
}

func initialModel() model {
	return model{
		duration:     DefaultWorkDuration,
		timeLeft:     DefaultWorkDuration,
		running:      false,
		state:        StateWork,
		sessionCount: 0,
		progress:     progress.New(progress.WithDefaultGradient()),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "s":
			if !m.running {
				m.running = true
				return m, tick()
			}
		case "p":
			m.running = false
		case "r":
			m.resetTimer()
		case "n":
			m.nextSession()
		}
	case tea.WindowSizeMsg:
		m.progress.Width = msg.Width - 10 // Padding
		if m.progress.Width > 80 {
			m.progress.Width = 80
		}
		return m, nil
	case tickMsg:
		if m.running && m.timeLeft > 0 {
			m.timeLeft -= time.Second
			
			// Update progress
			percent := 1.0 - (float64(m.timeLeft) / float64(m.duration))
			cmd := m.progress.SetPercent(percent)
			
			return m, tea.Batch(tick(), cmd)
		} else if m.running && m.timeLeft <= 0 {
			m.running = false
			m.progress.SetPercent(1.0)
			// Here we could auto-advance or just stop
		}
	case progress.FrameMsg:
		progressModel, cmd := m.progress.Update(msg)
		m.progress = progressModel.(progress.Model)
		return m, cmd
	}
	return m, nil
}

func (m *model) resetTimer() {
	m.running = false
	switch m.state {
	case StateWork:
		m.duration = DefaultWorkDuration
	case StateShortBreak:
		m.duration = DefaultShortBreakDuration
	case StateLongBreak:
		m.duration = DefaultLongBreakDuration
	}
	m.timeLeft = m.duration
	m.progress.SetPercent(0)
}

func (m *model) nextSession() {
	m.running = false
	// Logic to cycle sessions: Work -> Short Break -> Work -> Short Break -> Work -> Long Break
	if m.state == StateWork {
		m.sessionCount++
		if m.sessionCount%4 == 0 {
			m.state = StateLongBreak
			m.duration = DefaultLongBreakDuration
		} else {
			m.state = StateShortBreak
			m.duration = DefaultShortBreakDuration
		}
	} else {
		// From break to work
		m.state = StateWork
		m.duration = DefaultWorkDuration
	}
	m.timeLeft = m.duration
	m.progress.SetPercent(0)
}

func tick() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func (m model) View() string {
	minutes := int(m.timeLeft.Minutes())
	seconds := int(m.timeLeft.Seconds()) % 60

	header := titleStyle.Render("FocusBrew 🍵")
	
	status := fmt.Sprintf("%s", m.state)
	timeStr := fmt.Sprintf("%02d:%02d", minutes, seconds)
	
	timerView := timerStyle.Render(
		lipgloss.JoinVertical(lipgloss.Center,
			status,
			timeStr,
		),
	)

	// Update the logic to reflect current progress state even if paused
	// Just reusing the model's progress view
	progView := progressStyle.Render(m.progress.View())

	help := helpStyle.Render("s: start • p: pause • r: reset • n: next • q: quit")

	return lipgloss.JoinVertical(lipgloss.Center,
		header,
		timerView,
		progView,
		help,
	)
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
	}
}