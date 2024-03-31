package router

import (
	"gotest/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/car", controller.GetAllCars).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/car", controller.AddCar).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/car/{guid}", controller.GetCar).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/car/{guid}", controller.UpdateCar).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/car/{guid}", controller.DeleteCar).Methods("DELETE", "OPTIONS")	

	return router
}