package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// SessionState defines the current state of the timer
type SessionState int

const (
	StateWork SessionState = iota
	StateShortBreak
	StateLongBreak
)

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

	// Base timer style
	timerBaseStyle = lipgloss.NewStyle().
			Bold(true).
			Height(3).
			Width(25).
			Align(lipgloss.Center, lipgloss.Center).
			Border(lipgloss.RoundedBorder()).
			MarginBottom(1)

	workStyle = timerBaseStyle.Copy().
			Foreground(lipgloss.Color("#FF5F87")).
			BorderForeground(lipgloss.Color("#FF5F87"))

	breakStyle = timerBaseStyle.Copy().
			Foreground(lipgloss.Color("#04B575")).
			BorderForeground(lipgloss.Color("#04B575"))

	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#626262")).
			MarginTop(1)
	
	progressStyle = lipgloss.NewStyle().
			PaddingLeft(2).
			PaddingRight(2)

	sessionStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#AAA")).
			Italic(true).
			MarginTop(1)

	statusStyle = lipgloss.NewStyle().
			Bold(true).
			MarginTop(1)
)

type tickMsg time.Time

type model struct {
	workDuration       time.Duration
	shortBreakDuration time.Duration
	longBreakDuration  time.Duration
	
	duration     time.Duration
	timeLeft     time.Duration
	running      bool
	state        SessionState
	sessionCount int
	progress     progress.Model
}

func initialModel(work, short, long int) model {
	w := time.Duration(work) * time.Minute
	return model{
		workDuration:       w,
		shortBreakDuration: time.Duration(short) * time.Minute,
		longBreakDuration:  time.Duration(long) * time.Minute,
		duration:           w,
		timeLeft:           w,
		running:            false,
		state:              StateWork,
		sessionCount:       0,
		progress:           progress.New(progress.WithDefaultGradient()),
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
			if !m.running && m.timeLeft > 0 {
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
		m.progress.Width = msg.Width - 10
		if m.progress.Width > 80 {
			m.progress.Width = 80
		}
		return m, nil
	case tickMsg:
		if m.running && m.timeLeft > 0 {
			m.timeLeft -= time.Second
			percent := 1.0 - (float64(m.timeLeft) / float64(m.duration))
			cmd := m.progress.SetPercent(percent)
			return m, tea.Batch(tick(), cmd)
		} else if m.running && m.timeLeft <= 0 {
			m.running = false
			m.progress.SetPercent(1.0)
			// Audible alert
			fmt.Print("\a")
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
	m.timeLeft = m.duration
	m.progress.SetPercent(0)
}

func (m *model) nextSession() {
	m.running = false
	if m.state == StateWork {
		m.sessionCount++
		if m.sessionCount%4 == 0 {
			m.state = StateLongBreak
			m.duration = m.longBreakDuration
		} else {
			m.state = StateShortBreak
			m.duration = m.shortBreakDuration
		}
	} else {
		m.state = StateWork
		m.duration = m.workDuration
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

	header := titleStyle.Render("FocusBrew")
	
	statusText := fmt.Sprintf("%s", m.state)
	timeStr := fmt.Sprintf("%02d:%02d", minutes, seconds)

	var style lipgloss.Style
	if m.state == StateWork {
		style = workStyle
	} else {
		style = breakStyle
	}
	
	timerView := style.Render(
		lipgloss.JoinVertical(lipgloss.Center,
			statusText,
			timeStr,
		),
	)

	progView := progressStyle.Render(m.progress.View())
	
	sessions := sessionStyle.Render(fmt.Sprintf("Completed Sessions: %d", m.sessionCount))

	var currentStatus string
	if m.running {
		currentStatus = statusStyle.Foreground(lipgloss.Color("#04B575")).Render("RUNNING")
	} else if m.timeLeft <= 0 {
		currentStatus = statusStyle.Foreground(lipgloss.Color("#FF5F87")).Render("FINISHED")
	} else {
		currentStatus = statusStyle.Foreground(lipgloss.Color("#626262")).Render("PAUSED")
	}

	help := helpStyle.Render("s: start • p: pause • r: reset • n: next • q: quit")

	return lipgloss.JoinVertical(lipgloss.Center,
		header,
		timerView,
		progView,
		currentStatus,
		sessions,
		help,
	)
}

func main() {
	work := flag.Int("work", 25, "Work duration in minutes")
	short := flag.Int("short", 5, "Short break duration in minutes")
	long := flag.Int("long", 15, "Long break duration in minutes")
	flag.Parse()

	p := tea.NewProgram(initialModel(*work, *short, *long))
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}