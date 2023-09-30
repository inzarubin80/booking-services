package controllers

import (
	"booking-services/models"
	u "booking-services/utils"
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

var GetAllTypeBusiness = func (w http.ResponseWriter, r *http.Request)  {

	typeBusinesses := models.GetAllTypeBusiness();
	resp := u.Message(true, "success")
	resp["data"] = typeBusinesses
	u.Respond(w, resp)
}

var GetTypeBusiness = func (w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"]);

	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding id"))
		return
	}
	typeBusiness := models.GetTypeBusiness(id);
	resp := u.Message(true, "success")
	resp["data"] = typeBusiness
	u.Respond(w, resp)
}

var CreateTypeBusiness = func (w http.ResponseWriter, r *http.Request)  {

	typeBusiness := &models.TypeBusiness{}
	err := json.NewDecoder(r.Body).Decode(typeBusiness)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}
	id := models.CreateTypeBusiness(*typeBusiness);
	resp := u.Message(true, "success")
	resp["data"] = id
	u.Respond(w, resp)
}

var  DeleteTypeBusiness = func (w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	MarkDeletion, err := strconv.Atoi(vars["MarkDeletion"]);
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding MarkDeletion"))
		return
	}
	id, err := strconv.Atoi(vars["id"]);
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding id"))
		return
	}
	MarkDeletion = models.DeleteTypeBusiness(id, MarkDeletion);
	resp := u.Message(true, "success")
	resp["data"] = MarkDeletion
	u.Respond(w, resp)
}

