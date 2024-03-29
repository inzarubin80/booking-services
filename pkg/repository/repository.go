package repository
import (
	"github.com/inzarubin80/booking-services"
	"github.com/jmoiron/sqlx"
)

const (
	typeBusinessTable = "TypeBusiness"
	companiesTable	= "Companies"
	usersTable	= "users"
	serviceСentersTable = "ServiceСenters"
	serviceGroupsTable = "ServiceGroups"
	serviceProducersTable = "ServiceProducers"
	serviceItemsTable = "ServiceItems"
	slotsTable = "Slots"
	bookingSlotsTable = "bookingSlots"
	
)

type Authorization interface {
	CreateUser(user booking.User) (int, error)
	GetUser(username, password string) (booking.User, error)
}

type TypeBusiness interface {
	Create(item booking.TypeBusiness) (int, error)
	GetAll() ([]booking.TypeBusiness, error)
	GetById(id int) (booking.TypeBusiness, error)
	Delete(id int, markDeletion bool) error
	Update(id int, item booking.UpdateTypeBusinessInput) error
}

type Companies interface {
	Create(item booking.Companies) (int, error)
	GetAll() ([]booking.Companies, error)
	GetById(id int) (booking.Companies, error)
	Delete(id int, markDeletion bool) error
	Update(id int, item booking.UpdateCompaniesInput) error
}

type ServiceСenters interface {
	Create(item booking.ServiceСenters) (int, error)
	GetAll() ([]booking.ServiceСenters, error)
	GetById(id int) (booking.ServiceСenters, error)
	Delete(id int, markDeletion bool) error
	Update(id int, item booking.UpdateServiceСentersInput) error
}

type ServiceGroups interface {
	Create(item booking.ServiceGroups) (int, error)
	GetAll() ([]booking.ServiceGroups, error)
	GetById(id int) (booking.ServiceGroups, error)
	Delete(id int, markDeletion bool) error
	Update(id int, item booking.UpdateServiceGroupsInput) error
}

type ServiceProducers interface {
	Create(item booking.ServiceProducers) (int, error)
	GetAll() ([]booking.ServiceProducers, error)
	GetById(id int) (booking.ServiceProducers, error)
	Delete(id int, markDeletion bool) error
	Update(id int, item booking.UpdateServiceProducersInput) error
}

type ServiceItems interface {
	Create(item booking.ServiceItems) (int, error)
	GetAll() ([]booking.ServiceItems, error)
	GetById(id int) (booking.ServiceItems, error)
	Delete(id int, markDeletion bool) error
	Update(id int, item booking.UpdateServiceItemsInput) error
}

type Slots interface {
	Create(item booking.Slots) (int, error)
	GetAll() ([]booking.Slots, error)
	GetById(id int) (booking.Slots, error)
	Delete(id int, markDeletion bool) error
	Update(id int, item booking.UpdateSlotsInput) error
}


type BookingSlots interface {
	Create(item booking.BookingSlots) (int, error)
	GetAll() ([]booking.BookingSlots, error)
	GetById(id int) (booking.BookingSlots, error)
	Delete(id int, markDeletion bool) error
}



type Repository struct {
	TypeBusiness 
	Authorization
	Companies
	ServiceСenters
	ServiceGroups
	ServiceProducers
	ServiceItems
	Slots
	BookingSlots
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		TypeBusiness: NewTypeBusinessMSSQL(db),
		Authorization: NewAuthMssql(db),
		Companies: NewCompaniesMSSQL(db),
		ServiceСenters: NewServiceСentersMSSQL(db),
		ServiceGroups: NewServiceGroupsMSSQL(db),
		ServiceProducers:NewServiceProducersMSSQL(db),
		ServiceItems:NewServiceItemsMSSQL(db),
		Slots:NewSlotsMSSQL(db),
		BookingSlots:NewBookingSlotsMSSQL(db),
	}
}
