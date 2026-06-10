package tui

import (
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"

	"yearprogress/internal/core"
)

const tickInterval = 50 * time.Millisecond

type tickMsg time.Time

type model struct {
	state  core.TimeState
	width  int
	height int
	config core.Config
}

func newModel(cfg core.Config) model {
	return model{
		state:  core.ComputeTimeState(time.Now()),
		config: cfg,
	}
}

func (m model) Init() tea.Cmd {
	return m.tick()
}

func (m model) tick() tea.Cmd {
	return tea.Tick(tickInterval, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if core.IsExitKey(msg.String()) {
			return m, tea.Quit
		}

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tickMsg:
		m.state = core.ComputeTimeState(time.Time(msg))
		return m, m.tick()
	}

	return m, nil
}

func (m model) View() string {
	return Render(core.BuildScene(m.state, m.config), m.width, m.height)
}

// Run starts the terminal UI.
func Run(cfg core.Config) error {
	p := tea.NewProgram(newModel(cfg), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		return fmt.Errorf("tui: %w", err)
	}
	return nil
}

func RunOrExit(cfg core.Config) {
	if err := Run(cfg); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}