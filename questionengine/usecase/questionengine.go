package usecase

import (
	"context"
	"ddr13/riddlercore/domain"
	"ddr13/riddlercore/internal/model"
	"log/slog"
)

const (
	AdDraftState     = 0
	AdActiveState    = 1
	AdPausedState    = 2
	AdCompletedState = 3
)

func New(repo domain.AdvertisementRepository) domain.AdvertisementUseCase {
	return &AdvertisementUsecase{
		repo: repo,
	}
}

type AdvertisementUsecase struct {
	repo domain.AdvertisementRepository
}

func (a AdvertisementUsecase) CreateAdvertisement(ctx context.Context, req *model.Advertisement) error {
	err := a.repo.CreateAdvertisement(ctx, req)
	if err != nil {
		slog.Info("Advertisement Usecase: Failed to create questionengine", err)
		return err
	}
	return nil
}
