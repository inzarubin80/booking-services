package service

import (
	"github.com/inzarubin80/booking-services"
	"github.com/inzarubin80/booking-services/pkg/repository"
)

type ServiceItemsService struct {
	repo repository.ServiceItems
}

func NewServiceItemsService(repo repository.ServiceItems) *ServiceItemsService {
	return &ServiceItemsService{repo: repo}
}

func (s *ServiceItemsService) Create(item booking.ServiceItems) (int, error) {
	return s.repo.Create(item)
}

func (s *ServiceItemsService) GetAll() ([]booking.ServiceItems, error) {
	return s.repo.GetAll()
}

func (s *ServiceItemsService) GetById(id int) (booking.ServiceItems, error) {
	return s.repo.GetById(id)
}

func (s *ServiceItemsService) Delete(id int) error {
	return s.repo.Delete(id, true)
}

func (s *ServiceItemsService) Update(id int, input booking.UpdateServiceItemsInput) error {
	return s.repo.Update(id, input)
}
