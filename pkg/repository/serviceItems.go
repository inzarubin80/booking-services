package repository

import (
	"fmt"
	"github.com/inzarubin80/booking-services"
	"github.com/jmoiron/sqlx"
)

type ServiceItemsMSSQL( struct {
	db *sqlx.DB
})

func NewServiceItemsMSSQL(db *sqlx.DB) *ServiceItemsMSSQL {
	return &ServiceItemsMSSQL{db: db}
}

func (r *ServiceItemsMSSQL) Create(item booking.ServiceItems) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var ServiceItemsID int
	createItemQuery := fmt.Sprintf(`INSERT INTO %s 
	(ServiceItemsName, Description, ServiceCenterID, UnitPrice, DurationMinutes, MarkDeletion, ServiceGroupID) 
	values ($1, $2, $3) RETURNING ServiceItemsID`, serviceItemsTable)
	row := tx.QueryRow(createItemQuery, item.ServiceItemsName, item.Description, item.ServiceCenterID, item.UnitPrice, item.DurationMinutes, item.MarkDeletion, item.ServiceGroupID)
	err = row.Scan(&ServiceItemsID)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return ServiceItemsID, tx.Commit()
}

func (r *ServiceItemsMSSQL) GetAll() ([]booking.ServiceItems, error) {
	var items []booking.ServiceItems
	query := fmt.Sprintf(`SELECT ServiceItemsID, ServiceItemsName, Description, ServiceCenterID, UnitPrice, DurationMinutes, MarkDeletion, ServiceGroupID) 
	values ($1, $2, $3) RETURNING ServiceItemsID FROM %s WHERE MarkDeletion=0`, serviceItemsTable)
	if err := r.db.Select(&items, query); err != nil {
		return nil, err
	}
	return items, nil
}

func (r *ServiceItemsMSSQL) GetById(itemId int) (booking.ServiceItems, error) {
	var item booking.ServiceItems

	query := fmt.Sprintf(`SELECT ServiceСenterID, ServiceСentreName, CompanieId, Location FROM %s WHERE ServiceСenterID = $1`, serviceItemsTable)
	if err := r.db.Get(&item, query, itemId); err != nil {
		return item, err
	}
	return item, nil
}

func (r *ServiceItemsMSSQL) Delete(itemId int, markDeletion bool) error {

	query := fmt.Sprintf(`UPDATE %s  SET MarkDeletion = $1 WHERE ServiceСenterID = $2;`, typeBusinessTable)
	_, err := r.db.Exec(query, markDeletion, itemId)
	return err

}

func (r *ServiceItemsMSSQL) Update(id int, item booking.UpdateServiceItemsInput) error {
	
	query := fmt.Sprintf(`UPDATE %s  
	SET Description = $1, 
	DurationMinutes = $2,
	ServiceGroupID = $3,
	ServiceItemsName = $4,
	UnitPrice = $5,
	WHERE ServiceСenterID = $3`,serviceItemsTable)
	_, err := r.db.Exec(query, item.Description, item.DurationMinutes,item.ServiceGroupID, item.ServiceItemsName, item.UnitPrice, id)
	return err

}
