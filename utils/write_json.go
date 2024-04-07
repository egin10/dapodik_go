package utils

import (
	"encoding/json"
	"fmt"
	"gin-dapodik/model"
	"os"
	"strings"
	"time"
)

func WriteJSON(data []model.DataSekolah, satuanPendidikan model.SatuanPendidikan, provinsi model.Provinsi) error {
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		fmt.Println("Unable to create json file")
		return err
	}

	now := time.Now()
	filename := fmt.Sprintf("data_sekolah_%s_%s_%s.json", strings.ToLower(satuanPendidikan.Name), strings.ToLower(provinsi.Name), now.Format("2006_01_02"))
	_ = os.WriteFile(filename, file, 0644)

	return nil
}
