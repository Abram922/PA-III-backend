package komponen

type KomponenPAKFormatter struct {
	ID        int    `json:"id"`
	Judul     string `json:"judul"`
	Deskripsi string `json:"deskripsi"`
	Token     string `json:"token"`
}

func FormatKompoenenPAK(komponen Entity_pak, token string) KomponenPAKFormatter {
	formatter := KomponenPAKFormatter{
		ID:        komponen.ID,
		Judul:     komponen.Judul,
		Deskripsi: komponen.Deskripsi,
		Token:     token,
	}

	return formatter
}
