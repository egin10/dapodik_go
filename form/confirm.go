package form

import (
	"fmt"
	"gin-dapodik/model"
	"os"

	"github.com/charmbracelet/huh"
)

func FormConfirm(options *model.Options, accessible bool) {
	form := huh.NewForm(
		huh.NewGroup(
			// Konfirmasi
			huh.NewConfirm().
				Title("Apakah sudah yakin dengan pilihan anda?").
				Affirmative("YA!").
				Negative("TIDAK"),
		),
	).WithAccessible(accessible)

	errTwo := form.Run()
	if errTwo != nil {
		fmt.Println("Yah error:", errTwo)
		os.Exit(1)
	}
}
