package controller

type Film struct {
	id           int      `form : "id" json:"id"`
	judul        string   `form : "judul" json:"judul"`
	tahunRilis   int      `form : "tahunRilis" json:"tahunRilis"`
	genre        string   `form : "genre" json:"genre"`
	sutradara    string   `form : "sutradara" json:"sutradara"`
	daftarPemain []string `form : "daftarPemain" json:"daftarPemain"`
	sinopsis     string   `form : "sinopsis" json:"sinopsis"`
	filmType     string   `form : "filmType" json:"filmType"`
}

type FilmResponse struct {
	Status  int    `form : "status" json:"status"`
	Message string `form : "message" json:"message"`
	Data    []Film `form : "data" json:"data"`
}
