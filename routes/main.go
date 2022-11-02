package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/henriqueassiss/cars-api/controllers"
)

func Main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/cars", controllers.GetCars).Methods("Get")
	r.HandleFunc("/api/cars", controllers.InsertCar).Methods("Post")
	r.HandleFunc("/api/cars/{id}", controllers.EditCar).Methods("Put")
	r.HandleFunc("/api/cars/{id}", controllers.GetCarById).Methods("Get")
	r.HandleFunc("/api/cars/{id}", controllers.DeleteCar).Methods("Delete")
	log.Fatal(http.ListenAndServe(":8000", r))
}
