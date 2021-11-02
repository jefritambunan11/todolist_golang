package campaign

import (
	"time"
)

type Campaign struct {
	ID int
	UserID int 
	Name string 
	ShortDescription string 
	Description string 
	Perks string 
	BackerCount int 
	GoalAmount int
	CurrenctAmount int 
	Slug string
	CreatedAt time.Time  
	UpdatedAt time.Time  

	// Relasi Struct 
	CampaignImages []CampaignImage
}

type CampaignImage struct {
	ID int
	CampaignID int 
	FileName string 
	IsPrimary int 	
	CreatedAt time.Time  
	UpdatedAt time.Time  
}