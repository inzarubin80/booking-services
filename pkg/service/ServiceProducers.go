package service

import (
	"github.com/inzarubin80/booking-services"
	"github.com/inzarubin80/booking-services/pkg/repository"
)

type ServiceProducersService  struct {
	repo repository.ServiceProducers
}

func NewServiceProducersService(repo repository.ServiceProducers) *ServiceProducersService {
	return &ServiceProducersService{repo: repo}
}

func (s *ServiceProducersService) Create(item booking.ServiceProducers) (int, error) {
	return s.repo.Create(item)
}

func (s *ServiceProducersService) GetAll() ([] booking.ServiceProducers, error) {
	return s.repo.GetAll()
}

func (s *ServiceProducersService) GetById(id int) (booking.ServiceProducers, error) {
	return s.repo.GetById(id)
}

func (s *ServiceProducersService) Delete(id int) error {
	return s.repo.Delete(id, true)
}

func (s *ServiceProducersService) Update(id int, input booking.UpdateServiceProducersInput) error {

	return s.repo.Update(id, input)
}