package booking
import ("errors"
"time")

type TypeBusiness struct {
    TypeBusinessID         int64    `json:"TypeBusinessID" db:"TypeBusinessID"`
    TypeBusinessName       string  `json:"TypeBusinessName" db:"TypeBusinessName"`
    Description            string  `json:"Description" db:"Description"`
    NameServiceProducers   string  `json:"NameServiceProducers" db:"NameServiceProducers"`
    UseMultipleSlotBooking bool    `json:"UseMultipleSlotBooking" db:"UseMultipleSlotBooking"`
    MarkDeletion           bool    `json:"MarkDeletion" db:"MarkDeletion"`
    UseSelectSlotService   bool    `json:"UseSelectSlotService" db:"UseSelectSlotService"`
}

type UpdateTypeBusinessInput struct {	
    TypeBusinessName       *string  `json:"TypeBusinessName"`
    Description            string  `json:"description"`
    NameServiceProducers   string  `json:"NameServiceProducers"`
    UseMultipleSlotBooking bool    `json:"UseMultipleSlotBooking"`
    UseSelectSlotService   bool    `json:"UseSelectSlotService"`
}

type Companies struct {
    CompanieId         int64   `json:"CompanieId" db:"CompanieId"`
    CompanieName       string  `json:"CompanieName" db:"CompanieName"`
    TIN                string  `json:"TIN" db:"TIN"`
    MarkDeletion       bool    `json:"MarkDeletion" db:"MarkDeletion"`
    UserID             int   `json:"UserID" db:"UserID"`
    TypeBusinessID     int64   `json:"TypeBusinessID" db:"TypeBusinessID"`
}

type UpdateCompaniesInput struct {
    CompanieName       string  `json:"CompanieName" db:"CompanieName"`
    TIN                string  `json:"TIN" db:"TIN"`
    TypeBusinessID     int64   `json:"TypeBusinessID" db:"TypeBusinessID"`
    UserID             int     `json:"UserID" db:"UserID"`    
}

func (i UpdateCompaniesInput) Validate() error {
	if i.CompanieName == "" {
		return errors.New("update structure has no values")
	}
	return nil
}


type ServiceСenters struct {
    ServiceСenterID int64  `json:"ServiceСenterID" db:"ServiceСenterID"`
    ServiceСentreName   string `json:"ServiceСentreName" db:"ServiceСentreName"`
    CompanieId  int64  `json:"CompanieId" db:"CompanieId"`
    MarkDeletion    bool   `json:"MarkDeletion" db:"MarkDeletion"`
    Location string `json:"Location" db:"Location"`
}

type UpdateServiceСentersInput struct {
    ServiceСentreName   string `json:"ServiceСentreName" db:"ServiceСentreName"`
    Location string `json:"Location" db:"Location"`
}

type ServiceGroups struct {
    ServiceGroupID int64 `json:"ServiceGroupID" db:"ServiceGroupID"`
    ServiceGroupName int64 `json:"ServiceGroupName" db:"ServiceGroupName"`  
    Description   string `json:"Description" db:"Description"`
    MarkDeletion    bool `json:"MarkDeletion" db:"MarkDeletion"`
    TypeBusinessID int64 `json:"TypeBusinessID" db:"TypeBusinessID"`
}

type UpdateServiceGroupsInput struct {
    ServiceGroupName int64 `json:"ServiceGroupName" db:"ServiceGroupName"`  
    Description   string `json:"ServiceСentreName" db:"ServiceСentreName"`
    TypeBusinessID int64 `json:"TypeBusinessID" db:"TypeBusinessID"`
}


func (i UpdateTypeBusinessInput) Validate() error {
	if i.TypeBusinessName == nil {
		return errors.New("update structure has no values")
	}
	return nil
}

type User struct {
	Id       int    `json:"Id" db:"id"`
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ServiceProducers struct {
	ServiceProducerID      int    `json:"ServiceProducerID" db:"ServiceProducerID"`
	ServiceProducerName    string `json:"ServiceProducerName" binding:"ServiceProducerName"`
	Description string `json:"Description" binding:"Description"`
	MarkDeletion bool `json:"MarkDeletion" binding:"MarkDeletion"`
    ServiceCenterID string `json:"ServiceCenterID" binding:"ServiceCenterID"`
}

type UpdateServiceProducersInput struct {
    ServiceProducerName int64 `json:"ServiceProducerName" db:"ServiceProducerName"`  
    Description   string `json:"Description" db:"Description"`
}

type ServiceItems struct {
	ServiceID      int    `json:"ServiceID" db:"ServiceID"`
    ServiceItemsName string `json:"ServiceItemsName" db:"ServiceItemsName"`  
    Description    string `json:"Description" binding:"Description"`
    ServiceCenterID    string `json:"ServiceCenterID" binding:"ServiceCenterID"`
    UnitPrice    int `json:"UnitPrice" binding:"UnitPrice"`
    DurationMinutes    string `json:"DurationMinutes" binding:"DurationMinutes"`
    MarkDeletion    bool `json:"MarkDeletion" binding:"MarkDeletion"`
    ServiceGroupID    int `json:"ServiceGroupID]" binding:"ServiceGroupID]"` 
}

type UpdateServiceItemsInput struct {
    ServiceItemsName string `json:"ServiceItemsName" db:"ServiceItemsName"`  
    Description string `json:"Description" db:"Description"`  
    UnitPrice   int `json:"UnitPrice" db:"UnitPrice"`
    DurationMinutes   int `json:"DurationMinutes" db:"DurationMinutes"`
    ServiceGroupID   int  `json:"ServiceGroupID" db:"ServiceGroupID"`
}

type Slots struct {
	SlotID      int    `json:"SlotID" db:"SlotID"`
    SlotName string `json:"SlotName" db:"SlotName"`  
    ServiceСenterID int `json:"ServiceСenterID" db:"ServiceСenterID"`  
    ServiceGroupID int `json:"ServiceGroupID" db:"ServiceGroupID"`  
    ServiceItemsID int `json:"ServiceItemsID" db:"ServiceItemsID"` 
    ServiceProducerID int `json:"ServiceProducerID" db:"ServiceProducerID"`
    Date time.Time `json:"Date" db:"Date"` 
    StartTime time.Time `json:"StartTime" db:"StartTime"` 
    EndTime time.Time `json:"EndTime" db:"EndTime"`
    AvailableCapacity int `json:"AvailableCapacity" db:"AvailableCapacity"`  
    Description string `json:"Description" db:"Description"`  
}

type UpdateSlotsInput struct {
    ServiceItemsName string `json:"ServiceItemsName" db:"ServiceItemsName"`  
    UnitPrice   int `json:"UnitPrice" db:"UnitPrice"`
    DurationMinutes   int `json:"DurationMinutes" db:"DurationMinutes"`
    ServiceGroupID   int  `json:"ServiceGroupID" db:"ServiceGroupID"`
    ServiceItemsID int `json:"ServiceItemsID" db:"ServiceItemsID"` 
    ServiceProducerID int `json:"ServiceProducerID" db:"ServiceProducerID"`
    Date time.Time `json:"Date" db:"Date"` 
    StartTime time.Time `json:"StartTime" db:"StartTime"` 
    EndTime time.Time `json:"EndTime" db:"EndTime"`
    AvailableCapacity int `json:"AvailableCapacity" db:"AvailableCapacity"`  
    Description string `json:"Description" db:"Description"`  
}
