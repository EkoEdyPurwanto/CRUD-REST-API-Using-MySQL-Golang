package main

import (
	"CRUD_API_MEDIUM_COM/controllers"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/getKtp", controllers.AllKtp).Methods(http.MethodGet)
	router.HandleFunc("/insertKtp", controllers.InsertKtp).Methods(http.MethodPost)
	router.HandleFunc("/updateKtp", controllers.UpdateKtp).Methods(http.MethodPut)
	router.HandleFunc("/deleteKtp", controllers.DeleteKtp).Methods(http.MethodDelete)
	http.Handle("/", router)
	fmt.Print("CONNECTED TO PORT 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
