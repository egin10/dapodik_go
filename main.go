package main

import (
	"fmt"
	"gin-dapodik/scraper"
	"time"
)

func timeTrack(start time.Time) {
	duration := time.Since(start).Seconds()
	fmt.Printf("Durasi : %.2fs\n", duration)
}

func main() {
	defer timeTrack(time.Now())

	// Get Satuan Pendidikan
	listSatuanPendidikan := scraper.GetListSatuanPendidikan()
	// for index, value := range listSatuanPendidikan {
	// 	fmt.Printf("%d) %s \n", index+1, value.Name)
	// }

	// Pilih Satuan Pendidikan
	/**
	*	0 : paud
	*	1 : dikdas
	*	2 : dikmen
	*	3 : dikti
	*	4 : dikmas
	 */
	satuanPendidikan := listSatuanPendidikan[2]

	// Get List Provinsi by Satuan Pendidikan terpilih
	listProvinsi := scraper.GetListProvinsiBySatuanPendidikan(satuanPendidikan)
	// for _, value := range listProvinsi {
	// 	switch satuanPendidikan.Path {
	// 	case "paud":
	// 		fmt.Printf("%s) | %s \n", value.No, value.Name)
	// 		fmt.Printf("\tTK: %s |KB: %s |TPA: %s |SPS: %s |Total: %s \n", value.Paud.TK, value.Paud.KB, value.Paud.TPA, value.Paud.SPS, value.Total)
	// 	case "dikdas":
	// 		fmt.Printf("%s) | %s \n", value.No, value.Name)
	// 		fmt.Printf("\tSD: %s |SMP: %s |Total: %s \n", value.Dikdas.SD, value.Dikdas.SMP, value.Total)
	// 	case "dikmen":
	// 		fmt.Printf("%s) | %s \n", value.No, value.Name)
	// 		fmt.Printf("\tSMA: %s |SMK: %s |SLB: %s |Total: %s \n", value.Dikmen.SMA, value.Dikmen.SMK, value.Dikmen.SLB, value.Total)
	// 	case "dikti":
	// 		fmt.Printf("%s) | %s \n", value.No, value.Name)
	// 		fmt.Printf("\tAkademik: %s |Politeknik: %s \n", value.Dikti.Akademik, value.Dikti.Politeknik)
	// 	}
	// 	fmt.Println("=======================================")
	// }

	// Pilih Provinsi
	provinsi := listProvinsi[0]

	// Get Semua Data Sekolah di Provinsi terpilih
	// Get List Url Kabupaten/Kota
	listUrlKabKota := scraper.GetListDataUrl(provinsi.Url)
	for _, urlKabKota := range listUrlKabKota {
		fmt.Printf("\tKAB.KOTA : %s\n", urlKabKota.Name)

		// Get Url Kecamatan
		listUrlKecamatan := scraper.GetListDataUrl(urlKabKota.Url)
		for _, urlKecamatan := range listUrlKecamatan {
			fmt.Printf("\tKECAMATAN : %s\n", urlKecamatan.Name)

			// Get Url Sekolah
			listUrlSekolah := scraper.GetListDataUrl(urlKecamatan.Url)
			for _, urlSekolah := range listUrlSekolah {
				// Get Data Detail Sekolah
				dataSekolah := scraper.GetDataSekolah(urlSekolah.Url)
				fmt.Printf("Nama: %s\n", dataSekolah.IdentitasSekolah.Nama)
				fmt.Printf("NPSN: %s\n", dataSekolah.IdentitasSekolah.NPSN)
				fmt.Printf("Alamat: %s\n", dataSekolah.IdentitasSekolah.Alamat)
				fmt.Printf("BentukPendidikan: %s\n", dataSekolah.IdentitasSekolah.BentukPendidikan)
				fmt.Println("--------------------------------------")
			}
		}
		fmt.Println("=======================================")
	}

	fmt.Println("Finished 👾")
}
