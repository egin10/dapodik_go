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

	"github.com/charmbracelet/huh/spinner"
	"github.com/charmbracelet/lipgloss"

	"github.com/common-nighthawk/go-figure"
)

var (
	options         = model.Options{}
	listDataSekolah = make([]model.DataSekolah, 0)
)

func main() {
	// Should we run in accessible mode?
	accessible, _ := strconv.ParseBool(os.Getenv("ACCESSIBLE"))

	// TITLE APP
	myFigure := figure.NewFigure("GIN-DAPODIK", "", true)
	myFigure.Print()
	fmt.Println()

	form.FormSatuanPendidikan(&options, accessible)
	form.FormProvinsi(&options, accessible)
	form.FormConfirm(&options, accessible)

	// Get Semua Data Sekolah di Provinsi terpilih
	processData := func() {
		// Get List Url Kabupaten/Kota
		listUrlKabKota := scraper.GetListDataUrl(options.Provinsi.Url)
		for _, urlKabKota := range listUrlKabKota {
			// fmt.Printf("\tKAB.KOTA : %s\n", urlKabKota.Name)

			// Get Url Kecamatan
			listUrlKecamatan := scraper.GetListDataUrl(urlKabKota.Url)
			for _, urlKecamatan := range listUrlKecamatan {
				// fmt.Printf("\tKECAMATAN : %s\n", urlKecamatan.Name)

				// Get Url Sekolah
				listUrlSekolah := scraper.GetListDataUrl(urlKecamatan.Url)
				for _, urlSekolah := range listUrlSekolah {
					// Get Data Detail Sekolah
					dataSekolah := scraper.GetDataSekolah(urlSekolah.Url)
					fmt.Printf("Nama: %s\n", dataSekolah.IdentitasSekolah.Nama)
					fmt.Printf("NPSN: %s\n", dataSekolah.IdentitasSekolah.NPSN)
					fmt.Printf("Alamat: %s\n", dataSekolah.IdentitasSekolah.Alamat)
					fmt.Printf("BentukPendidikan: %s\n", dataSekolah.IdentitasSekolah.BentukPendidikan)
					fmt.Println("--------------------------------------")

					listDataSekolah = append(listDataSekolah, dataSekolah)
				}
			}
			fmt.Println("=======================================")
		}
	}

	_ = spinner.New().
		Title(fmt.Sprintf("Downloading data sekolah %s di provinsi %s...", options.SatuanPendidikan.Name, options.Provinsi.Name)).
		Accessible(accessible).
		Action(processData).
		Run()

	// Finished
	{
		var sb strings.Builder
		keyword := func(s string) string {
			return lipgloss.NewStyle().Foreground(lipgloss.Color("212")).Render(s)
		}
		fmt.Fprintf(&sb,
			"%s\n\nData sekolah untuk satuan pendidikan %s pada provinsi %s\nSelesai diunduh 👾.\n\nTerima kasih sudah menggunakan GIN-DAPODIK!",
			lipgloss.NewStyle().Bold(true).Render("GIN-DAPODIK 👾"),
			keyword("PAUD"),
			keyword("DKI JAKARTA"),
		)

		fmt.Println(
			lipgloss.NewStyle().
				Width(40).
				BorderStyle(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("63")).
				Padding(1, 2).
				Render(sb.String()),
		)

		// Generate to JSON file
		utils.WriteJSON(listDataSekolah)
	}
}
