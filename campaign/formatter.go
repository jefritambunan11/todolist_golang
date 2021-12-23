 package campaign

import (
	"strings"
)


type CampaignFormatter struct {
	ID int `json:"id"`
	UserID int `json:"user_id"`
	Name string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageURL string `json:"image_url"`
	GoalAmount int `json:"goal_amount"`
	CurrentAmount int `json:"current_amount"` 
	Slug string `json:"slug"` 
}

func FormatCampaign(campaign Campaign) CampaignFormatter {
	campaignFormatter := CampaignFormatter{}
	campaignFormatter.ID = campaign.ID
	campaignFormatter.UserID = campaign.UserID
	campaignFormatter.Name = campaign.Name
	campaignFormatter.ShortDescription = campaign.ShortDescription
	campaignFormatter.ImageURL = ""
	campaignFormatter.GoalAmount = campaign.GoalAmount
	campaignFormatter.CurrentAmount = campaign.CurrentAmount
	campaignFormatter.Slug = campaign.Slug

	if len(campaign.CampaignImages) > 0 {
		campaignFormatter.ImageURL = campaign.CampaignImages[0].FileName
	}

	return campaignFormatter
}

func FormatCampaigns(campaigns []Campaign) []CampaignFormatter {
	if len(campaigns) == 0 {
		return []CampaignFormatter{}
	}

	var campaignsFormatter []CampaignFormatter 

	for _, campaign := range campaigns {
		campaignFormatter := FormatCampaign(campaign) 
		campaignsFormatter = append(campaignsFormatter, campaignFormatter)
	}

	return campaignsFormatter
}


type CampaignDetailFormatter struct {
	ID int `json:"id"`
	Name string `json:"name"`
	ShortDescription string `json:"short_description"`
	Description string `json:"description"`
	ImageURL string `json:"image_url"`
	GoalAmount int `json:"goal_amount"`
	CurrentAmount int `json:"current_amount"`
	UserID int `json:"user_id"`
	Slug string `json:"slug"`
	Perks []string `json:"perks"`

	User CampaignUserFormatter `json:"user"`
	Images []CampaignImageFormatter `json:"images"`
}

type CampaignUserFormatter struct {
	Name string `json:"name"`
	ImageURL string `json:"image_url"`
}

type CampaignImageFormatter struct {
	ImageURL  string `json:"image_url"`
	IsPrimary bool   `json:"is_primary"`
}

func FormatCampaignDetail(campaign Campaign) CampaignDetailFormatter {
	c := CampaignDetailFormatter{}
	c.ID = campaign.ID	
	c.Name = campaign.Name
	c.ShortDescription = campaign.ShortDescription
	c.Description = campaign.Description
	c.GoalAmount = campaign.GoalAmount
	c.CurrentAmount = campaign.CurrentAmount
	c.UserID = campaign.UserID
	c.Slug = campaign.Slug
	//c.Perks = campaign.Perks
	
	c.ImageURL = ""

	if len(campaign.CampaignImages) > 0 {
		c.ImageURL = campaign.CampaignImages[0].FileName
	}

	var perks []string 

	for _, perk := range strings.Split(campaign.Perks, ", ") {
		perks = append(perks, perk)
	}

	c.Perks = perks

	// struct user - Package user
	user := campaign.User 
	campaignUserFormatter := CampaignUserFormatter{}
	campaignUserFormatter.Name = user.Name
	campaignUserFormatter.ImageURL = user.AvatarFileName
	c.User = campaignUserFormatter


	// Images
	images := campaign.CampaignImages
	imagesArray := []CampaignImageFormatter{}
	for _, image := range images {
		campaignImageFormatter := CampaignImageFormatter{}
		campaignImageFormatter.ImageURL = image.FileName 

		isPrimary := false
		if image.IsPrimary == 1 {
			isPrimary = true
		}
		campaignImageFormatter.IsPrimary = isPrimary 		

		imagesArray = append(imagesArray, campaignImageFormatter)

	}

	c.Images = imagesArray


	return c
}