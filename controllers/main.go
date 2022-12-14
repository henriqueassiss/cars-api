package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/henriqueassiss/cars-api/database"
	"github.com/henriqueassiss/cars-api/models"
)

func GetCars(w http.ResponseWriter, r *http.Request) {
	var p []models.Car

	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

func GetCarById(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	id := v["id"]
	var p models.Car

	database.DB.First(&p, id)
	json.NewEncoder(w).Encode(p)
}

func InsertCar(w http.ResponseWriter, r *http.Request) {
	var p models.Car
	json.NewDecoder(r.Body).Decode(&p)

	database.DB.Create(&p)
	json.NewEncoder(w).Encode(p)
}

func EditCar(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	id := v["id"]
	var p models.Car

	database.DB.First(&p, id)
	json.NewDecoder(r.Body).Decode(&p)

	database.DB.Save(&p)
	json.NewEncoder(w).Encode(p)
}

func DeleteCar(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	id := v["id"]
	var p models.Car

	database.DB.Delete(&p, id)
	json.NewEncoder(w).Encode(p)
}
