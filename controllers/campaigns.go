package controllers

import (
	"log"
	"mime/multipart"
	"namgay/jampa/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func GetCampaigns(c *gin.Context) {
	status := c.Query("status")
	query := c.Query("query")
	page := c.Query("page")
	pageSize := c.Query("pageSize")

	campaigns := models.FetchCampaigns(status, query, page, pageSize)

	c.IndentedJSON(http.StatusOK, gin.H{"data": campaigns})
}

func GetCampaign(c *gin.Context) {
	var campaign models.Campaign

	if err:= models.DB.Where("id = ?", c.Param("id")).Preload("Image").First(&campaign).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, campaign)
}

func CreateCampaign(c *gin.Context) {
	var input createCampaignInput
	
	if err:= c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var uuid = uuid.New().String() + ".jpg"

	var filePath = "./public/images/" + uuid

	if err:= c.SaveUploadedFile(input.Image, filePath); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	campaign := models.Campaign{ 
		Header: input.Header, 
		SubHeader: input.SubHeader, 
		Description: input.Description, 
		EndDate: input.EndDate, 
		Category: input.Category,
		TargetAmount: input.TargetAmount,
		Status: "active",
		Image: models.Image{ ImageUrl: c.Request.Host + "/public/images/" + uuid }, 
	}

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
	if err:= c.ShouldBind(&input); err != nil {
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
	Header string `form:"header" binding:"required,max=50"`
	SubHeader string `form:"sub_header" binding:"required,max=100"`
	Description string `form:"description" binding:"required,max=500"`
	Category string `form:"category" binding:"required"`
	EndDate time.Time `form:"end_date" binding:"CampaignDateValidation" time_format:"2006-01-02"`
	Image *multipart.FileHeader `form:"image" binding:"required"`
	TargetAmount int `form:"target_amount" binding:"required"`
}

type updateCampaignInput struct {
	Header string `form:"header" binding:"max=50"`
	SubHeader string `form:"sub_header" binding:"max=100"`
	Description string `form:"description" binding:"max=500"`
	Category string `form:"category"`
	EndDate time.Time `form:"end_date" binding:"CampaignDateValidation" time_format:"2006-01-02"`
	Image *multipart.FileHeader `form:"image"`
	TargetAmount int `form:"target_amount"`
}

var CampaignDateValidation validator.Func = func(fl validator.FieldLevel) bool { 
	date, ok := fl.Field().Interface().(time.Time)
	log.Println(date)
	if ok {
		today := time.Now()
		if today.After(date) {
			return false 
		}
	}
	return true
}
