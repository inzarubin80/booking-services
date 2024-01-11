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

type Companies interface {
	Create(item booking.Companies) (int, error)
	GetAll() ([]booking.Companies, error)
	GetById(id int) (booking.Companies, error)
	Delete(id int) error
	Update(id int, input booking.UpdateCompaniesInput) error
}

type ServiceСenters interface {
	Create(item booking.ServiceСenters) (int, error)
	GetAll() ([]booking.ServiceСenters, error)
	GetById(id int) (booking.ServiceСenters, error)
	Delete(id int) error
	Update(id int, input booking.UpdateServiceСentersInput) error
}

type ServiceGroups interface {
	Create(item booking.ServiceGroups) (int, error)
	GetAll() ([]booking.ServiceGroups, error)
	GetById(id int) (booking.ServiceGroups, error)
	Delete(id int) error
	Update(id int, input booking.UpdateServiceGroupsInput) error
}

type ServiceProducers interface {
	Create(item booking.ServiceProducers) (int, error)
	GetAll() ([]booking.ServiceProducers, error)
	GetById(id int) (booking.ServiceProducers, error)
	Delete(id int) error
	Update(id int, input booking.UpdateServiceProducersInput) error
}

type ServiceItems interface {
	Create(item booking.ServiceItems) (int, error)
	GetAll() ([]booking.ServiceItems, error)
	GetById(id int) (booking.ServiceItems, error)
	Delete(id int) error
	Update(id int, input booking.UpdateServiceItemsInput) error
}

type Slots interface {
	Create(item booking.Slots) (int, error)
	GetAll() ([]booking.Slots, error)
	GetById(id int) (booking.Slots, error)
	Delete(id int) error
	Update(id int, input booking.UpdateSlotsInput) error

}


type BookingSlots interface {
	Create(item booking.BookingSlots) (int, error)
}

type Service struct {
	Authorization
	TypeBusiness
	Companies
	ServiceСenters
	ServiceGroups
	ServiceProducers
	ServiceItems
	Slots
	BookingSlots
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TypeBusiness:NewTypeBusinessService(repos.TypeBusiness),
		Companies:NewCompaniesService(repos.Companies),
		ServiceСenters:NewServiceСentersService(repos.ServiceСenters),
		ServiceGroups:NewServiceGroupsServic(repos.ServiceGroups),
		ServiceProducers: NewServiceProducersService(repos.ServiceProducers),
		ServiceItems:NewServiceItemsService(repos.ServiceItems),
		Slots:NewSlotsService(repos.Slots),
		BookingSlots:NewServiceBookingSlots(repos.BookingSlots),
	}
}
