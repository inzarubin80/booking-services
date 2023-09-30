package controllers
import (
	"booking-services/models"	
	"net/http"
	"encoding/json"
	
)

var GetAllTypeBusiness = func (w http.ResponseWriter, r *http.Request)  {

	dataResponseGetAllTypeBusiness := models.GetAllTypeBusiness();
	json.NewEncoder(w).Encode(dataResponseGetAllTypeBusiness)

}
