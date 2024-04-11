package main

import (
	"fmt"
	"gin-dapodik/form"
	"gin-dapodik/model"
	"gin-dapodik/scraper"
	"gin-dapodik/utils"
	"os"
	"strconv"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/common-nighthawk/go-figure"
)

type Model struct {
	header    string
	menuIndex int
	form      *huh.Form

	optionValue model.OptionValue

	accessible bool
	process    string
}

type done int

func initialModel() Model {
	accessible, _ := strconv.ParseBool(os.Getenv("ACCESSIBLE"))

	myFigure := figure.NewFigure("GIN-DAPODIK", "", true)
	welcomeMsg := "Aplikasi ini hanya bisa digunakan untuk mengambil data sekolah\n"
	welcomeMsg += "sesuai dengan satuan pendidikan & provinsi yang diinginkan pada website\n"
	welcomeMsg += "https://referensi.data.kemdikbud.go.id\n\n"
	welcomeMsg += "Tekan ctrl+c untuk close program"

	header := fmt.Sprintf(
		"%s\n%s",
		myFigure.String(),
		welcomeMsg,
	)

	return Model{
		menuIndex:   0,
		optionValue: model.OptionValue{},
		accessible:  accessible,
		process:     "",
		header:      header,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit
		}

	case done:
		m.menuIndex = 0
		m.process = fmt.Sprintf("Data selesai diunduh dengan format file (.%s) 🗿\n", m.optionValue.ExportTo)
		m.process += "Tekan Enter untuk kembali ke menu awal.\n"
		// m.process += "Tekan ctrl+c untuk menutup aplikasi.\n"
		return m, nil
	}

	switch m.menuIndex {
	case 0:
		m.reset()
		m.form = form.FormSatuanPendidikan(&m.optionValue, m.accessible)
		m.menuIndex++
	case 1:
		m.form = form.FormProvinsi(&m.optionValue, m.accessible)
		m.menuIndex++
	case 2:
		m.form = form.FormExportTo(&m.optionValue, m.accessible)
		m.menuIndex++
	case 3:
		m.form = form.FormConfirm(&m.optionValue, m.accessible)
		m.menuIndex++

		str := fmt.Sprintf(
			"🗂️  Sedang mengunduh data %s di %s...\nProses ini tergantung koneksi internet anda.\n",
			strings.ToUpper(m.optionValue.SatuanPendidikan.Path),
			strings.ToUpper(m.optionValue.Provinsi.Name),
		)
		m.process = str

	case 4:
		return m, m.scraping()
	}

	err := m.form.Run()
	if err != nil {
		fmt.Println("🗿 Yah error:", err)
		os.Exit(1)
	}

	return m, nil
}

func (m Model) View() string {
	return fmt.Sprintf(
		"%s\n%s\n",
		m.header,
		m.process,
	)
}

func (m Model) scraping() tea.Cmd {
	listDataSekolah := make([]model.DataSekolah, 0)

	listUrlKabKota := scraper.GetListDataUrl(m.optionValue.Provinsi.Url)
	for _, urlKabKota := range listUrlKabKota {
		// Get Url semua Kecamatan
		listUrlKecamatan := scraper.GetListDataUrl(urlKabKota.Url)
		for _, urlKecamatan := range listUrlKecamatan {
			// Get Url semua Sekolah
			listUrlSekolah := scraper.GetListDataUrl(urlKecamatan.Url)
			for _, urlSekolah := range listUrlSekolah {
				// Get Data Detail Sekolah
				dataSekolah := scraper.GetDataSekolah(urlSekolah.Url)

				listDataSekolah = append(listDataSekolah, dataSekolah)
			}
		}
	}

	switch m.optionValue.ExportTo {
	case "xlsx":
		// Generate to xlsx file
		errxlsx := utils.WireToExcel(listDataSekolah, m.optionValue.SatuanPendidikan, m.optionValue.Provinsi)
		if errxlsx != nil {
			fmt.Println("Unable to create xlsx file 🗿")
		}
	case "json":
		// Generate to json file
		errjson := utils.WriteToJSON(listDataSekolah, m.optionValue.SatuanPendidikan, m.optionValue.Provinsi)
		if errjson != nil {
			fmt.Println("Unable to create json file 🗿")
		}
	}

	return func() tea.Msg {
		return done(1)
	}
}

func (m *Model) reset() {
	m.menuIndex = 0
	m.process = ""
	m.optionValue.Provinsi = model.Provinsi{}
	m.optionValue.SatuanPendidikan = model.SatuanPendidikan{}
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
