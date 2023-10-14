package service

import (
	"github.com/inzarubin80/booking-services"
	"github.com/inzarubin80/booking-services/pkg/repository"
)

type TypeBusinessService struct {
	repo repository.TypeBusiness
}

func NewTypeBusinessService(repo repository.TypeBusiness) *TypeBusinessService {
	return &TypeBusinessService{repo: repo}
}

func (s *TypeBusinessService) Create(item booking.TypeBusiness) (int, error) {
	return s.repo.Create(item)
}

func (s *TypeBusinessService) GetAll() ([] booking.TypeBusiness, error) {
	return s.repo.GetAll()
}

func (s *TypeBusinessService) GetById(id int) (booking.TypeBusiness, error) {
	return s.repo.GetById(id)
}

func (s *TypeBusinessService) Delete(id int) error {
	return s.repo.Delete(id, true)
}

func (s *TypeBusinessService) Update(id int, input booking.UpdateTypeBusinessInput) error {

	if err := input.Validate(); err != nil {
		return err
	}

	return s.repo.Update(id, input)
}