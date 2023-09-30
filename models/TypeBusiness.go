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
    UseMultipleSlotBooking int    `json:"UseMultipleSlotBooking"`
    MarkDeletion           int    `json:"MarkDeletion"`
    UseSelectSlotService   int    `json:"UseSelectSlotService"`
}


var GetAllTypeBusiness = func() []TypeBusiness{
   
    statement := "SELECT TypeBusinessID, TypeBusinessName, Description, NameServiceProducers, UseMultipleSlotBooking, MarkDeletion, UseSelectSlotService FROM TypeBusiness WHERE MarkDeletion=0;"

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


    return typeBusinesses;
}
var CreateTypeBusiness = func (typeBusiness TypeBusiness) int64 {

    statement := "INSERT INTO TypeBusiness (TypeBusinessName, Description, NameServiceProducers, UseMultipleSlotBooking, MarkDeletion, UseSelectSlotService) OUTPUT INSERTED.TypeBusinessID VALUES (@TypeBusinessName, @Description, @NameServiceProducers, @UseMultipleSlotBooking, @MarkDeletion, @UseSelectSlotService); select ID = convert(bigint, SCOPE_IDENTITY())"
    rows, err := GetDB().Query(statement,
        sql.Named("TypeBusinessName", typeBusiness.TypeBusinessName),
        sql.Named("Description", typeBusiness.Description),
        sql.Named("NameServiceProducers", typeBusiness.NameServiceProducers),
        sql.Named("UseMultipleSlotBooking", typeBusiness.UseMultipleSlotBooking),
        sql.Named("MarkDeletion", typeBusiness.MarkDeletion),
        sql.Named("UseSelectSlotService", typeBusiness.UseSelectSlotService),
    )
    
    if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var lastInsertId1 int64
	for rows.Next() {
		rows.Scan(&lastInsertId1)
	}

    return lastInsertId1
}

var GetTypeBusiness = func(id int) *TypeBusiness{
    
    typeBusiness:= &TypeBusiness{}

    statement := "SELECT TypeBusinessID, TypeBusinessName, Description, NameServiceProducers, UseMultipleSlotBooking, MarkDeletion, UseSelectSlotService FROM TypeBusiness WHERE TypeBusinessID = @TypeBusinessID;"
    rows, err := GetDB().Query(statement, sql.Named("TypeBusinessID", id))
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()
    for rows.Next() {
        err = rows.Scan(&typeBusiness.TypeBusinessID, &typeBusiness.TypeBusinessName, &typeBusiness.Description, &typeBusiness.NameServiceProducers, &typeBusiness.UseMultipleSlotBooking, &typeBusiness.MarkDeletion, &typeBusiness.UseSelectSlotService)
        if err != nil {
            log.Fatal(err)
        }
    }

    return typeBusiness;
}

var UpdateTypeBusiness =  func(typeBusiness *TypeBusiness) TypeBusiness {
   
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
        return *typeBusiness
    }

    return *typeBusiness;
}

var DeleteTypeBusiness = func(id int, markDeletion int) int {
  
    if markDeletion!=0 {
        markDeletion = 1 
    }
    statement := "UPDATE TypeBusiness SET MarkDeletion = @MarkDeletion WHERE TypeBusinessID = @TypeBusinessID;"
    _, err := GetDB().Exec(statement,
        sql.Named("TypeBusinessID", id),
        sql.Named("MarkDeletion", markDeletion),
    )

    if err != nil {
        log.Fatal(err)
    }
    return markDeletion;
}
