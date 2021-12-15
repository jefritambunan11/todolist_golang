package handler

// tangkap parameter di handler
// handler ke service
// service yg menentukan repository yg dipanggil
// repository: GetAll, GetUserByID
// db

import(
	"bwastartup/campaign"	
	"bwastartup/helper"	
// 	"bwastartup/auth"	

	"github.com/gin-gonic/gin"
	
	"net/http"
// 	"fmt"
		"strconv"
)

type campaignHandler struct {
	// campaignService campaign.Service
	// authService auth.Service
	service campaign.Service
}

func NewCampaignHandler(service campaign.Service) *campaignHandler {
	return &campaignHandler {service}
}


func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))
	
	campaigns, err := h.service.GetCampaigns(userID)
		
	if err != nil {				
		//errors := helper.FormatValidationError(err)			
		response := helper.APIResponse("Error to get campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	
	response := helper.APIResponse("List of campaigns", http.StatusOK, "success", campaign.FormatCampaigns(campaigns))
	c.JSON(http.StatusOK, response)
	
}

func (h *campaignHandler) GetCampaign(c *gin.Context) {
	var input campaign.GetCampaignDetailInput 

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	campaign, err := h.service.GetCampaignByID(input)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Campaign detail", http.StatusOK, "success", campaign)
	c.JSON(http.StatusOK, response)
}