package models

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/google/uuid"
	"gotest/config"
)

type Car struct {
	GUID          string `json:"guid"`
	CarName       string `json:"car_name"`
	CarBrand      string `json:"car_brand"`
	CarType       string `json:"car_type"`
	CarYear       int64  `json:"car_year"`
	CarDescription string `json:"car_description"`
}

func AddCar(car Car) string {
	db := config.CreateConnection()
	defer db.Close()

	car.GUID = uuid.New().String()
	sqlStatement := `INSERT INTO car (guid, name, brand, type, year, description) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := db.Exec(sqlStatement, car.GUID, car.CarName, car.CarBrand, car.CarType, car.CarYear, car.CarDescription)
	if err != nil {
		log.Fatalf("Gagal mengeksekusi query. %v", err)
	}

	fmt.Printf("Data mobil berhasil ditambahkan: %v\n", car.GUID)
	return car.GUID
}

func GetAllCars() ([]Car, error) {
	db := config.CreateConnection()
	defer db.Close()

	var cars []Car

	sqlStatement := `SELECT * FROM car`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		log.Fatalf("Gagal mengeksekusi query. %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var car Car

		err = rows.Scan(&car.GUID, &car.CarName, &car.CarBrand, &car.CarType, &car.CarYear, &car.CarDescription)
		if err != nil {
			log.Fatalf("Gagal mendapatkan data. %v", err)
		}
		cars = append(cars, car)
	}
	return cars, nil
}

func GetCar(guid string) (Car, error) {
	db := config.CreateConnection()
	defer db.Close()

	var car Car

	sqlStatement := `SELECT * FROM car WHERE guid=$1`
	row := db.QueryRow(sqlStatement, guid)
	err := row.Scan(&car.GUID, &car.CarName, &car.CarBrand, &car.CarType, &car.CarYear, &car.CarDescription)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("Data tidak ditemukan!")
		return car, nil
	case nil:
		return car, nil
	default:
		log.Fatalf("Gagal mendapatkan data. %v", err)
	}
	return car, err
}

func UpdateCar(guid string, car Car) int64 {
	db := config.CreateConnection()
	defer db.Close()

	sqlStatement := `UPDATE car SET name=$2, brand=$3, type=$4, year=$5, description=$6 WHERE guid=$1`
	res, err := db.Exec(sqlStatement, guid, car.CarName, car.CarBrand, car.CarType, car.CarYear, car.CarDescription)
	if err != nil {
		log.Fatalf("Gagal mengeksekusi query. %v", err)
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("Kesalahan saat memperbarui data. %v", err)
	}
	fmt.Printf("Total baris/rekaman yang diperbarui adalah %v\n", rowsAffected)
	return rowsAffected
}

func DeleteCar(guid string) int64 {
	db := config.CreateConnection()
	defer db.Close()

	sqlStatement := `DELETE FROM car WHERE guid=$1`
	res, err := db.Exec(sqlStatement, guid)
	if err != nil {
		log.Fatalf("Gagal mengeksekusi query. %v", err)
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("Data tidak ditemukan. %v", err)
	}
	fmt.Printf("Total data yang dihapus adalah %v", rowsAffected)
	return rowsAffected
}
