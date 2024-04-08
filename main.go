package main

import (
	"fmt"
	"gin-dapodik/form"
	"gin-dapodik/model"
	"log"
	"os"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/common-nighthawk/go-figure"
)

func main() {
	p := tea.NewProgram(newModel())

	_, err := p.Run()
	if err != nil {
		log.Fatalln(err)
	}
}

type Model struct {
	menuIndex int
	form      *huh.Form

	optionValue model.OptionValue
}

func newModel() Model {
	accessible, _ := strconv.ParseBool(os.Getenv("ACCESSIBLE"))

	m := Model{}
	m.menuIndex = 0
	m.optionValue = model.OptionValue{
		SatuanPendidikan: model.SatuanPendidikan{},
		Provinsi:         model.Provinsi{},
		ExportTo:         "",
	}
	m.form = form.FormSatuanPendidikan(&m.optionValue, accessible)

	return m
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			// switch m.menuIndex {
			// case 0:
			// }
			err := m.form.Run()
			if err != nil {
				fmt.Println("Yah error:", err)
				os.Exit(1)
			}
			return m, nil
		case tea.KeyCtrlC:
			return m, tea.Quit
		}
	}

	return m, cmd
}

func (m Model) View() string {
	myFigure := figure.NewFigure("GIN-DAPODIK", "", true)
	devider := "======================================================="
	return fmt.Sprintf(
		"%s\n%s\n",
		myFigure.String(),
		devider,
	)
}

// Cmd
