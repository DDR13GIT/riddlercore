package repository

import (
	"context"
	"log/slog"

	"ddr13/riddlercore/domain"
	"ddr13/riddlercore/internal/model"

	"gorm.io/gorm"
)

func New(db *gorm.DB) domain.AdvertisementRepository {
	return AdvertisementSqlStorage{db: db}
}

type AdvertisementSqlStorage struct {
	db *gorm.DB
}

func (a AdvertisementSqlStorage) CreateAdvertisement(ctx context.Context, advertisement *model.Advertisement) error {
	result := a.db.Create(advertisement)
	if result.Error != nil {
		slog.Info("Advertisement Repository: Failed to create questionengine", result.Error)
		return result.Error
	}
	return nil
}

func (a AdvertisementSqlStorage) UpdateAdvertisement(ctx context.Context, advertisement *model.Advertisement) error {
	result := a.db.Save(advertisement)
	if result.Error != nil {
		slog.Info("Advertisement Repository: Failed to update questionengine", result.Error)
		return result.Error
	}
	return nil
}

func (a AdvertisementSqlStorage) FetchOne(ctx context.Context, id int) (*model.Advertisement, error) {
	var ad model.Advertisement
	result := a.db.First(&ad, id)
	if result.Error != nil {
		slog.Info("Advertisement Repository: Failed to fetch questionengine", result.Error)
		return nil, result.Error
	}
	return &ad, nil
}
