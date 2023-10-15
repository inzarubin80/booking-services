package repository
import (
	"github.com/inzarubin80/booking-services"
	"github.com/jmoiron/sqlx"
)

const (
	typeBusinessTable = "TypeBusiness"
	companiesTable	= "Companies"
	usersTable	= "users"
)

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

type Authorization interface {
	CreateUser(user booking.User) (int, error)
	GetUser(username, password string) (booking.User, error)
}

type Repository struct {
	TypeBusiness 
	Authorization
	Companies
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		TypeBusiness: NewTypeBusinessMSSQL(db),
		Authorization: NewAuthMssql(db),
		Companies: NewCompaniesMSSQL(db),
	}
}
