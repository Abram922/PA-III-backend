package komponen

import "gorm.io/gorm"

type Repository interface {
	Save(komponen Entity_pak) (Entity_pak, error)
	FindAll() ([]Entity_pak, error)
	FindByUserID(UserID int) ([]Entity_pak, error)
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

func (r *repository) FindAll() ([]Entity_pak, error) {
	var komponenpaks []Entity_pak
	err := r.db.Find(&komponenpaks).Error

	//err := r.db.Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error

	if err != nil {
		return komponenpaks, err
	}

	return komponenpaks, nil
}

func (r *repository) FindByUserID(UserID int) ([]Entity_pak, error) {
	var komponenpaks []Entity_pak

	err := r.db.Where("user_id = ?", UserID).Find(&komponenpaks).Error
	//err := r.db.Where("user_id = ?", userID).Find(&campaigns).Error

	if err != nil {
		return komponenpaks, err
	}

	return komponenpaks, nil
}
