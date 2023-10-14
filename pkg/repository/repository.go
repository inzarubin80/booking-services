package repository
import (
	"github.com/inzarubin80/booking-services"
	"github.com/jmoiron/sqlx"
)

const (
	typeBusinessTable      = "users"
)

type TypeBusiness interface {
	Create(item booking.TypeBusiness) (int, error)
	GetAll() ([]booking.TypeBusiness, error)
	GetById(id int) (booking.TypeBusiness, error)
	Delete(id int, markDeletion bool) error
	Update(id int, item booking.UpdateTypeBusinessInput) error
}

type Repository struct {
	TypeBusiness 
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		TypeBusiness: NewTypeBusinessMSSQL(db),
	}
}
