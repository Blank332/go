package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"gotest/models"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type response struct {
	GUID      string `json:"guid,omitempty"`
	Message   string `json:"message,omitempty"`
	CarBrand  string `json:"car_brand,omitempty"`
	CarName   string `json:"car_name,omitempty"`
}

type Response struct {
	Status  int           `json:"status"`
	Message string        `json:"message"`
	Data    []models.Car  `json:"data"`
}

// CRUD BY ID

func AddCar(w http.ResponseWriter, r *http.Request) {
	var car models.Car

	err := json.NewDecoder(r.Body).Decode(&car)
	if err != nil {
		log.Fatalf("Gagal mendekode permintaan body.  %v", err)
	}
	insertGUID := models.AddCar(car)
	res := response{
		GUID:    insertGUID,
		Message: "Mobil telah ditambahkan",
	}
	json.NewEncoder(w).Encode(res)
}

func GetCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	params := mux.Vars(r)
	guid := params["guid"]

	// Validasi GUID
	if _, err := uuid.Parse(guid); err != nil {
		http.Error(w, "GUID tidak valid", http.StatusBadRequest)
		return
	}

	car, err := models.GetCar(guid)
	if err != nil {
		log.Printf("Gagal mendapatkan data mobil dengan GUID %s: %v", guid, err)
		http.Error(w, "Gagal mendapatkan data mobil", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(car)
}



func GetAllCars(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	cars, err := models.GetAllCars()
	if err != nil {
		log.Fatalf("Gagal mendapatkan data. %v", err)
	}

	var response Response
	response.Status = 1
	response.Message = "Sukses"
	response.Data = cars

	json.NewEncoder(w).Encode(response)
}

func UpdateCar(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	guid := params["guid"]

	var car models.Car

	err := json.NewDecoder(r.Body).Decode(&car)
	if err != nil {
		log.Fatalf("Gagal mendekode permintaan body.  %v", err)
	}

	updatedRows := models.UpdateCar(guid, car)

	msg := fmt.Sprintf("Daftar mobil berhasil diperbarui. Total data yang diperbarui adalah %v baris/rekaman", updatedRows)
	res := response{
		GUID:    guid,
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}

func DeleteCar(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	guid := params["guid"]

	deletedRows := models.DeleteCar(guid)

	msg := fmt.Sprintf("Mobil berhasil dihapus. Total data yang dihapus adalah %v", deletedRows)
	res := response{
		GUID:    guid,
		Message: msg,
	}
	json.NewEncoder(w).Encode(res)
}
