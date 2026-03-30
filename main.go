package main

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	keyStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("15")).
			Background(lipgloss.Color("240")).
			Padding(0, 1).
			Margin(0, 1)

	activeKeyStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("15")).
			Background(lipgloss.Color("82")).
			Padding(0, 1).
			Margin(0, 1)

	resultStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("82")).
			Bold(true)

	subStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("245"))
)

type model struct {
	state     string
	startTime time.Time
	endTime   time.Time
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.Type == tea.KeyEsc || msg.String() == "q" {
			return m, tea.Quit
		}
		if m.state == "waiting" {
			if msg.Type == tea.KeyUp || msg.Type == tea.KeyDown ||
				msg.Type == tea.KeyLeft || msg.Type == tea.KeyRight {
				m.state = "started"
				m.startTime = time.Now()
			}
		} else if m.state == "started" {
			if msg.Type == tea.KeyEnter {
				m.state = "finished"
				m.endTime = time.Now()
			}
		} else if m.state == "finished" {
			if msg.Type == tea.KeyEnter {
				m.state = "waiting"
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	var content string

	switch m.state {
	case "waiting":
		content = fmt.Sprintf(`%s

%s

%s

%s

%s`, activeKeyStyle.Render("↑"),
			keyStyle.Render("ENTER"),
			subStyle.Render("Press ↑ to start"),
			subStyle.Render("Press Enter as fast as you can"),
			subStyle.Render("Press Esc or q to quit"))
	case "started":
		content = fmt.Sprintf(`%s

%s

%s

%s`, keyStyle.Render("↑"),
			activeKeyStyle.Render("ENTER"),
			subStyle.Render("GO!"),
			subStyle.Render("Press Esc or q to quit"))
	case "finished":
		duration := m.endTime.Sub(m.startTime)
		content = fmt.Sprintf(`%s

%s

%s

%s

%s`, resultStyle.Render(fmt.Sprintf("%v", duration)),
			keyStyle.Render("↑"),
			activeKeyStyle.Render("ENTER"),
			subStyle.Render("Press Enter to try again"),
			subStyle.Render("Press Esc or q to quit"))
	}

	w := 20
	if lipgloss.Width(content) > w {
		w = lipgloss.Width(content)
	}
	h := 12
	return lipgloss.Place(w, h, lipgloss.Center, lipgloss.Center, content)
}

func main() {
	p := tea.NewProgram(model{state: "waiting"})
	p.Run()
}
