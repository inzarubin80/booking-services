package repository

import (
	"fmt"
	"github.com/inzarubin80/booking-services"
	"github.com/jmoiron/sqlx"
)

type ServiceGroupsMSSQL struct {
	db *sqlx.DB
}

func NewServiceGroupsMSSQL(db *sqlx.DB) *ServiceGroupsMSSQL {
	return &ServiceGroupsMSSQL{db: db}
}

func (r *ServiceGroupsMSSQL) Create(item booking.ServiceGroups) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	var ServiceGroupID int
	createItemQuery := fmt.Sprintf(`INSERT INTO %s 
	(ServiceGroupName, Description, TypeBusinessID, MarkDeletion) 
	OUTPUT Inserted.ServiceGroupID
	values ($1, $2, $3, $4)`, serviceGroupsTable)
	row := tx.QueryRow(createItemQuery, item.ServiceGroupName, item.Description, item.TypeBusinessID, item.MarkDeletion)
	err = row.Scan(&ServiceGroupID)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return ServiceGroupID, tx.Commit()
}

func (r *ServiceGroupsMSSQL) GetAll() ([]booking.ServiceGroups, error) {
	var items []booking.ServiceGroups
	query := fmt.Sprintf(`SELECT ServiceGroupID, ServiceGroupName, Description, TypeBusinessID FROM %s WHERE MarkDeletion=0`, serviceGroupsTable)
	if err := r.db.Select(&items, query); err != nil {
		return nil, err
	}
	return items, nil
}

func (r *ServiceGroupsMSSQL) GetById(itemId int) (booking.ServiceGroups, error) {
	var item booking.ServiceGroups

	query := fmt.Sprintf(`SELECT ServiceGroupID, ServiceGroupName, Description, TypeBusinessID FROM %s WHERE ServiceGroupID = $1`, serviceGroupsTable)
	if err := r.db.Get(&item, query, itemId); err != nil {
		return item, err
	}
	return item, nil
}

func (r *ServiceGroupsMSSQL) Delete(itemId int, markDeletion bool) error {

	query := fmt.Sprintf(`UPDATE %s  SET MarkDeletion = $1 WHERE ServiceGroupID = $2;`, typeBusinessTable)
	_, err := r.db.Exec(query, markDeletion, itemId)
	return err

}

func (r *ServiceGroupsMSSQL) Update(id int, item booking.UpdateServiceGroupsInput) error {
	query := fmt.Sprintf(`UPDATE %s  
	SET Description = $1, 
	ServiceGroupName = $2,
	TypeBusinessID = $3
	WHERE ServiceGroupID = $3`,serviceGroupsTable)
	_, err := r.db.Exec(query, item.Description, item.ServiceGroupName, item.TypeBusinessID, id)
	return err

}
