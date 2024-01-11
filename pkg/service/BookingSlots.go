package service

import (
	"github.com/inzarubin80/booking-services"
	"github.com/inzarubin80/booking-services/pkg/repository"
)

type ServiceBookingSlots struct {
	repo repository.BookingSlots
}

func NewServiceBookingSlots(repo repository.BookingSlots) *ServiceBookingSlots {
	return &ServiceBookingSlots{repo: repo}
}

func (s *ServiceBookingSlots) Create(item booking.BookingSlots) (int, error) {
	return s.repo.Create(item)
}
