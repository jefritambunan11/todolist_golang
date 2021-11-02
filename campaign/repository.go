package campaign

import (
	"gorm.io/gorm"
)


type Repository interface {
	FindALL() ([]Campaign, error)
	FindByUserID(ID int) ([]Campaign, error)
}

type repository struct {
	db *gorm.DB 
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
} 

func (r *repository) FindALL() ([]Campaign, error) {
	var campaign []Campaign

	err := r.db.Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaign).Error
	if err != nil {
		return campaign, err
	}
	return campaign, nil
}

func (r *repository) FindByUserID(ID int) ([]Campaign, error)	{
	var campaign []Campaign

	err := r.db.Where("user_id = ?", ID).Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaign).Error

	if err != nil {
		return campaign, err
	}
	return campaign, nil	
}