package user

import "gorm.io/gorm"

type Repository interface {
	Register(user User) (User, error)
	Login(email string) (User, error)
	FindByEmail(email string) (User, error)
	FindByID(ID int) (User, error)
	Update(user User) (User, error)
}
type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
	// jika kita memanggil newrepository maka kita akan membuat objek baru
	// dari repository(atas huruf keci) dan nilai dbnya akan diisi
	// sesuai denan parameter NewRepository
}
func (r *repository) Register(user User) (User, error) {
	err := r.db.Create(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) Login(email string) (User, error) {
	var user User

	err := r.db.Where("email = ?", email).Find(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindByEmail(email string) (User, error) {
	var user User
	err := r.db.Table("users").Where("Email = ?", email).Scan(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil

}

func (r *repository) FindByID(ID int) (User, error) {
	var user User

	err := r.db.Where("ID = ?", ID).Find(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) Update(user User) (User, error) {
	err := r.db.Save(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}
