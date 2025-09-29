package tui

import (
	"os"
	"path/filepath"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
)

func RunTui(directoryWithTransforms string) error {
	files, err := getFilesFromDirectory(directoryWithTransforms)
	if err != nil {
		return err
	}

	var items []list.Item
	for _, file := range files {
		items = append(items, file)
	}

	l := list.New(items, itemDelegate{}, 20, 25)
	l.Title = "Transform files"

	m := model{list: l}

	if _, err := tea.NewProgram(m).Run(); err != nil {
		return err
	}

	return nil
}

func getFilesFromDirectory(directoryWithTransforms string) ([]item, error) {
	entries, err := os.ReadDir(directoryWithTransforms)
	if err != nil {
		return nil, err
	}

	var files []item

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		path := filepath.Join(directoryWithTransforms, entry.Name())
		files = append(files, item{
			fileName: entry.Name(),
			path:     path,
		})
	}

	return files, nil
}
