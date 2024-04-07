package model

type SatuanPendidikan struct {
	Name string
	Path string
	Url  string
}

type Provinsi struct {
	No     string
	Name   string
	Url    string
	Total  string
	Paud   Paud
	Dikdas Dikdas
	Dikmen Dikmen
	Dikti  Dikti
}

type Paud struct {
	TK  string
	KB  string
	TPA string
	SPS string
}

type Dikdas struct {
	SD  string
	SMP string
}

type Dikmen struct {
	SMA string
	SMK string
	SLB string
}

type Dikti struct {
	Akademik   string
	Politeknik string
}

type DataUrl struct {
	Name string
	Url  string
}

type DataSekolah struct {
	IdentitasSekolah IdentitasSekolah
	DokumenPerijinan DokumenPerijinan
	SaranaPrasarana  SaranaPrasarana
	Kontak           Kontak
}

type IdentitasSekolah struct {
	Nama             string
	NPSN             string
	Alamat           string
	DesaKeluarahan   string
	Kecamatan        string
	KabupatenKota    string
	Provinsi         string
	Status           string
	BentukPendidikan string
}

type DokumenPerijinan struct {
	KementerianPembina         string
	Nauangan                   string
	NPYP                       string
	NoSKPendirian              string
	TanggalSKpendirian         string
	NoSKOperasional            string
	TanggalSKOperasional       string
	TanggalUploadSKOperasional string
	Akreditasi                 string
}

type SaranaPrasarana struct {
	LuasTanah      string
	AksesInternet1 string
	AksesInternet2 string
	SumberListrik  string
}

type Kontak struct {
	Fax     string
	Telepon string
	Email   string
	Website string
}

// FORM
type Options struct {
	SatuanPendidikan SatuanPendidikan
	Provinsi         Provinsi
	ExportTo         string
}
