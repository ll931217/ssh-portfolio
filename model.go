package main

import (
	"fmt"
	"sort"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	titleStyle = func() lipgloss.Style {
		return lipgloss.NewStyle().Padding(0, 1)
	}()

	selectedTitleStyle = func() lipgloss.Style {
		return titleStyle.Underline(true)
	}()
)

type model struct {
	pages     map[int]page
	pageIndex int
	viewport  viewport.Model
	ready     bool
}

func initModel(pgs pages) (model, error) {
	return model{pages: pgs}, nil
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	lineJumps := 3
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "l", "right":
			if m.pageIndex < len(m.pages)-1 {
				m.pageIndex++
			}
			m.viewport.SetContent(m.pages[m.pageIndex].rendered)
			m.viewport.GotoTop()
		case "h", "left":
			if m.pageIndex > 0 {
				m.pageIndex--
			}
			m.viewport.SetContent(m.pages[m.pageIndex].rendered)
			m.viewport.GotoTop()
		case "j", "down":
			m.viewport.LineDown(lineJumps)
		case "k", "up":
			m.viewport.LineUp(lineJumps)
		case "g":
			m.viewport.GotoTop()
		case "G":
			m.viewport.GotoBottom()
		default:
			m.viewport.Update(msg)
		}
	case tea.WindowSizeMsg:
		headerHeight := lipgloss.Height(m.headerView())
		footerHeight := lipgloss.Height(m.helpView())
		verticalMarginHeight := headerHeight + footerHeight

		if !m.ready {
			m.viewport = viewport.New(msg.Width, msg.Height-verticalMarginHeight)
			m.viewport.YPosition = headerHeight
			m.viewport.SetContent(m.pages[m.pageIndex].rendered)
			m.ready = true
		} else {
			m.viewport.Width = msg.Width
			m.viewport.Height = msg.Height - verticalMarginHeight
		}

	}

	var cmd tea.Cmd
	m.viewport, cmd = m.viewport.Update(msg)
	return m, tea.Batch(cmd)
}

func (m model) View() string {
	if !m.ready {
		return "Loading..."
	}

	return fmt.Sprintf(
		"%s\n%s\n%s",
		m.headerView(),
		m.viewport.View(),
		m.helpView(),
	)
}

func (m model) helpView() string {
	return lipgloss.JoinHorizontal(
		lipgloss.Left,
		titleStyle.Render("h: Left"),
		titleStyle.Render("j: Down"),
		titleStyle.Render("k: Up"),
		titleStyle.Render("l: Right"),
		titleStyle.Render("g: Top"),
		titleStyle.Render("G: Bottom"),
		titleStyle.Render("q: Quit"),
	)
}

func (m model) headerView() string {
	indices := make([]int, 0, len(m.pages))
	for i := range m.pages {
		indices = append(indices, i)
	}
	sort.Ints(indices)

	var pagesTotalWidth int
	var header []string
	for _, i := range indices {
		var pageRender string
		if i == m.pageIndex {
			pageRender = selectedTitleStyle.Render(m.pages[i].name)
		} else {
			pageRender = titleStyle.Render(m.pages[i].name)
		}

		header = append(header, pageRender)
		pagesTotalWidth += lipgloss.Width(pageRender)
	}

	return lipgloss.JoinHorizontal(lipgloss.Center, header...)
}
