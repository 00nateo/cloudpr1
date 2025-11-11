/*
https://www.youtube.com/watch?v=d_L64KT3SFM
Microservice to handle get requests from frontend
Is this containerized?
Need a db?
Volumes?

and then deploy to rancher?

This go file must connect to databse and handle HTTP GET
GET format:
http://nateo.discovery.cs.vt.edu?question=text
text = value from QUESTION table (MariaDB)
*/
package main

import (
	"database/sql"
	"fmt"
	// "encoding/json"
	"log"
	"net/http"

	"example.com/database"
	"github.com/gorilla/mux"
)
var db *sql.DB
type Response struct {
	Message string `json:"message"`
}

//main 
func main(){
	log.Println("Chatbot Service POST")
	output := database.Test("yo")
	db = database.ConnectDB()
	fmt.Println(output)
	
	router := mux.NewRouter()
	// router.HandleFunc("/", getHandler).Methods("GET")

	log.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", router)
}


// //GET handler
// func getHandler(writer http.ResponseWriter, r *http.Request){
// 	//set response type to json
// 	queryParams := r.URL.Query()
// 	question := queryParams.Get("question")
//
// 	writer.Header().Set("Content-Type", "application/json")
// 	//get from database
// 	dbResponse := getDatabase(question)
// 	resp := Response{Message: dbResponse}
//
// 	//encode and send json
// 	json.NewEncoder(writer).Encode(resp)
// }
//
