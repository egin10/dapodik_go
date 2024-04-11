package form

import (
	"fmt"
	"gin-dapodik/model"

	"github.com/charmbracelet/huh"
)

func FormConfirm(optionValue *model.OptionValue, accessible bool) *huh.Form {
	form := huh.NewForm(
		huh.NewGroup(
			// Konfirmasi
			huh.NewConfirm().
				Title(fmt.Sprintf("Apakah yakin dengan data satuan pendidikan %s di %s?", optionValue.SatuanPendidikan.Path, optionValue.Provinsi.Name)).
				Affirmative("YA!").
				Negative("TIDAK"),
		),
	).WithAccessible(accessible)

	return form
}
