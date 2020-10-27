package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	//"io"
	//"io/ioutil"
)

type Coche struct {
	Id          int    `json:id`
	Brand       string `json:brand`
	Model       string `json:model`
	Horse_power int    `json:horse_power`
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/service/v1/cars/{ident}", endpointGetter).Methods("GET")
	log.Fatal(http.ListenAndServe(":8082", myRouter))

}

var config string

func initialize() {
	data, err := ioutil.ReadFile("/service/conf")
	if err != nil {
		fmt.Println("error reading configuration", err)
		return
	}
	config = string(data)
	fmt.Println("IP DE LA DB:", config)
}
func endpointGetter(w http.ResponseWriter, r *http.Request) {

	//	fmt.Println("GETTER MYSQL CONNECTION")
	vars := mux.Vars(r)
	ident := vars["ident"]

	db, err := sql.Open("mysql", "root:Miquelpiloto1@tcp("+config+":3306)/Concesionario")

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	w.Header().Set("Content-Type", "application/json")
	results, err := db.Query("SELECT id,brand,model,horse_power FROM Concesionario.Coches where id=?", ident)

	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var coche Coche

		err = results.Scan(&coche.Id, &coche.Brand, &coche.Model, &coche.Horse_power)
		if err != nil {
			panic(err.Error())
		}
		b, err := json.Marshal(coche)

		if err != nil {
			panic(err.Error())
		}
		fmt.Fprintln(w, string(b))
		//		fmt.Println("GET FINAL")

	}

}

func main() {
	initialize()
	handleRequests()

}
