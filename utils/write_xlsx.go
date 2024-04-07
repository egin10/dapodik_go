package utils

import (
	"fmt"
	"gin-dapodik/model"
	"strconv"
	"strings"
	"time"

	"github.com/xuri/excelize/v2"
)

func WireToExcel(data []model.DataSekolah, satuanPendidikan model.SatuanPendidikan, provinsi model.Provinsi) error {
	f := excelize.NewFile()

	// Header
	f.SetCellValue("Sheet1", "A1", "No")
	f.SetCellValue("Sheet1", "B1", "Nama")
	f.SetCellValue("Sheet1", "C1", "NPSN")
	f.SetCellValue("Sheet1", "D1", "Alamat")
	f.SetCellValue("Sheet1", "E1", "Desa/Kelurahan")
	f.SetCellValue("Sheet1", "F1", "Kecamatan/Kota (LN)")
	f.SetCellValue("Sheet1", "G1", "Kab.-Kota/Negara (LN)")
	f.SetCellValue("Sheet1", "H1", "Propinsi/Luar Negeri (LN)")
	f.SetCellValue("Sheet1", "I1", "Status Sekolah")
	f.SetCellValue("Sheet1", "J1", "Bentuk Pendidikan")
	f.SetCellValue("Sheet1", "K1", "Kementerian Pembina")
	f.SetCellValue("Sheet1", "L1", "Naungan")
	f.SetCellValue("Sheet1", "M1", "NPYP")
	f.SetCellValue("Sheet1", "N1", "No. SK. Pendirian")
	f.SetCellValue("Sheet1", "O1", "Tanggal SK. Pendirian")
	f.SetCellValue("Sheet1", "P1", "Nomor SK Operasional")
	f.SetCellValue("Sheet1", "Q1", "Tanggal SK Operasional")
	f.SetCellValue("Sheet1", "R1", "Tanggal Upload SK Op.")
	f.SetCellValue("Sheet1", "S1", "Akreditasi")
	f.SetCellValue("Sheet1", "T1", "Luas Tanah")
	f.SetCellValue("Sheet1", "U1", "Akses Internet 1")
	f.SetCellValue("Sheet1", "V1", "Akses Internet 2")
	f.SetCellValue("Sheet1", "W1", "Sumber Listrik")
	f.SetCellValue("Sheet1", "X1", "Fax")
	f.SetCellValue("Sheet1", "Y1", "Telepon")
	f.SetCellValue("Sheet1", "Z1", "Email")
	f.SetCellValue("Sheet1", "AA1", "Website")

	for i, value := range data {
		f.SetCellValue("Sheet1", "A"+strconv.Itoa(i+2), i+1)
		f.SetCellValue("Sheet1", "B"+strconv.Itoa(i+2), value.IdentitasSekolah.Nama)
		f.SetCellValue("Sheet1", "C"+strconv.Itoa(i+2), value.IdentitasSekolah.NPSN)
		f.SetCellValue("Sheet1", "D"+strconv.Itoa(i+2), value.IdentitasSekolah.Alamat)
		f.SetCellValue("Sheet1", "E"+strconv.Itoa(i+2), value.IdentitasSekolah.DesaKeluarahan)
		f.SetCellValue("Sheet1", "F"+strconv.Itoa(i+2), value.IdentitasSekolah.Kecamatan)
		f.SetCellValue("Sheet1", "G"+strconv.Itoa(i+2), value.IdentitasSekolah.KabupatenKota)
		f.SetCellValue("Sheet1", "H"+strconv.Itoa(i+2), value.IdentitasSekolah.Provinsi)
		f.SetCellValue("Sheet1", "I"+strconv.Itoa(i+2), value.IdentitasSekolah.Status)
		f.SetCellValue("Sheet1", "J"+strconv.Itoa(i+2), value.IdentitasSekolah.BentukPendidikan)
		f.SetCellValue("Sheet1", "K"+strconv.Itoa(i+2), value.DokumenPerijinan.KementerianPembina)
		f.SetCellValue("Sheet1", "L"+strconv.Itoa(i+2), value.DokumenPerijinan.Nauangan)
		f.SetCellValue("Sheet1", "M"+strconv.Itoa(i+2), value.DokumenPerijinan.NPYP)
		f.SetCellValue("Sheet1", "N"+strconv.Itoa(i+2), value.DokumenPerijinan.NoSKPendirian)
		f.SetCellValue("Sheet1", "O"+strconv.Itoa(i+2), value.DokumenPerijinan.TanggalSKpendirian)
		f.SetCellValue("Sheet1", "P"+strconv.Itoa(i+2), value.DokumenPerijinan.NoSKOperasional)
		f.SetCellValue("Sheet1", "Q"+strconv.Itoa(i+2), value.DokumenPerijinan.TanggalSKOperasional)
		f.SetCellValue("Sheet1", "R"+strconv.Itoa(i+2), value.DokumenPerijinan.TanggalUploadSKOperasional)
		f.SetCellValue("Sheet1", "S"+strconv.Itoa(i+2), value.DokumenPerijinan.Akreditasi)
		f.SetCellValue("Sheet1", "T"+strconv.Itoa(i+2), value.SaranaPrasarana.LuasTanah)
		f.SetCellValue("Sheet1", "U"+strconv.Itoa(i+2), value.SaranaPrasarana.AksesInternet1)
		f.SetCellValue("Sheet1", "V"+strconv.Itoa(i+2), value.SaranaPrasarana.AksesInternet2)
		f.SetCellValue("Sheet1", "W"+strconv.Itoa(i+2), value.SaranaPrasarana.SumberListrik)
		f.SetCellValue("Sheet1", "X"+strconv.Itoa(i+2), value.Kontak.Fax)
		f.SetCellValue("Sheet1", "Y"+strconv.Itoa(i+2), value.Kontak.Telepon)
		f.SetCellValue("Sheet1", "Z"+strconv.Itoa(i+2), value.Kontak.Email)
		f.SetCellValue("Sheet1", "AA"+strconv.Itoa(i+2), value.Kontak.Website)
	}

	now := time.Now()
	filename := fmt.Sprintf("data_sekolah_%s_%s_%s.xlsx", strings.ToLower(satuanPendidikan.Name), strings.ToLower(provinsi.Name), now.Format("2006_01_02"))

	if err := f.SaveAs(filename); err != nil {
		return err
	}
	return nil
}
