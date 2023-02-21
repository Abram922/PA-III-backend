package user

import "golang.org/x/crypto/bcrypt"

type Service interface {
	Daftar(input DaftarField) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Daftar(input DaftarField) (User, error) {
	user := User{}
	user.Email = input.Email

	user.Nama = input.Nama

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)

	if err != nil {
		return user, err
	}

	user.Password = string(passwordHash)
	user.Role = 2
	newUser, err := s.repository.Register(user)

	if err != nil {
		return newUser, err
	}

	return newUser, nil

}
