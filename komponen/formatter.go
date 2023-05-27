package komponen

/*
	ID             int
	KomponenPAK    string
	Kategori       string
	Batas_Maksimal string
	Score          int
*/

type KomponenPAKFormatter struct {
	ID                  int    `json:"id"`
	KomponenPAK         string `json:"komponenpak"`
	Kategori            string `json:"kategori"`
	BatasMaksimalDiakui string `json:"batas_maksimal_diakui"`
	Score               int    `json:"score"`
}

func FormatKompoenenPAK(komponen Entity_pak) KomponenPAKFormatter {
	formatter := KomponenPAKFormatter{
		ID:                  komponen.ID,
		KomponenPAK:         komponen.KomponenPAK,
		Kategori:            komponen.Kategori,
		BatasMaksimalDiakui: komponen.BatasMaksimalDiakui,
		Score:               komponen.Score,
	}

	return formatter
}
