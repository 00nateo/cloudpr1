package main
import (
	"database/sql"
	// "encoding/json"
	"net/http"
	"log"
	// "fmt"
	"github.com/gorilla/mux"
	"example.com/database"
)
var db *sql.DB
type Response struct {
	Message string `json:"message"`
}

//main 
func main(){
	log.Println("Chatbot Service Adminer")
	
	router := mux.NewRouter()
	router.HandleFunc("/", putHandler).Methods("PUT")

	log.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", router)
}



//GET handler
func putHandler(writer http.ResponseWriter, r *http.Request){
	// //set response type to json
	// queryParams := r.URL.Query()
	// question := queryParams.Get("question")
	//
	// writer.Header().Set("Content-Type", "application/json")
	// //get from database
	// dbResponse := getDatabase(question)
	// resp := Response{Message: dbResponse}
	//
	// //encode and send json
	// json.NewEncoder(writer).Encode(resp)
}

