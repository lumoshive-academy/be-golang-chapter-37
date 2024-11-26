package service

import (
	"golang-chapter-37/implem-gin-shipping/repository"

	"go.uber.org/zap"
)

type AllService struct {
	CustomerService ShippingServiceInterface
}

func NewAllService(repo repository.AllRepository, log *zap.Logger) AllService {
	return AllService{
		CustomerService: NewShippingService(repo, log),
	}
}
