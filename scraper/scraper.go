package scraper

import (
	"fmt"
	"gin-dapodik/model"
	"strings"

	"github.com/gocolly/colly"
)

var base_url = "https://referensi.data.kemdikbud.go.id"

func GetListSatuanPendidikan() []model.SatuanPendidikan {
	c := colly.NewCollector()

	listSatuanPendidikan := make([]model.SatuanPendidikan, 0)

	// Callback
	c.OnHTML("ul.nav.navbar-nav > li:nth-child(2) > ul.dropdown-menu > li:first-child > ul", func(h *colly.HTMLElement) {
		h.ForEach("li", func(i int, e *colly.HTMLElement) {
			item := model.SatuanPendidikan{}
			item.Name = e.Text
			item.Path = strings.Split(e.ChildAttr("a", "href"), "/pendidikan/")[1]
			item.Url = e.ChildAttr("a", "href")
			listSatuanPendidikan = append(listSatuanPendidikan, item)
		})
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Got an error :", err)
	})

	c.Visit(base_url)

	return listSatuanPendidikan
}

func GetListProvinsiBySatuanPendidikan(satuanPendidikan model.SatuanPendidikan) []model.Provinsi {
	c := colly.NewCollector()
	listProvinsi := make([]model.Provinsi, 0)

	// Callback
	c.OnHTML("table tbody", func(h *colly.HTMLElement) {
		h.ForEach("tr", func(i int, e *colly.HTMLElement) {
			item := model.Provinsi{}
			item.No = e.ChildText("td:nth-child(1)")
			item.Name = e.ChildText("td:nth-child(2)")
			item.Url = e.ChildAttr("td:nth-child(2) a", "href")

			switch satuanPendidikan.Path {
			case "paud":
				item.Paud.TK = e.ChildText("td:nth-child(3)")
				item.Paud.KB = e.ChildText("td:nth-child(4)")
				item.Paud.TPA = e.ChildText("td:nth-child(5)")
				item.Paud.SPS = e.ChildText("td:nth-child(6)")
				item.Total = e.ChildText("td:nth-child(7)")

			case "dikdas":
				item.Dikdas.SD = e.ChildText("td:nth-child(3)")
				item.Dikdas.SMP = e.ChildText("td:nth-child(4)")
				item.Total = e.ChildText("td:nth-child(5)")

			case "dikmen":
				item.Dikmen.SMA = e.ChildText("td:nth-child(3)")
				item.Dikmen.SMK = e.ChildText("td:nth-child(4)")
				item.Dikmen.SLB = e.ChildText("td:nth-child(5)")
				item.Total = e.ChildText("td:nth-child(6)")

			case "dikti":
				item.Dikti.Akademik = e.ChildText("td:nth-child(3)")
				item.Dikti.Politeknik = e.ChildText("td:nth-child(4)")
			}
			listProvinsi = append(listProvinsi, item)
		})
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Got an error :", err)
	})

	c.Visit(satuanPendidikan.Url)

	return listProvinsi
}

func GetListDataUrl(url string) []model.DataUrl {
	c := colly.NewCollector()
	listDataUrl := make([]model.DataUrl, 0)

	// Callback
	c.OnHTML("table tbody", func(h *colly.HTMLElement) {
		h.ForEach("tr", func(i int, e *colly.HTMLElement) {
			dataUrl := model.DataUrl{}
			dataUrl.Name = e.ChildText("td:nth-child(2)")
			dataUrl.Url = e.ChildAttr("td:nth-child(2) a", "href")
			listDataUrl = append(listDataUrl, dataUrl)
		})
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Got an error :", err)
	})

	c.Visit(url)

	return listDataUrl
}

