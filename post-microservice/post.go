//create information in the database
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
	// database.ConnectDB()
	database.Test("hola")
	
	router := mux.NewRouter()
	router.HandleFunc("/", postHandler).Methods("POST")

	log.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", router)
}


//syntax
//POST <request-target>["?"<query>]
func postHandler(writer http.ResponseWriter, r *http.Request){
	//set response type to json
	// queryParams := r.URL.Query()
	err := r.ParseForm()
	if err != nil{
		fmt.Println("Error: ", err)
	}

	database.PostDB(r.PostForm)
	database.GetDB(r.PostForm)
	database.PutDB(r.PostForm)
	database.DeleteDB(r.PostForm)
//	question := queryParams.Get("question")

	writer.Header().Set("Content-Type", "application/json")
	//write to db??
	// dbResponse := getDatabase(question)
	//
	// resp := Response{Message: dbResponse}
	//
	// //encode and send json
	// json.NewEncoder(writer).Encode(resp)
}

