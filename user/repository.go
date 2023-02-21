package user

import "gorm.io/gorm"

type Repository interface {
	Register(user User) (User, error)
}
type repository struct {
	db *gorm.DB
}

func (r *repository) Register(user User) (User, error){
	err := r.db.Create(&user).Error

	if err != nil{
		return user, err
	}

	return user, nil
}
