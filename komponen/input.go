package komponen

type CreateKomponenPAK struct {
	Judul     string `json:"judul" binding :"required"`
	Deskripsi string `json:"deskripsi" binding :"required"`
}
