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
		// Welcome Message
		huh.NewGroup(huh.NewNote().
			Description("Selamat datang di Gin-Dapodik👾.\n\n").
			Description("Aplikasi ini hanya bisa digunakan untuk mengambil data sekolah \nsesuai dengan satuan pendidikan & provinsi yang diinginkan 👾")),

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

	// err := form.Run()
	// if err != nil {
	// 	fmt.Println("Yah error:", err)
	// 	os.Exit(1)
	// }

	return form
}
