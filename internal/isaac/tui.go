package isaac

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	choices []string
	cursor  int
	checks  map[int]string
}

func StartScraper() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error")
		os.Exit(1)
	}
}

func initialModel() model {
	return model{
		choices: []string{"Items", "Characters", "Trinkets", "Transformations", "Pills"},
		checks:  make(map[int]string),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {

	case tea.KeyMsg:

		switch msg.String() {
		case "ctrl + c", "q":
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		case " ":
			_, ok := m.checks[m.cursor]
			if ok {
				delete(m.checks, m.cursor)
			} else {
				m.checks[m.cursor] = m.choices[m.cursor]
			}

		case "enter":

			if len(m.checks) <= 0 {
				return m, nil
			}

			fmt.Println("Starting Download")
			createSelectedCsv(m.checks)
			fmt.Println("Dowloading finished, exiting...")

			return m, tea.Quit
		}
	}

	return m, nil
}

func (m model) View() string {

	var style = lipgloss.NewStyle().
		Bold(true).
		Align(lipgloss.Center).
		Margin(1, 0).
		Border(lipgloss.RoundedBorder())

	s := "Mark what you want to download\n\n"

	for i, choice := range m.choices {
		cursor := " "

		if m.cursor == i {
			cursor = ">"
		}

		checked := " "
		if _, ok := m.checks[i]; ok {
			checked = "†"
		}

		s += fmt.Sprintf("%s %s %s\n", cursor, checked, choice)

	}

	p := "\nSpacebar to start  download.\n"

	p += "\nPress q to quit.\n"

	return fmt.Sprintf("%s%s", s, style.Render(p))
}

func createSelectedCsv(checks map[int]string) {

	categories := map[string]func() error{
		"Transformations": CreateTransformationsCsv,
		"Pills":           CreatePillsCsv,
		"Items":           CreateItemsCsv,
		"Trinkets":        CreateTrinketsCsv,
		"Characters":      CreateCharactersCsv,
	}

	for _, check := range checks {
		if categoryFunc, exists := categories[check]; exists {
			if err := categoryFunc(); err != nil {
				fmt.Printf("Error: %s", err)
				os.Exit(1)
			}
		}
	}

}
