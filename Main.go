package main
import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "booking-services/controllers"
   // "booking-services/models"	


)



func main() {
    
    
 
    
    router := mux.NewRouter()
    router.HandleFunc("/typebusiness", controllers.GetAllTypeBusiness).Methods("GET")
    router.HandleFunc("/typebusiness/{id}", controllers.GetTypeBusiness).Methods("GET")
    //router.HandleFunc("/typebusiness/{id}", controllers.updateTypeBusiness).Methods("PUT")
    router.HandleFunc("/typebusiness/{id}/", controllers.DeleteTypeBusiness).Methods("DELETE")
    router.HandleFunc("/typebusiness", controllers.CreateTypeBusiness).Methods("POST")
    

    log.Fatal(http.ListenAndServe(":8000", router))

    //defer models.GetDB().Close()	
}