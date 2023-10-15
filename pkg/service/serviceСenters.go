package service

import (
	"github.com/inzarubin80/booking-services"
	"github.com/inzarubin80/booking-services/pkg/repository"
)

type ServiceСentersService  struct {
	repo repository.ServiceСenters
}

func NewServiceСentersService(repo repository.ServiceСenters) *ServiceСentersService {
	return &ServiceСentersService{repo: repo}
}

func (s *ServiceСentersService) Create(item booking.ServiceСenters) (int, error) {
	return s.repo.Create(item)
}

func (s *ServiceСentersService) GetAll() ([] booking.ServiceСenters, error) {
	return s.repo.GetAll()
}

func (s *ServiceСentersService) GetById(id int) (booking.ServiceСenters, error) {
	return s.repo.GetById(id)
}

func (s *ServiceСentersService) Delete(id int) error {
	return s.repo.Delete(id, true)
}

func (s *ServiceСentersService) Update(id int, input booking.UpdateServiceСentersInput) error {

	return s.repo.Update(id, input)
}