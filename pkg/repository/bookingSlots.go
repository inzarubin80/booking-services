package repository

import (
	"errors"
	"fmt"
	"github.com/inzarubin80/booking-services"
	"github.com/jmoiron/sqlx"
)

type BookingSlotsMSSQL struct {
	db *sqlx.DB
}

func NewBookingSlotsMSSQL(db *sqlx.DB) *BookingSlotsMSSQL {
	return &BookingSlotsMSSQL{db: db}
}

func (r *BookingSlotsMSSQL) Create(item booking.BookingSlots) (int, error) {

	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	query := `Select 
				isNUll(Count(BookingSlotsID), 0) as NumberBookings
				,Min(Slots.AvailableCapacity) as AvailableCapacity 
				,Slots.SlotID 
			from BookingSlots as BookingSlots 
				join Slots as Slots on
					BookingSlots.SlotID = Slots.SlotID
					AND BookingSlots.MarkDeletion = 0
			Where 
				BookingSlots.SlotID = $1
			Group By
				Slots.SlotID`

	var bookingState booking.BookingState

	if err := r.db.Get(&bookingState, query, item.SlotID); err != nil {
		return 0, err
	}

	if bookingState.AvailableCapacity <= bookingState.NumberBookings {
		return 0, errors.New("no free places")
	}

	var BookingSlotsID int
	createItemQuery := fmt.Sprintf(`
	INSERT INTO %s 
	(Note, ServiceID, SlotID, UserID) 
	values ($1, $2, $3, $4) RETURNING BookingSlotsID`, bookingSlotsTable)
	row := tx.QueryRow(createItemQuery, item.Note, item.ServiceID, item.SlotID, item.UserID)
	err = row.Scan(&BookingSlotsID)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return BookingSlotsID, tx.Commit()
}

func (r *BookingSlotsMSSQL) GetAll() ([]booking.BookingSlots, error) {
	var items []booking.BookingSlots
	query := fmt.Sprintf(`SELECT BookingSlotsID, SlotID, UserID, ServiceID, MarkDeletion, Note FROM %s WHERE MarkDeletion=0`, bookingSlotsTable)
	if err := r.db.Select(&items, query); err != nil {
		return nil, err
	}
	return items, nil
}

func (r *BookingSlotsMSSQL) GetById(itemId int) (booking.BookingSlots, error) {
	var item booking.BookingSlots

	query := fmt.Sprintf(`SELECT BookingSlotsID, SlotID, UserID, ServiceID, MarkDeletion, Note FROM %s WHERE BookingSlotsID = $1`, bookingSlotsTable)
	if err := r.db.Get(&item, query, itemId); err != nil {
		return item, err
	}
	return item, nil
}

func (r *BookingSlotsMSSQL) Delete(itemId int, markDeletion bool) error {

	query := fmt.Sprintf(`UPDATE %s  SET MarkDeletion = $1 WHERE BookingSlotsID = $2;`, bookingSlotsTable)
	_, err := r.db.Exec(query, markDeletion, itemId)
	return err

}
