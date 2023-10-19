package service

import (
	"github.com/inzarubin80/booking-services"
	"github.com/inzarubin80/booking-services/pkg/repository"
)

type ServiceGroupsService struct {
	repo repository.ServiceGroups
}

func NewServiceGroupsServic(repo repository.ServiceGroups) *ServiceGroupsService {
	return &ServiceGroupsService{repo: repo}
}

func (s *ServiceGroupsService) Create(item booking.ServiceGroups) (int, error) {
	return s.repo.Create(item)
}

func (s *ServiceGroupsService) GetAll() ([] booking.ServiceGroups, error) {
	return s.repo.GetAll()
}

func (s *ServiceGroupsService) GetById(id int) (booking.ServiceGroups, error) {
	return s.repo.GetById(id)
}

func (s *ServiceGroupsService) Delete(id int) error {
	return s.repo.Delete(id, true)
}

func (s *ServiceGroupsService) Update(id int, input booking.UpdateServiceGroupsInput) error {

	return s.repo.Update(id, input)
}