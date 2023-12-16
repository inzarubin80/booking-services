package repository

import (
	"fmt"
	"github.com/inzarubin80/booking-services"
	"github.com/jmoiron/sqlx"
)

type TypeBusinessMSSQL struct {
	db *sqlx.DB
}

func NewTypeBusinessMSSQL(db *sqlx.DB) *TypeBusinessMSSQL {
	return &TypeBusinessMSSQL{db: db}
}

func (r *TypeBusinessMSSQL) Create(item booking.TypeBusiness) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var TypeBusinessID int
	createItemQuery := fmt.Sprintf(`INSERT INTO %s 
	(TypeBusinessName, Description, NameServiceProducers, UseMultipleSlotBooking, MarkDeletion, UseSelectSlotService) 
	OUTPUT Inserted.TypeBusinessID
	values ($1, $2, $3,$4,$5,$6)`, typeBusinessTable)
	row := tx.QueryRow(createItemQuery, item.TypeBusinessName, item.Description, item.NameServiceProducers, item.UseMultipleSlotBooking, item.MarkDeletion, item.UseSelectSlotService)
	err = row.Scan(&TypeBusinessID)
	if err != nil {
		tx.Rollback()
		return 0, err		
	}
	return TypeBusinessID, tx.Commit()
}

func (r *TypeBusinessMSSQL) GetAll() ([]booking.TypeBusiness, error) {
	var items []booking.TypeBusiness
	query := fmt.Sprintf(`SELECT TypeBusinessID, TypeBusinessName, 
	Description, NameServiceProducers, UseMultipleSlotBooking, 
	MarkDeletion, UseSelectSlotService FROM %s WHERE MarkDeletion=0`, typeBusinessTable)
	if err := r.db.Select(&items, query); err != nil {
		return nil, err
	}

	return items, nil
}

func (r *TypeBusinessMSSQL) GetById(itemId int) (booking.TypeBusiness, error) {
	var item booking.TypeBusiness

	query := fmt.Sprintf(`SELECT 
	TypeBusinessID, 
	TypeBusinessName,
	Description,
	NameServiceProducers, 
	UseMultipleSlotBooking,
	MarkDeletion, 
	UseSelectSlotService 
		FROM %s 
	WHERE TypeBusinessID = $1`, typeBusinessTable)
	if err := r.db.Get(&item, query, itemId); err != nil {
		return item, err
	}
	return item, nil
}

func (r *TypeBusinessMSSQL) Delete(itemId int, markDeletion bool) error {

	query := fmt.Sprintf(`UPDATE %s  SET MarkDeletion = $1 WHERE TypeBusinessID = $2;`, typeBusinessTable)
	_, err := r.db.Exec(query, markDeletion, itemId)
	return err

}

func (r *TypeBusinessMSSQL) Update(id int, item booking.UpdateTypeBusinessInput) error {
	query := fmt.Sprintf(`UPDATE %s  
	SET TypeBusinessName = $1, 
	Description = $2,
	NameServiceProducers = $3, 
	UseMultipleSlotBooking = $4, 
	UseSelectSlotService = $5
	WHERE TypeBusinessID = $6`,typeBusinessTable)
	_, err := r.db.Exec(query, item.TypeBusinessName, item.Description, item.NameServiceProducers, item.UseMultipleSlotBooking, item.UseSelectSlotService, id)
	return err

}
