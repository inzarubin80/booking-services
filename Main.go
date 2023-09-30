package main
import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "booking-services/controllers"
    "booking-services/models"	
)



func main() {
    
    
    router := mux.NewRouter()
    router.HandleFunc("/typebusiness", controllers.GetAllTypeBusiness).Methods("GET")

    // router.HandleFunc("/typebusiness", controllers.createTypeBusiness).Methods("POST")
    //router.HandleFunc("/typebusiness/{id}", controllers.getTypeBusiness).Methods("GET")
    //router.HandleFunc("/typebusiness", controllers.getAllTypeBusiness).Methods("GET")
    //router.HandleFunc("/typebusiness/{id}", controllers.updateTypeBusiness).Methods("PUT")
    //router.HandleFunc("/typebusiness/{id}", controllers.deleteTypeBusiness).Methods("DELETE")

    log.Fatal(http.ListenAndServe(":8000", router))

    defer models.GetDB().Close()	
}