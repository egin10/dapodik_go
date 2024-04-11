package form

import (
	"gin-dapodik/model"

	"github.com/charmbracelet/huh"
)

func FormExportTo(optionValue *model.OptionValue, accessible bool) *huh.Form {
	form := huh.NewForm(
		huh.NewGroup(
			// Pilih Export
			huh.NewSelect[string]().
				Options(
					huh.NewOption("Excel (.xlsx)", "xlsx"),
					huh.NewOption("JSON (.json)", "json"),
				).
				Title("Ekstrak data").
				Key("export_to").
				Description("Silahkan pilih hasil file download").
				Value(&optionValue.ExportTo),
		),
	).WithAccessible(accessible)

	return form
}
