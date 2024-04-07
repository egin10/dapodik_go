package form

import (
	"fmt"
	"gin-dapodik/model"
	"os"

	"github.com/charmbracelet/huh"
)

func FormExportTo(options *model.Options, accessible bool) {
	form := huh.NewForm(
		huh.NewGroup(
			// Pilih Export
			huh.NewSelect[string]().
				Options(
					huh.NewOption("Excel (.xlsx)", "xlsx"),
					huh.NewOption("JSON (.json)", "json"),
				).
				Title("Ekstrak data menjadi").
				Description("Silahkan pilih hasil file download").
				Value(&options.ExportTo),
		),
	).WithAccessible(accessible)

	errTwo := form.Run()
	if errTwo != nil {
		fmt.Println("Yah error:", errTwo)
		os.Exit(1)
	}
}
