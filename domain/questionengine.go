package domain

import (
	"context"
	"errors"

	"ddr13/riddlercore/internal/model"
)

type AdvertisementRepository interface {
	CreateAdvertisement(ctx context.Context, advertisement *model.Advertisement) error
}

// AdvertisementUseCase represents Advertisement's usecase contract
type AdvertisementUseCase interface {
	CreateAdvertisement(ctx context.Context, req *model.Advertisement) error
}

var (
	ErrAdvertisementNotFound = errors.New("questionengine not found")
	ErrAdvertisementConflict = errors.New("questionengine conflict")
)
