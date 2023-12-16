package repository

import (
	"fmt"
	"github.com/inzarubin80/booking-services"
	"github.com/jmoiron/sqlx"
)

type ServiceProducersMSSQL struct {
	db *sqlx.DB
}

func NewServiceProducersMSSQL(db *sqlx.DB) *ServiceProducersMSSQL {
	return &ServiceProducersMSSQL{db: db}
}

func (r *ServiceProducersMSSQL) Create(item booking.ServiceProducers) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	var ServiceProducerID int
	createItemQuery := fmt.Sprintf(`INSERT INTO %s 
	(ServiceProducerName, Description, MarkDeletion, ServiceCenterID) 
	OUTPUT Inserted.ServiceProducerID
	values ($1, $2, $3, $4)`, serviceProducersTable)
	row := tx.QueryRow(createItemQuery, item.ServiceProducerName, item.Description, item.MarkDeletion, item.ServiceCenterID)
	err = row.Scan(&ServiceProducerID)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return ServiceProducerID, tx.Commit()
}

func (r *ServiceProducersMSSQL) GetAll() ([]booking.ServiceProducers, error) {
	var items []booking.ServiceProducers
	query := fmt.Sprintf(`SELECT ServiceProducerID, ServiceProducerName,Description, MarkDeletion, ServiceCenterID FROM %s WHERE MarkDeletion=0`, serviceProducersTable)
	if err := r.db.Select(&items, query); err != nil {
		return nil, err
	}
	return items, nil
}

func (r *ServiceProducersMSSQL) GetById(itemId int) (booking.ServiceProducers, error) {
	var item booking.ServiceProducers

	query := fmt.Sprintf(`SELECT ServiceProducerID, ServiceProducerName, Description, MarkDeletion, ServiceCenterID FROM %s WHERE ServiceProducerID = $1`, serviceProducersTable)
	if err := r.db.Get(&item, query, itemId); err != nil {
		return item, err
	}
	return item, nil
}

func (r *ServiceProducersMSSQL) Delete(itemId int, markDeletion bool) error {
	query := fmt.Sprintf(`UPDATE %s  SET MarkDeletion = $1 WHERE ServiceProducerID = $2;`, typeBusinessTable)
	_, err := r.db.Exec(query, markDeletion, itemId)
	return err
}

func (r *ServiceProducersMSSQL) Update(id int, item booking.UpdateServiceProducersInput) error {
	query := fmt.Sprintf(`UPDATE %s  
	SET Description = $1, 
	ServiceProducerName = $2
	WHERE ServiceProducerID = $3`,serviceGroupsTable)
	_, err := r.db.Exec(query, item.Description, item.ServiceProducerName,  id)
	return err
}
