package komponen

/*

ID                  int
	UserID              int
	KomponenPAK         string
	Kategori            string
	BatasMaksimalDiakui string
	Score               int
*/
type CreateKomponenPAK struct {
	KomponenPAK         string `json:"komponenpak" binding :"required"`
	Kategori            string `json:"kategori" binding :"required"`
	BatasMaksimalDiakui string `json:"batas_maksimal_diakui"`
	Score               int    `json:"score"`
}


