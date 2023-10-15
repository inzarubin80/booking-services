package repository

import (
	"fmt"
	"github.com/inzarubin80/booking-services"
	"github.com/jmoiron/sqlx"
)

type CompaniesMSSQL struct {
	db *sqlx.DB
}

func NewCompaniesMSSQL(db *sqlx.DB) *CompaniesMSSQL {
	return &CompaniesMSSQL{db: db}
}

func (r *CompaniesMSSQL) Create(item booking.Companies) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	var CompanieId int
	createItemQuery := fmt.Sprintf(`INSERT INTO %s 
	(CompanieName, TIN, MarkDeletion, UserID,TypeBusinessID) 
	values ($1, $2) RETURNING CompanieId`, companiesTable)
	row := tx.QueryRow(createItemQuery, item.CompanieName, item.TIN, item.UserID, item.TypeBusinessID)
	err = row.Scan(&CompanieId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return CompanieId, tx.Commit()
}

func (r *CompaniesMSSQL) GetAll() ([]booking.Companies, error) {
	var items []booking.Companies
	query := fmt.Sprintf(`SELECT CompanieId, CompanieName, 
	TIN, UserID,  TypeBusinessID FROM %s WHERE MarkDeletion=0`, companiesTable)
	if err := r.db.Select(&items, query); err != nil {
		return nil, err
	}
	return items, nil
}

func (r *CompaniesMSSQL) GetById(itemId int) (booking.Companies, error) {
	var item booking.Companies

	query := fmt.Sprintf(`SELECT 
	CompanieId, 
	CompanieName,
	TIN,
	UserID,
	TypeBusinessID
		FROM %s 
	WHERE CompanieId = $1`, companiesTable)
	if err := r.db.Get(&item, query, itemId); err != nil {
		return item, err
	}
	return item, nil
}

func (r *CompaniesMSSQL) Delete(itemId int, markDeletion bool) error {

	query := fmt.Sprintf(`UPDATE %s  SET MarkDeletion = $1 WHERE CompanieId = $2;`, companiesTable)
	_, err := r.db.Exec(query, markDeletion, itemId)
	return err

}

func (r *CompaniesMSSQL) Update(id int, item booking.UpdateCompaniesInput) error {
	query := fmt.Sprintf(`UPDATE %s  
	SET CompanieName = $1, TIN = $2 TypeBusinessID=$3 WHERE CompanieId = $3`, companiesTable)
	_, err := r.db.Exec(query, item.CompanieName, item.TIN, id)
	return err
}
