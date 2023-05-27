package komponen

type Service interface {
	Create(input CreateKomponenPAK) (Entity_pak, error)
	FindKomponenPAK(UserID int) ([]Entity_pak, error)
}

type service struct {
	repository Repository
}

func NewServiceKomponenPAK(repository Repository) *service {
	return &service{repository}
}

func (s *service) Create(input CreateKomponenPAK) (Entity_pak, error) {

	komponenpak := Entity_pak{}

	//komponenpak.KomponenPAK = input.
	//komponenpak.Deskripsi = input.Deskripsi

	newKomponenPAK, err := s.repository.Save(komponenpak)

	if err != nil {
		return newKomponenPAK, err
	}

	return newKomponenPAK, nil

}

func (s *service) FindKomponenPAK(userID int) ([]Entity_pak, error) {
	if userID != 0 {
		komponen, err := s.repository.FindByUserID(userID)

		if err != nil {
			return komponen, err
		}

		return komponen, nil
	}

	komponen, err := s.repository.FindAll()
	if err != nil {
		return komponen, err
	}
	return komponen, nil
}
