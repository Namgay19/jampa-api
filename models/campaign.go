package models

import (
	"time"

	"gorm.io/gorm"
)

type Campaign struct {
	Model
	Header string `json:"header" gorm:"not null"`
	SubHeader string `json:"sub_header" gorm:"not null"` 
	Description string `json:"description" gorm:"not null"`
	Category string `json:"category" gorm:"not null"`
	EndDate time.Time `json:"end_date" gorm:"not null"`
	Image Image `gorm:"polymorphic:Owner;"`
	Status string `json:"status" gorm:"default:active;not null"`
	TargetAmount int `json:"target_amount" gorm:"not null"`
	CollectedAmount int `json:"collected_amount" gorm:"default:0"`
	DonationCount int `json:"donation_count" gorm:"default:0"`
	Donations []Donation
}

func FetchCampaigns(status, query, page, pageSize string) []Campaign {
	var campaigns []Campaign
	DB.Scopes(FilterByStatus(status), FetchByQuery(query), Paginate(page, pageSize)).
	Order("created_at desc").
	Preload("Image").
	Find(&campaigns) 

	return campaigns
}

func FilterByStatus(status string) func(db *gorm.DB) *gorm.DB {
	return func (db *gorm.DB) *gorm.DB {
		return db.Where(&Campaign{Status: status})
	}
}

func FetchByQuery(query string) func(db *gorm.DB) *gorm.DB {
	return func (db *gorm.DB) *gorm.DB {
		return db.Where("LOWER(header) like '%' || ? || '%' OR LOWER(sub_header) like '%' || ? || '%'", query, query)
	}
}
