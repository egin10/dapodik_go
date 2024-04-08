package main

// import (
// 	"fmt"
// 	"gin-dapodik/form"
// 	"gin-dapodik/model"
// 	"gin-dapodik/scraper"
// 	"gin-dapodik/utils"
// 	"os"
// 	"strconv"
// 	"strings"

// 	"github.com/charmbracelet/huh/spinner"
// 	"github.com/charmbracelet/lipgloss"

// 	"github.com/common-nighthawk/go-figure"
// )

// var (
// 	options         = model.Options{}
// 	listDataSekolah = make([]model.DataSekolah, 0)
// )

// func mainHuh() {
// 	// Should we run in accessible mode?
// 	accessible, _ := strconv.ParseBool(os.Getenv("ACCESSIBLE"))

// 	// TITLE APP
// 	myFigure := figure.NewFigure("GIN-DAPODIK", "", true)
// 	myFigure.Print()
// 	fmt.Println()

// 	form.FormSatuanPendidikan(&options, accessible)
// 	form.FormProvinsi(&options, accessible)
// 	form.FormExportTo(&options, accessible)
// 	form.FormConfirm(&options, accessible)

// 	// Get Semua Data Sekolah di Provinsi terpilih
// 	processData := func() {
// 		// Get List Url Kabupaten/Kota
// 		listUrlKabKota := scraper.GetListDataUrl(options.Provinsi.Url)
// 		for _, urlKabKota := range listUrlKabKota {
// 			// Get Url Kecamatan
// 			listUrlKecamatan := scraper.GetListDataUrl(urlKabKota.Url)
// 			for _, urlKecamatan := range listUrlKecamatan {
// 				// Get Url Sekolah
// 				listUrlSekolah := scraper.GetListDataUrl(urlKecamatan.Url)
// 				for _, urlSekolah := range listUrlSekolah {
// 					// Get Data Detail Sekolah
// 					dataSekolah := scraper.GetDataSekolah(urlSekolah.Url)

// 					listDataSekolah = append(listDataSekolah, dataSekolah)
// 				}
// 			}
// 		}
// 	}

// 	_ = spinner.New().
// 		Title(fmt.Sprintf("Downloading data sekolah %s di provinsi %s...", options.SatuanPendidikan.Name, options.Provinsi.Name)).
// 		Accessible(accessible).
// 		Action(processData).
// 		Run()

// 	// Finished
// 	{
// 		var sb strings.Builder
// 		keyword := func(s string) string {
// 			return lipgloss.NewStyle().Foreground(lipgloss.Color("212")).Render(s)
// 		}
// 		fmt.Fprintf(&sb,
// 			"%s\n\nData sekolah untuk satuan pendidikan %s pada provinsi %s\nSelesai diunduh 👾.\n\nTerima kasih sudah menggunakan GIN-DAPODIK!",
// 			lipgloss.NewStyle().Bold(true).Render("GIN-DAPODIK 👾"),
// 			keyword("PAUD"),
// 			keyword("DKI JAKARTA"),
// 		)

// 		fmt.Println(
// 			lipgloss.NewStyle().
// 				Width(40).
// 				BorderStyle(lipgloss.RoundedBorder()).
// 				BorderForeground(lipgloss.Color("63")).
// 				Padding(1, 2).
// 				Render(sb.String()),
// 		)

// 		switch options.ExportTo {
// 		case "xlsx":
// 			// Generate to xlsx file
// 			errxlsx := utils.WireToExcel(listDataSekolah, options.SatuanPendidikan, options.Provinsi)
// 			if errxlsx != nil {
// 				fmt.Println("Unable to create xlsx file 🗿")
// 			}
// 		case "json":
// 			// Generate to json file
// 			errjson := utils.WriteJSON(listDataSekolah, options.SatuanPendidikan, options.Provinsi)
// 			if errjson != nil {
// 				fmt.Println("Unable to create json file 🗿")
// 			}
// 		}
// 	}
// }
