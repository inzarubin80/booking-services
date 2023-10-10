package booking

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
	
    TypeBusinessName       string  `json:"TypeBusinessName"`
    Description            string  `json:"description"`
    NameServiceProducers   string  `json:"NameServiceProducers"`
    UseMultipleSlotBooking bool    `json:"UseMultipleSlotBooking"`
    UseSelectSlotService   bool    `json:"UseSelectSlotService"`

}