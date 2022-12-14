package controllers

import (
	"CRUD_API_MEDIUM_COM/config"
	"CRUD_API_MEDIUM_COM/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// READ ALL KTP
func AllKtp(w http.ResponseWriter, r *http.Request) {
	var ktp models.Ktp
	var response models.Response
	var arrKtp []models.Ktp

	db := config.ConnectDB()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM ktp")
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		err = rows.Scan(&ktp.Nik, &ktp.Nama, &ktp.Agama, &ktp.Negara)
		if err != nil {
			log.Fatal(err.Error())
		} else {
			arrKtp = append(arrKtp, ktp)
		}
	}

	response.Status = 200
	response.Message = "SUCCESS"
	response.Data = arrKtp

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}

// INSERT KTP
func InsertKtp(w http.ResponseWriter, r *http.Request) {
	var response models.Response

	db := config.ConnectDB()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	nama := r.FormValue("nama")
	agama := r.FormValue("agama")
	negara := r.FormValue("negara")

	_, err = db.Exec("INSERT INTO ktp(nama, agama, negara) VALUES(?, ?, ?)", nama, agama, negara)
	if err != nil {
		log.Print(err)
		return
	}

	response.Status = 200
	response.Message = "INSERT DATA SUCCESFULLY"
	fmt.Print("INSERT DATA TO DATABASE")

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}

// UPDATE KTP
func UpdateKtp(w http.ResponseWriter, r *http.Request) {
	var response models.Response

	db := config.ConnectDB()
	defer db.Close()

	err := r.ParseMultipartForm(4096)

	if err != nil {
		panic(err)
	}

	nik := r.FormValue("nik")
	nama := r.FormValue("nama")
	agama := r.FormValue("agama")
	negara := r.FormValue("negara")

	if nama != "" && agama == "" && negara == "" {
		_, err = db.Exec("UPDATE ktp SET nama=? WHERE nik=?", nama, nik)
	} else if nama == "" && agama != "" && negara == "" {
		_, err = db.Exec("UPDATE ktp SET agama=? WHERE nik=?", agama, nik)
	} else if nama == "" && agama == "" && negara != "" {
		_, err = db.Exec("UPDATE ktp SET negara=? WHERE nik=?", negara, nik)
	} else {
		_, err = db.Exec("UPDATE ktp SET nama=?, agama=?, negara=? WHERE nik=?", nama, agama, negara, nik)
	}

	if err != nil {
		log.Print(err)
	}

	response.Status = 200
	response.Message = "UPDATE DATA SUCCESFULLY"
	fmt.Print("UPDATE DATA SUCCESFULLY")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// DELETE KTP
func DeleteKtp(w http.ResponseWriter, r *http.Request) {
	var response models.Response

	db := config.ConnectDB()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	nik := r.FormValue("nik")
	_, err = db.Exec("DELETE FROM ktp WHERE nik=?", nik)

	if err != nil {
		log.Print(err)
		return
	}
	response.Status = 200
	response.Message = "DELETE DATA SUCCESSFULLY"
	fmt.Print("DELETE DATA SUCCESSFULLY")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
