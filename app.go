package main

import (
	"encoding/json"
	"log"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	. "github.com/gonzaloescobar/customers-restapi/config"
	. "github.com/gonzaloescobar/customers-restapi/dao"
	. "github.com/gonzaloescobar/customers-restapi/models"
	"github.com/gorilla/mux"
)

var config = Config{}
var dao = CustomersDAO{}

// GET list of customers
func AllCustomersEndPoint(w http.ResponseWriter, r *http.Request) {
	customers, err := dao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, customers)
}

// GET a customer by its ID
func FindCustomerEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	customer, err := dao.FindById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Customer ID")
		return
	}
	respondWithJson(w, http.StatusOK, customer)
}

// POST a new customer
func CreateCustomerEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var customer Customer
	if err := json.NewDecoder(r.Body).Decode(&customer); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	customer.ID = bson.NewObjectId()
	if err := dao.Insert(customer); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, customer)
}

// PUT update an existing customer
func UpdateCustomerEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var customer Customer
	if err := json.NewDecoder(r.Body).Decode(&customer); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Update(customer); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

// DELETE an existing customer
func DeleteCustomerEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var customer Customer
	if err := json.NewDecoder(r.Body).Decode(&customer); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Delete(customer); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}

// Define HTTP request routes
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/customers", AllCustomersEndPoint).Methods("GET")
	r.HandleFunc("/customers", CreateCustomerEndPoint).Methods("POST")
	r.HandleFunc("/customers", UpdateCustomerEndPoint).Methods("PUT")
	r.HandleFunc("/customers", DeleteCustomerEndPoint).Methods("DELETE")
	r.HandleFunc("/customers/{id}", FindCustomerEndpoint).Methods("GET")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
