package booking

import "errors"

type TypeBusiness struct {
    TypeBusinessID         int     `json:"TypeBusinessID"`
    TypeBusinessName       string  `json:"TypeBusinessName"`
    Description            string  `json:"description"`
    NameServiceProducers   string  `json:"NameServiceProducers"`
    UseMultipleSlotBooking bool    `json:"UseMultipleSlotBooking"`
    MarkDeletion           bool    `json:"MarkDeletion"`
    UseSelectSlotService   bool    `json:"UseSelectSlotService"`
}


type UpdateTypeBusinessInput struct {
	
    TypeBusinessName       *string  `json:"TypeBusinessName"`
    Description            string  `json:"description"`
    NameServiceProducers   string  `json:"NameServiceProducers"`
    UseMultipleSlotBooking bool    `json:"UseMultipleSlotBooking"`
    UseSelectSlotService   bool    `json:"UseSelectSlotService"`

}

func (i UpdateTypeBusinessInput) Validate() error {
	if i.TypeBusinessName == nil {
		return errors.New("update structure has no values")
	}

	return nil
}

type User struct {
	Id       int    `json:"-" db:"id"`
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
