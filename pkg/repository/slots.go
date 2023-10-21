package repository

import (
	"fmt"
	"github.com/inzarubin80/booking-services"
	"github.com/jmoiron/sqlx"
)

type SlotsMSSQL struct {
	db *sqlx.DB
}

func NewSlotsMSSQL(db *sqlx.DB) *SlotsMSSQL {
	return &SlotsMSSQL{db: db}
}

func (r *SlotsMSSQL) Create(item booking.Slots) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var SlotID int
	createItemQuery := fmt.Sprintf(`INSERT INTO %s 
	(SlotName, ServiceСenterID, ServiceGroupID, ServiceProducerID, Date, StartTime, EndTime, AvailableCapacity, Description, ServiceItemsID) 
	values ($1, $2, $3,$4,$5,$6,$7,$8,$9) RETURNING SlotID`, slotsTable)
	row := tx.QueryRow(createItemQuery, item.SlotName, item.ServiceСenterID,  item.ServiceGroupID, item.ServiceProducerID,  item.Date, item.StartTime,   item.EndTime,  item.AvailableCapacity, item.Description, item.ServiceItemsID)
	err = row.Scan(&SlotID)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return SlotID, tx.Commit()
}

func (r *SlotsMSSQL) GetAll() ([]booking.Slots, error) {
	var items []booking.Slots
	query := fmt.Sprintf(`SELECT SlotID, SlotName, ServiceСenterID, ServiceGroupID, ServiceProducerID, Date, StartTime, EndTime, AvailableCapacity, Description, ServiceItemsID
	FROM %s WHERE MarkDeletion=0`, slotsTable)
	if err := r.db.Select(&items, query); err != nil {
		return nil, err
	}
	return items, nil
}

func (r *SlotsMSSQL) GetById(itemId int) (booking.Slots, error) {
	var item booking.Slots

	query := fmt.Sprintf(`SELECT SlotID, SlotName, ServiceСenterID, ServiceGroupID, ServiceProducerID, Date, StartTime, EndTime, AvailableCapacity, Description, ServiceItemsID FROM %s WHERE SlotID = $1`, slotsTable)
	if err := r.db.Get(&item, query, itemId); err != nil {
		return item, err
	}
	return item, nil
}

func (r *SlotsMSSQL) Delete(itemId int, markDeletion bool) error {

	query := fmt.Sprintf(`UPDATE %s  SET MarkDeletion = $1 WHERE SlotID = $2;`, slotsTable)
	_, err := r.db.Exec(query, markDeletion, itemId)
	return err

}

func (r *SlotsMSSQL) Update(id int, item booking.UpdateSlotsInput) error {
	
	//SlotID, SlotName, ServiceСenterID, ServiceGroupID, ServiceProducerID, Date, StartTime, EndTime, AvailableCapacity, Description, ServiceItemsID

	query := fmt.Sprintf(`UPDATE %s  
	SET SlotName = $1, 

	
	WHERE SlotID = $3`,slotsTable)
	_, err := r.db.Exec(query, item.Description, item.DurationMinutes,item.ServiceGroupID, item.ServiceItemsName, item.UnitPrice, id)
	return err

}
