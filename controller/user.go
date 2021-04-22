package controller

type User struct {
	id           int    `form : "id" json:"id"`
	nama         string `form : "nama" json:"nama"`
	email        string `form : "email" json:"email"`
	password     string `form : "password" json:"password"`
	tglLahir     string `form : "tglLahir" json:"tglLahir"`
	jenisKelamin string `form : "jenisKelamin" json:"jenisKelamin"`
	asalNegara   string `form : "asalNegara" json:"asalNegara"`
	userType     int    `form : "userType" json:"userType"`
	userMember   string `form : "userMember" json:"userMember"`
}

type UserResponse struct {
	Status  int    `form : "status" json:"status"`
	Message string `form : "message" json:"message"`
	Data    []User `form : "data" json:"data"`
}
