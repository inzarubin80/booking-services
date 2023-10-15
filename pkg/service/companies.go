package service

import (
	"github.com/inzarubin80/booking-services"
	"github.com/inzarubin80/booking-services/pkg/repository"
)

type CompaniesService  struct {
	repo repository.Companies
}

func NewCompaniesService(repo repository.Companies) *CompaniesService {
	return &CompaniesService{repo: repo}
}

func (s *CompaniesService) Create(item booking.Companies) (int, error) {
	return s.repo.Create(item)
}

func (s *CompaniesService) GetAll() ([] booking.Companies, error) {
	return s.repo.GetAll()
}

func (s *CompaniesService) GetById(id int) (booking.Companies, error) {
	return s.repo.GetById(id)
}

func (s *CompaniesService) Delete(id int) error {
	return s.repo.Delete(id, true)
}

func (s *CompaniesService) Update(id int, input booking.UpdateCompaniesInput) error {

	if err := input.Validate(); err != nil {
		return err
	}

	return s.repo.Update(id, input)
}