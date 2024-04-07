package utils

import (
	"encoding/json"
	"fmt"
	"gin-dapodik/model"
	"os"
)

func WriteJSON(data []model.DataSekolah) error {
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		fmt.Println("Unable to create json file")
		return err
	}

	_ = os.WriteFile("data_sekolah.json", file, 0644)

	return nil
}
