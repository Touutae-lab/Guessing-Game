package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// I implement middleware by simple function call
// Documentation version of middleware commended here
// ##################################################
//
//	func auth(next *http.HandlerFunc) http.HandlerFunc {
//		return func(res http.ReponseWriter, req *http.Request) {
//			token := req.Header.Get("token")
//			if token == "token" {
//				next(res, req)
//			}
//			else {
//				return res.Write()
//			}
//		}
//	}
type Password struct {
	value string `json:"id"`
}

// Login

var x int = rand.Int()
var current_password = new(Password)

func login(res http.ResponseWriter, req *http.Request) {
	result := make(map[string]string)
	if req.PostFormValue("password") == current_password.value {
		createJson(res, req)
	} else {
		res.WriteHeader(http.StatusNotAcceptable)
		result["message"] = "Not Authorized"
		data, err := json.Marshal(result)
		if err != nil {
			log.Println("Failed to marshaling json")
			log.Println(err)
		}
		res.Write(data)
	}
}

func getValue(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusAccepted)
	res.Header().Set("Content-Type", "application/json")

	result := make(map[string]string)
	result["value"] = strconv.Itoa(x)

	data, err := json.Marshal(result)
	if err != nil {
		log.Println("Failed to marshaling json")
		log.Println(err)
	}
	res.Write(data)
}

func guessing(res http.ResponseWriter, req *http.Request) {
	return
}

// Create Content for token session
func createJson(res http.ResponseWriter, req *http.Request) {

	res.WriteHeader(http.StatusCreated)
	res.Header().Set("Content-Type", "application/json")
	res.Header().Add("token", "token")

	result := make(map[string]string)
	result["message"] = "Session Created"

	data, err := json.Marshal(result)
	if err != nil {
		log.Println("Failed to marshaling json")
		log.Println(err)
	}
	res.Write(data)
}

// 2nd Middleware
func auth(next http.HandlerFunc) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		if req.PostFormValue("token") == "token" {
			next(res, req)
		} else {
			var result = make(map[string]string)
			res.WriteHeader(http.StatusNotAcceptable)
			result["message"] = "Not Authorized"
			data, err := json.Marshal(result)
			if err != nil {
				log.Println("Failed to marshaling json")
				log.Println(err)
			}
			res.Write(data)
		}
	}
}

func main() {
	fmt.Printf("Starting Server at Port 8000\n")
	current_password.value = "123456"
	fmt.Printf(current_password.value)
	router := mux.NewRouter()
	router.HandleFunc("/login", login).Methods("POST")
	router.HandleFunc("/guessing", auth(guessing)).Methods("POST")
	router.HandleFunc("/value", auth(getValue)).Methods("GET")
	router.HandleFunc("/getPassword")
	log.Fatal(http.ListenAndServe(":8000", router))
}
