package user

type User struct {
	ID                      int
	Email                   string
	Password                string
	Nama                    string
	JabatanTerakhir        string
	NIP                     int
	PendidikanTerakhir     string
	MataKuliahyangDiampu string
	BidangIlmu           string
	Role                 int
	AvatarFileName       string
}

type Role struct {
	ID   int
	Role string
}
