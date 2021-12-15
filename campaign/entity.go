package campaign

import (
	"time"
	"bwastartup/user"
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
	CurrentAmount int 
	Slug string
	CreatedAt time.Time  
	UpdatedAt time.Time  

	// Relasi Struct 
	CampaignImages []CampaignImage

	User user.User
}

type CampaignImage struct {
	ID int
	CampaignID int 
	FileName string 
	IsPrimary int 	
	CreatedAt time.Time  
	UpdatedAt time.Time  
}