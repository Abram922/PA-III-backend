package user

type User struct {
	ID                          int
	Email                   string
	Password                string
	Nama                    string
	Jabatan_Terakhir        string
	NIP                     int
	Pendidikan_Terakhir     string
	Mata_Kuliah_yang_Diampu string
	Bidang_Ilmu             string
	Role                    int
}

type Role struct {
	ID   int
	Role string
}
