package models

type Ktp struct {
	Nik    int    `json:"nik"`
	Nama   string `json:"nama"`
	Agama  string `json:"agama"`
	Negara string `json:"negara"`
}

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Ktp
}
