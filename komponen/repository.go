package komponen

import "gorm.io/gorm"

type Repository interface {
	Save(komponen Entity_pak) (Entity_pak, error)
}

type repository struct { // jika huruf kecil maka struct hanya dapat digunkaan pada file ini
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(komponen Entity_pak) (Entity_pak, error) {
	err := r.db.Create(&komponen).Error //

	if err != nil {
		return komponen, err
	}

	return komponen, nil
}
