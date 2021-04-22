package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func RegisterMember(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		return
	}

	nama := r.Form.Get("nama")
	email := r.Form.Get("email")
	password := r.Form.Get("password")
	tglLahir := r.Form.Get("tglLahir")
	jenisKelamin := r.Form.Get("jenisKelamin")
	asalNegara := r.Form.Get("asalNegara")

	_, errQuery := db.Exec("INSERT INTO users(nama, email, password, tglLahir, jenisKelamin, asalNegara, usertype, usermember) values (?, ?, ?, ?, ?, ?, ?, ?)",
		nama,
		email,
		password,
		tglLahir,
		jenisKelamin,
		asalNegara,
		1,
		"FREE",
	)

	fmt.Println(errQuery)

	var response UserResponse
	if errQuery == nil {
		response.Status = 200
		response.Message = "Success"
	} else {
		response.Status = 400
		response.Message = "Insert Failed"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}




func sendUnAuthorizedResponse(w http.ResponseWriter) {
	var response UserResponse
	response.Status = 401
	response.Message = "Unauthorized Access"

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func sendErrorResponse(w http.ResponseWriter) {
	var response UserResponse
	response.Status = 400
	response.Message = "Failed"

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func sendSuccessResponse(w http.ResponseWriter) {
	var response UserResponse
	response.Status = 200
	response.Message = "Success"

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
