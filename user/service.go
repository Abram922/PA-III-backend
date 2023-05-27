package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Daftar(input DaftarField) (User, error)
	LoginUser(input LoginField) (User, error)
	IsEmailAvailable(input CheckEmailInput) (bool, error)
	SaveAvatar(ID int, fileLocation string) (User, error)
	GetUserbyID(ID int) (User, error)
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

func (s *service) LoginUser(input LoginField) (User, error) {
	email := input.Email
	pasword := input.Password

	user, err := s.repository.Login(email)

	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("Tolong periksa email atau katasandi anda")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pasword))

	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) IsEmailAvailable(input CheckEmailInput) (bool, error) {
	email := input.Email

	user, err := s.repository.FindByEmail(email)

	if err != nil {
		return false, err
	}

	if user.ID == 0 {
		return true, nil
	}

	return false, nil

}

func (s *service) SaveAvatar(ID int, fileLocation string) (User, error) {
	//dapatkan user berdasarkan ID
	// update attribute avatar file name
	// simpan perubahan avatar file name

	user, err := s.repository.FindByID(ID)

	if err != nil {
		return user, err
	}

	user.AvatarFileName = fileLocation

	updateUser, err := s.repository.Update(user)

	if err != nil {
		return updateUser, err
	}

	return updateUser, nil
}

func (s *service) GetUserbyID(ID int) (User, error) {
	user, err := s.repository.FindByID(ID)

	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("No user found on with that ID")
	}

	return user, nil
}
