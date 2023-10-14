package service
import (
	"github.com/inzarubin80/booking-services"
	"github.com/inzarubin80/booking-services/pkg/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Authorization interface {
	CreateUser(user booking.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}


type TypeBusiness interface {
	Create(item booking.TypeBusiness) (int, error)
	GetAll() ([]booking.TypeBusiness, error)
	GetById(id int) (booking.TypeBusiness, error)
	Delete(id int) error
	Update(id int, input booking.UpdateTypeBusinessInput) error
}


type Service struct {
	Authorization
	TypeBusiness

	
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TypeBusiness:NewTypeBusinessService(repos.TypeBusiness),
	}
}
