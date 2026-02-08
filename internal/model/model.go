package model

import (
	"time"

	"gorm.io/datatypes"
)

type Advertisement struct {
	ID               int                              `json:"id" gorm:"primaryKey;autoIncrement"`
	Name             string                           `json:"name" gorm:"size:255;not null"`
	CountryID        int                              `json:"country_id" gorm:"not null"`
	CityID           int                              `json:"city_id" gorm:"not null"`
	CompanyID        int                              `json:"company_id" gorm:"not null"`
	StartDateTime    *time.Time                       `json:"start_date_time" gorm:"not null"`
	EndDateTime      *time.Time                       `json:"end_date_time"`
	ContentType      string                           `json:"content_type" gorm:"type:text"`
	ContentURL       string                           `json:"content_url" gorm:"type:text"`
	ShowCompanyLogo  bool                             `json:"show_company_logo"`
	IsSponsored      bool                             `json:"is_sponsored"`
	Title            string                           `json:"title" gorm:"type:text"`
	Body             string                           `json:"body" gorm:"type:text"`
	Thumbnail        string                           `json:"thumbnail" gorm:"type:text"`
	Deeplink         string                           `json:"deeplink" gorm:"type:text"`
	TargetImpression int                              `json:"target_impression"`
	TargetClick      int                              `json:"target_click"`
	AudienceType     string                           `json:"audience_type" gorm:"size:255;not null"`
	AudienceMeta     datatypes.JSONType[AudienceMeta] `json:"audience_meta" gorm:"type:jsonb"`
	Status           int                              `json:"status"`
	CreatedAt        *time.Time                       `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt        *time.Time                       `json:"updated_at" gorm:"autoUpdateTime"`
}

type AudienceMeta struct {
	Gender   string  `json:"gender"`
	AgeRange string  `json:"age_range"`
	Lat      float64 `json:"lat"`
	Long     float64 `json:"long"`
	Radius   int     `json:"radius"`
}

type Company struct {
	ID        int        `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string     `json:"name" gorm:"size:255"`
	Logo      string     `json:"logo" gorm:"size:255"`
	CreatedAt *time.Time `json:"created_at" gorm:"autoCreateTime"`
}

type Analytics struct {
	ID                *int       `json:"id" gorm:"primaryKey;autoIncrement"`
	CompanyID         *int       `json:"company_id"`
	Date              *time.Time `json:"date" gorm:"type:date"`
	UniqueImpressions *int       `json:"unique_impressions"`
	TotalImpressions  *int       `json:"impressions"`
	UniqueClicks      *int       `json:"unique_clicks"`
	TotalClicks       *int       `json:"clicks"`
}
