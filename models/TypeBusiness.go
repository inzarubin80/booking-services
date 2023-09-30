package models
import (
    "database/sql"
    "log"
)

type TypeBusiness struct {
    TypeBusinessID         int64     `json:"TypeBusinessID"`
    TypeBusinessName       string  `json:"TypeBusinessName"`
    Description            string  `json:"description"`
    NameServiceProducers   string  `json:"NameServiceProducers"`
    UseMultipleSlotBooking bool    `json:"UseMultipleSlotBooking"`
    MarkDeletion           bool    `json:"MarkDeletion"`
    UseSelectSlotService   bool    `json:"UseSelectSlotService"`
}

type DataResponseGetAllTypeBusiness struct {
    MessageResponse MessageResponse `json:"MessageResponse"`
    Data []TypeBusiness `json:"Data"`
}

type DataResponseCreateTypeBusinesss struct {
    MessageResponse MessageResponse `json:"MessageResponse"`
    Data int64 `json:"data"`
}

type DataResponseGetTypeBusiness struct {
    MessageResponse MessageResponse `json:"MessageResponse"`
    Data TypeBusiness `json:"data"`
}

type DataResponseUpdateTypeBusiness struct {
    MessageResponse MessageResponse `json:"MessageResponse"`
}

type DataResponseDeleteTypeBusiness struct {
    MessageResponse MessageResponse `json:"MessageResponse"`
}

var GetAllTypeBusiness = func() DataResponseGetAllTypeBusiness{
   
    dataResponseGetAllTypeBusiness := DataResponseGetAllTypeBusiness{}
    dataResponseGetAllTypeBusiness.MessageResponse = MessageResponse{}

    statement := "SELECT TypeBusinessID, TypeBusinessName, Description, NameServiceProducers, UseMultipleSlotBooking, MarkDeletion, UseSelectSlotService FROM TypeBusiness;"

    rows, err := GetDB().Query(statement)
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    var typeBusinesses []TypeBusiness

    for rows.Next() {
        var typeBusiness TypeBusiness
        err = rows.Scan(&typeBusiness.TypeBusinessID, &typeBusiness.TypeBusinessName, &typeBusiness.Description, &typeBusiness.NameServiceProducers, &typeBusiness.UseMultipleSlotBooking, &typeBusiness.MarkDeletion, &typeBusiness.UseSelectSlotService)
        if err != nil {
            log.Fatal(err)
        }
        typeBusinesses = append(typeBusinesses, typeBusiness)
    }

    dataResponseGetAllTypeBusiness.Data=typeBusinesses;
    return dataResponseGetAllTypeBusiness;
}
var CreateTypeBusiness = func (typeBusiness TypeBusiness) DataResponseCreateTypeBusinesss {
   
    dataResponseCreateTypeBusinesss := DataResponseCreateTypeBusinesss{}
    dataResponseCreateTypeBusinesss.MessageResponse = MessageResponse{}
    statement := "INSERT INTO TypeBusiness (TypeBusinessName, Description, NameServiceProducers, UseMultipleSlotBooking, MarkDeletion, UseSelectSlotService) OUTPUT INSERTED.TypeBusinessID VALUES (@TypeBusinessName, @Description, @NameServiceProducers, @UseMultipleSlotBooking, @MarkDeletion, @UseSelectSlotService);"
    result, err := GetDB().Exec(statement,
        sql.Named("TypeBusinessName", typeBusiness.TypeBusinessName),
        sql.Named("Description", typeBusiness.Description),
        sql.Named("NameServiceProducers", typeBusiness.NameServiceProducers),
        sql.Named("UseMultipleSlotBooking", typeBusiness.UseMultipleSlotBooking),
        sql.Named("MarkDeletion", typeBusiness.MarkDeletion),
        sql.Named("UseSelectSlotService", typeBusiness.UseSelectSlotService),
    )

    if err != nil {
        log.Fatal(err)
        return dataResponseCreateTypeBusinesss
    }

    id, err := result.LastInsertId()
    if err != nil {
        log.Fatal(err)
        return  dataResponseCreateTypeBusinesss
    }

    typeBusiness.TypeBusinessID = id
    dataResponseCreateTypeBusinesss.Data=id;
    return dataResponseCreateTypeBusinesss

}

