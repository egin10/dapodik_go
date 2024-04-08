package form

import (
	"gin-dapodik/model"

	"github.com/charmbracelet/huh"
)

func FormConfirm(optionValue *model.OptionValue, accessible bool) *huh.Form {
	form := huh.NewForm(
		huh.NewGroup(
			// Konfirmasi
			huh.NewConfirm().
				Title("Apakah sudah yakin dengan pilihan anda?").
				Affirmative("YA!").
				Negative("TIDAK"),
		),
	).WithAccessible(accessible)

	// err := form.Run()
	// if err != nil {
	// 	fmt.Println("Yah error:", err)
	// 	os.Exit(1)
	// }

	return form
}
