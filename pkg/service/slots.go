package service

import (
	"github.com/inzarubin80/booking-services"
	"github.com/inzarubin80/booking-services/pkg/repository"
)

type SlotsService struct {
	repo repository.Slots
}

func NewSlotsService(repo repository.Slots) *SlotsService {
	return &SlotsService{repo: repo}
}

func (s *SlotsService) Create(item booking.Slots) (int, error) {
	return s.repo.Create(item)
}

func (s *SlotsService) GetAll() ([]booking.Slots, error) {
	return s.repo.GetAll()
}

func (s *SlotsService) GetById(id int) (booking.Slots, error) {
	return s.repo.GetById(id)
}

func (s *SlotsService) Delete(id int) error {
	return s.repo.Delete(id, true)
}

func (s *SlotsService) Update(id int, input booking.UpdateSlotsInput) error {
	return s.repo.Update(id, input)
}
