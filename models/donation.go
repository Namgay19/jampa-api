package models

import (
	"errors"

	"gorm.io/gorm"
)

type Donation struct {
	Model
	FirstName     string `json:"first_name" gorm:"type:varchar(20);not null"`
	LastName      string `json:"last_name" gorm:"type:varchar(20)"`
	Email         string `json:"email"`
	PaymentMethod string `json:"payment_method" gorm:"type:varchar(25);not null"`
	Amount        int    `json:"amount" gorm:"not null"`
	Bank          string `json:"bank" gorm:"type:varchar(25)"`
	AccountNumber string `json:"account_number" gorm:"type:varchar(25)"`
	CardNumber    string `json:"card_number" gorm:"type:varchar(25)"`
	Cvv           string `json:"cvv" gorm:"type:varchar(5)"`
	ExpiryDate    string `json:"expiry_date"`
	CampaignId    int    `json:"campaign_id"`
	Campaign      Campaign
}

func (d *Donation) BeforeCreate(tx *gorm.DB) (err error) {
	var campaign Campaign
	tx.First(&campaign, d.CampaignId)
	if campaign.Status == "completed" {
		return errors.New("Campaign is completed")
	}
	return nil
}

func (d *Donation) AfterCreate(tx *gorm.DB) (err error) {
	var campaign Campaign
	tx.First(&campaign, d.CampaignId)
	campaign.CollectedAmount = d.Amount + campaign.CollectedAmount
	if (d.Amount + campaign.CollectedAmount) > campaign.TargetAmount {
		campaign.Status = "completed"
		campaign.DonationCount = campaign.DonationCount + 1
	}
	tx.Save(&campaign)

	return
}

func FetchDonations(campaignId int, sortBy, page, pageSize string) []Donation {
	var donations []Donation
	DB.Scopes(FilterByCampaign(campaignId), SortDonations(sortBy), Paginate(page, pageSize)).Preload("Campaign").Find(&donations)

	return donations
}

func FilterByCampaign(campaignId int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(&Donation{CampaignId: campaignId})
	}
}

func SortDonations(sortBy string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if sortBy == "" || sortBy == "Recent" {
			sortBy = "created_at"
		} else {
			sortBy = "amount"
		}
		return db.Order(sortBy + " " + "desc")
	}
}
