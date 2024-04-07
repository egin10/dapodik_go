package form

import (
	"fmt"
	"gin-dapodik/model"
	"gin-dapodik/scraper"
	"os"

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

func FormSatuanPendidikan(options *model.Options, accessible bool) {
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
				Value(&options.SatuanPendidikan),
		),
	).WithAccessible(accessible)

	errOne := form.Run()
	if errOne != nil {
		fmt.Println("Yah error:", errOne)
		os.Exit(1)
	}
}
