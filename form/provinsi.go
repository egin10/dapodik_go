package form

import (
	"gin-dapodik/model"
	"gin-dapodik/scraper"

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

func FormProvinsi(optionValue *model.OptionValue, accessible bool) *huh.Form {
	form := huh.NewForm(
		huh.NewGroup(
			// Pilih Provinsi
			huh.NewSelect[model.Provinsi]().
				Options(generateListOptionsProvinsi(optionValue.SatuanPendidikan)...).
				Key("provinsi").
				Title("Provinsi").
				Description("Silahkan pilih provinsi yang diinginkan").
				Value(&optionValue.Provinsi),
		),
	).WithAccessible(accessible)

	return form
}
