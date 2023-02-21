package komponen

type Service interface {
	Create(input CreateKomponenPAK) (Entity_pak, error)
}

type service struct {
	repository Repository
}

func NewServiceKomponenPAK(repository Repository) *service {
	return &service{repository}
}

func (s *service) Create(input CreateKomponenPAK) (Entity_pak, error) {

	komponenpak := Entity_pak{}

	komponenpak.Judul = input.Judul
	komponenpak.Deskripsi = input.Deskripsi

	newKomponenPAK, err := s.repository.Save(komponenpak)

	if err != nil {
		return newKomponenPAK, err
	}

	return newKomponenPAK, nil

}
