package http

import (
	"ddr13/riddlercore/domain"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"

	"net/http"
)

type AdvertisementHandler struct {
	advertisementUsecase domain.AdvertisementUseCase
	validator            *validator.Validate
}

func NewHTTPHandler(
	r *chi.Mux,
	usecase domain.AdvertisementUseCase,
	validator *validator.Validate,
) {
	handler := &AdvertisementHandler{
		advertisementUsecase: usecase,
		validator:            validator,
	}
	r.Route("/v1/advertisements", func(r chi.Router) {
		r.Post("/", handler.CreateAdvertisement)

	})
}

func (h *AdvertisementHandler) CreateAdvertisement(w http.ResponseWriter, r *http.Request) {
}
