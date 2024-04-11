package form

import (
	"gin-dapodik/model"
	"gin-dapodik/scraper"

	"github.com/charmbracelet/huh"
)

func generateListOptionsSatuanPendidikan() []huh.Option[model.SatuanPendidikan] {
	dataList := []huh.Option[model.SatuanPendidikan]{}
	for _, value := range scraper.GetListSatuanPendidikan() {
		option := huh.NewOption(value.Name, value)
		dataList = append(dataList, option)
	}
	return dataList
}

func FormSatuanPendidikan(optionValue *model.OptionValue, accessible bool) *huh.Form {
	form := huh.NewForm(
		// Menu utama
		huh.NewGroup(
			// Pilih Satuan Pendidikan
			huh.NewSelect[model.SatuanPendidikan]().
				Key("satuan_pendidikan").
				Options(generateListOptionsSatuanPendidikan()...).
				Title("Satuan Pendidikan").
				Description("Silahkan pilih satuan pendidikan.").
				Value(&optionValue.SatuanPendidikan),
		),
	).WithAccessible(accessible)

	return form
}
