package repository

import (
	"fmt"
	"github.com/inzarubin80/booking-services"
	"github.com/jmoiron/sqlx"
)

type ServiceСentersMSSQL struct {
	db *sqlx.DB
}

func NewServiceСentersMSSQL(db *sqlx.DB) *ServiceСentersMSSQL {
	return &ServiceСentersMSSQL{db: db}
}

func (r *ServiceСentersMSSQL) Create(item booking.ServiceСenters) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var ServiceСenterID int
	createItemQuery := fmt.Sprintf(`INSERT INTO %s 
	(ServiceСentreName, CompanieId, Location) 
	values ($1, $2, $3) RETURNING ServiceСenterID`, serviceСentersTable)
	row := tx.QueryRow(createItemQuery, item.ServiceСentreName, item.CompanieId, item.Location)
	err = row.Scan(&ServiceСenterID)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return ServiceСenterID, tx.Commit()
}

func (r *ServiceСentersMSSQL) GetAll() ([]booking.ServiceСenters, error) {
	var items []booking.ServiceСenters
	query := fmt.Sprintf(`SELECT ServiceСenterID, ServiceСentreName, CompanieId, Location FROM %s WHERE MarkDeletion=0`, serviceСentersTable)
	if err := r.db.Select(&items, query); err != nil {
		return nil, err
	}
	return items, nil
}

func (r *ServiceСentersMSSQL) GetById(itemId int) (booking.ServiceСenters, error) {
	var item booking.ServiceСenters

	query := fmt.Sprintf(`SELECT ServiceСenterID, ServiceСentreName, CompanieId, Location FROM %s WHERE ServiceСenterID = $1`, serviceСentersTable)
	if err := r.db.Get(&item, query, itemId); err != nil {
		return item, err
	}
	return item, nil
}

func (r *ServiceСentersMSSQL) Delete(itemId int, markDeletion bool) error {

	query := fmt.Sprintf(`UPDATE %s  SET MarkDeletion = $1 WHERE ServiceСenterID = $2;`, typeBusinessTable)
	_, err := r.db.Exec(query, markDeletion, itemId)
	return err

}

func (r *ServiceСentersMSSQL) Update(id int, item booking.UpdateServiceСentersInput) error {
	
	query := fmt.Sprintf(`UPDATE %s  
	SET ServiceСentreName = $1, 
	Location = $2,
	WHERE ServiceСenterID = $3`,serviceСentersTable)
	_, err := r.db.Exec(query, item.ServiceСentreName, item.Location, id)
	return err

}
