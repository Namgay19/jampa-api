package controllers

import (
	"namgay/jampa/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateDonation(c *gin.Context) {
	var input DonationInput
	campaignId, _ := strconv.Atoi(c.Param("id"));
	
	if err:= c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var donation = models.Donation{
		CampaignId: campaignId,
		FirstName: input.FirstName,
		LastName: input.LastName,
		Email: input.Email,
		PaymentMethod: input.PaymentMethod,
		Amount: input.Amount,
		Bank: input.Bank,
		AccountNumber: input.AccountNumber,
		CardNumber: input.CardNumber,
		Cvv: input.Cvv,
		ExpiryDate: input.ExpiryDate,
	}

	result := models.DB.Create(&donation);

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": donation.ID })
}

func GetDonations(c *gin.Context) {
	sortBy := c.Query("sortBy")
	page := c.Query("page")
	pageSize := c.Query("pageSize")
	campaignId, _ := strconv.Atoi(c.Param("id"));

	donations := models.FetchDonations(campaignId, sortBy, page, pageSize)

	c.IndentedJSON(http.StatusOK, gin.H{"data": donations})
}

type DonationInput struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email" binding:"email"`
	PaymentMethod string `json:"payment_method" binding:"required,oneof=national international"`
	Amount int `json:"amount" binding:"required"`
	Bank string `json:"bank"`
	AccountNumber string `json:"account_number"`
	CardNumber string `json:"card_number"`
	Cvv string `json:"cvv"`
	ExpiryDate string `json:"expiry_time"`
}
