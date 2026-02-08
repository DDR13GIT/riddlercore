package request

import (
	"time"
)

type AdvertisementRequest struct {
	ID               int          `json:"id"`
	Name             string       `json:"name" validate:"required,notblank"`
	CountryID        int          `json:"country_id" validate:"required,oneof=1 2"`
	CityID           int          `json:"city_id" validate:"required,gt=0,lt=10"`
	CompanyID        int          `json:"company_id" validate:"required,gt=0"`
	StartDateTime    *time.Time   `json:"start_date_time" validate:"required,future"`
	EndDateTime      *time.Time   `json:"end_date_time" validate:"omitempty,gtfield=StartDateTime"`
	ContentType      string       `json:"content_type" validate:"required,oneof=image video"`
	ContentURL       string       `json:"content_url" validate:"required"`
	ShowCompanyLogo  bool         `json:"show_company_logo" validate:"required"`
	IsSponsored      bool         `json:"is_sponsored" validate:"required"`
	Title            string       `json:"title" validate:"required,notblank,max=255"`
	Body             string       `json:"body" validate:"required,notblank,max=500"`
	Thumbnail        string       `json:"thumbnail" validate:"required,url"`
	Deeplink         string       `json:"deeplink" validate:"required,url"`
	TargetImpression int          `json:"target_impression" validate:"gte=0"`
	TargetClick      int          `json:"target_click" validate:"gte=0"`
	AudienceType     string       `json:"audience_type" validate:"required,oneof=geo global"`
	AudienceMeta     AudienceMeta `json:"audience_meta" validate:"required"`
	Status           int          `json:"status" validate:"oneof=0 1 2 3"`
	CreatedAt        *time.Time   `json:"created_at"`
	UpdatedAt        *time.Time   `json:"updated_at"`
}

type AudienceMeta struct {
	Gender   string  `json:"gender"`
	AgeRange string  `json:"age_range"`
	Lat      float64 `json:"lat"`
	Long     float64 `json:"long"`
	Radius   int     `json:"radius"`
}