func GetDataSekolah(url string) model.DataSekolah {
	c := colly.NewCollector()
	item := model.DataSekolah{}

	// Callback
	c.OnHTML("div.tabs", func(h *colly.HTMLElement) {
		// Get Identitas Pendidikan
		h.ForEach("div.tabby-tab:nth-child(1) > .tabby-content > table > tbody", func(i int, e *colly.HTMLElement) {
			item.IdentitasSekolah.Nama = e.ChildText("tr:nth-child(1) td:nth-child(4)")
			item.IdentitasSekolah.NPSN = e.ChildText("tr:nth-child(2) td:nth-child(4)")
			item.IdentitasSekolah.Alamat = e.ChildText("tr:nth-child(3) td:nth-child(4)")
			item.IdentitasSekolah.DesaKeluarahan = e.ChildText("tr:nth-child(5) td:nth-child(4)")
			item.IdentitasSekolah.Kecamatan = e.ChildText("tr:nth-child(6) td:nth-child(4)")
			item.IdentitasSekolah.KabupatenKota = e.ChildText("tr:nth-child(7) td:nth-child(4)")
			item.IdentitasSekolah.Provinsi = e.ChildText("tr:nth-child(8) td:nth-child(4)")
			item.IdentitasSekolah.Status = e.ChildText("tr:nth-child(9) td:nth-child(4)")
			item.IdentitasSekolah.BentukPendidikan = e.ChildText("tr:nth-child(10) td:nth-child(4)")
		})

		// Get Dokumen Perijinan
		h.ForEach("div.tabby-tab:nth-child(2) > .tabby-content > table > tbody", func(i int, e *colly.HTMLElement) {
			item.DokumenPerijinan.KementerianPembina = e.ChildText("tr:nth-child(1) td:nth-child(4)")
			item.DokumenPerijinan.Nauangan = e.ChildText("tr:nth-child(2) td:nth-child(4)")
			item.DokumenPerijinan.NPYP = e.ChildText("tr:nth-child(3) td:nth-child(4)")
			item.DokumenPerijinan.NoSKPendirian = e.ChildText("tr:nth-child(4) td:nth-child(4)")
			item.DokumenPerijinan.TanggalSKpendirian = e.ChildText("tr:nth-child(5) td:nth-child(4)")
			item.DokumenPerijinan.NoSKOperasional = e.ChildText("tr:nth-child(6) td:nth-child(4)")
			item.DokumenPerijinan.TanggalSKOperasional = e.ChildText("tr:nth-child(7) td:nth-child(4)")
			item.DokumenPerijinan.TanggalUploadSKOperasional = e.ChildText("tr:nth-child(9) td:nth-child(4)")
			item.DokumenPerijinan.Akreditasi = e.ChildText("tr:nth-child(10) td:nth-child(4)")
		})

		// Get Sarana Prasarana
		h.ForEach("div.tabby-tab:nth-child(3) > .tabby-content > table > tbody", func(i int, e *colly.HTMLElement) {
			item.SaranaPrasarana.LuasTanah = e.ChildText("tr:nth-child(1) td:nth-child(4)")
			item.SaranaPrasarana.AksesInternet1 = e.ChildText("tr:nth-child(2) td:nth-child(4)")
			item.SaranaPrasarana.AksesInternet2 = e.ChildText("tr:nth-child(3) td:nth-child(4)")
			item.SaranaPrasarana.SumberListrik = e.ChildText("tr:nth-child(4) td:nth-child(4)")
		})

		// Get Kontak
		h.ForEach("div.tabby-tab:nth-child(4) > .tabby-content > table > tbody", func(i int, e *colly.HTMLElement) {
			item.Kontak.Fax = e.ChildText("tr:nth-child(1) td:nth-child(4)")
			item.Kontak.Telepon = e.ChildText("tr:nth-child(2) td:nth-child(4)")
			item.Kontak.Email = e.ChildText("tr:nth-child(3) td:nth-child(4)")
			item.Kontak.Website = e.ChildText("tr:nth-child(4) td:nth-child(4)")
		})
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Got an error :", err)
	})

	c.Visit(url)

	return item
}
