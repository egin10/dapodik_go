package form

import (
	"fmt"
	"gin-dapodik/model"
	"gin-dapodik/scraper"
	"os"

	"github.com/charmbracelet/huh"
)

func generateListOptionsProvinsi(satuanPendidikan model.SatuanPendidikan) []huh.Option[model.Provinsi] {
	dataList := []huh.Option[model.Provinsi]{}
	for _, value := range scraper.GetListProvinsiBySatuanPendidikan(satuanPendidikan) {
		option := huh.NewOption(value.Name, value)
		dataList = append(dataList, option)
	}
	return dataList
}

func FormProvinsi(options *model.Options, accessible bool) {
	form := huh.NewForm(
		huh.NewGroup(
			// Pilih Provinsi
			huh.NewSelect[model.Provinsi]().
				Options(generateListOptionsProvinsi(options.SatuanPendidikan)...).
				Title("Provinsi").
				Description("Silahkan pilih provinsi yang diinginkan").
				Value(&options.Provinsi),
		),
	).WithAccessible(accessible)

	errTwo := form.Run()
	if errTwo != nil {
		fmt.Println("Yah error:", errTwo)
		os.Exit(1)
	}
}