var GetTypeBusiness = func(id int) DataResponseGetTypeBusiness{
    
    dataResponseGetTypeBusiness := DataResponseGetTypeBusiness{}
    dataResponseGetTypeBusiness.MessageResponse = MessageResponse{}

    statement := "SELECT TypeBusinessID, TypeBusinessName, Description, NameServiceProducers, UseMultipleSlotBooking, MarkDeletion, UseSelectSlotService FROM TypeBusiness WHERE TypeBusinessID = @TypeBusinessID;"
    rows, err := GetDB().Query(statement, sql.Named("TypeBusinessID", id))
    if err != nil {
        log.Fatal(err)
        return dataResponseGetTypeBusiness;
    }
    defer rows.Close()

    var typeBusiness TypeBusiness

    for rows.Next() {
        err = rows.Scan(&typeBusiness.TypeBusinessID, &typeBusiness.TypeBusinessName, &typeBusiness.Description, &typeBusiness.NameServiceProducers, &typeBusiness.UseMultipleSlotBooking, &typeBusiness.MarkDeletion, &typeBusiness.UseSelectSlotService)
        if err != nil {
            log.Fatal(err)
            return dataResponseGetTypeBusiness;
        }
    }

    dataResponseGetTypeBusiness.Data = typeBusiness;
    return dataResponseGetTypeBusiness;
}

var UpdateTypeBusiness =  func(typeBusiness *TypeBusiness) DataResponseUpdateTypeBusiness {
   
    dataResponseupdateTypeBusiness := DataResponseUpdateTypeBusiness{}
    dataResponseupdateTypeBusiness.MessageResponse = MessageResponse{}
  
    statement := "UPDATE TypeBusiness SET TypeBusinessName = @TypeBusinessName, Description = @Description, NameServiceProducers = @NameServiceProducers, UseMultipleSlotBooking = @UseMultipleSlotBooking, MarkDeletion = @MarkDeletion, UseSelectSlotService = @UseSelectSlotService WHERE TypeBusinessID = @TypeBusinessID;"
    _, err := GetDB().Exec(statement,
        sql.Named("TypeBusinessID", typeBusiness.TypeBusinessID),
        sql.Named("TypeBusinessName", typeBusiness.TypeBusinessName),
        sql.Named("Description", typeBusiness.Description),
        sql.Named("NameServiceProducers", typeBusiness.NameServiceProducers),
        sql.Named("UseMultipleSlotBooking", typeBusiness.UseMultipleSlotBooking),
        sql.Named("MarkDeletion", typeBusiness.MarkDeletion),
        sql.Named("UseSelectSlotService", typeBusiness.UseSelectSlotService),
    )

    if err != nil {
        log.Fatal(err)
        return dataResponseupdateTypeBusiness
    }

    return dataResponseupdateTypeBusiness;
}

var DeleteTypeBusiness = func(id int, markDeletion bool) DataResponseDeleteTypeBusiness {

    dataResponseDeleteTypeBusiness := DataResponseDeleteTypeBusiness{}
    dataResponseDeleteTypeBusiness.MessageResponse = MessageResponse{}
  
    statement := "UPDATE TypeBusiness SET MarkDeletion = @MarkDeletion WHERE TypeBusinessID = @TypeBusinessID;"
    _, err := GetDB().Exec(statement,
        sql.Named("TypeBusinessID", id),
        sql.Named("MarkDeletion", markDeletion),
    )

    if err != nil {
        log.Fatal(err)
        return dataResponseDeleteTypeBusiness
    }
    return dataResponseDeleteTypeBusiness;
}
