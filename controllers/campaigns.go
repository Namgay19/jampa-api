package controllers

import (
	"namgay/jampa/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetCampaigns(c *gin.Context) {
	var campaigns []models.Campaign
	models.DB.Find(&campaigns)

	c.IndentedJSON(http.StatusOK, gin.H{"data": campaigns})
}

func GetCampaign(c *gin.Context) {
	var campaign models.Campaign

	if err:= models.DB.Where("id = ?", c.Param("id")).First(&campaign).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, campaign)
}

func CreateCampaign(c *gin.Context) {
	var input createCampaignInput
	
	if err:= c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	campaign := models.Campaign{ Header: input.Header, SubHeader: input.SubHeader, Description: input.Description, EndDate: input.EndDate, Category: input.Category }
	result := models.DB.Create(&campaign);

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": campaign.ID })
}

func UpdateCampaign(c *gin.Context) {
	var campaign models.Campaign

	if err:= models.DB.Where("id = ?", c.Param("id")).First(&campaign).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}	

	var input updateCampaignInput
	if err:= c.ShouldBindJSON(&input); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	models.DB.Model(&campaign).Updates(map[string]interface{}{"Header": input.Header, "SubHeader": input.SubHeader, "Category": input.Category, "Description": input.Description, "EndDate": input.EndDate})
	c.IndentedJSON(http.StatusOK, gin.H{"data": campaign})
}

func DeleteCampaign(c *gin.Context) {
	var campaign models.Campaign

	if err:= models.DB.Where("id = ?", c.Param("id")).First(&campaign).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Record Not Found"})
		return
	}
	models.DB.Delete(&campaign)
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Campaign deleted"})
}

type createCampaignInput struct {
	Header string `json:"header" binding:"required"`
	SubHeader string `json:"sub_header" binding:"required"`
	Description string `json:"description" binding:"required"`
	Category string `json:"category" binding:"required"`
	EndDate time.Time `json:"end_date"`
}

type updateCampaignInput struct {
	Header string `json:"header"`
	SubHeader string `json:"sub_header"`
	Description string `json:"description"`
	Category string `json:"category"`
	EndDate time.Time `json:"end_date"`
}
