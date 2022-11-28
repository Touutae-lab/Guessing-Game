package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// I implement middleware by simple function call
// Documentation version of middleware commended here
// ##################################################
// func auth(next *http.HandlerFunc) http.HandlerFunc {
// 	return func(res http.ReponseWriter, req *http.Request) {
// 		token := req.Header.Get("token")
// 		if token == "token" {
// 			next(res, req)
// 		}
// 		else {
// 			return res.Write()
// 		}
// 	}
// }

// Login
func login(res http.ResponseWriter, req *http.Request) {
	result := make(map[string]string)
	if req.PostFormValue("token") == "token" {
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

func guessing(res http.ResponseWriter, req *http.Request) {

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

func main() {
	fmt.Printf("Starting Server at Port 8000\n")
	router := mux.NewRouter()
	// router.HandleFunc("/", login).Methods("GET")
	router.HandleFunc("/login", login).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", router))
}
