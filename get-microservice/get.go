package main
import (
	"database/sql"
	"encoding/json"
	"net/http"
	"log"
	"example.com/database"
	"github.com/gorilla/mux"
)
var db *sql.DB
type Response struct {
	Message string `json:"message"`
}

//main 
func main(){
	log.Println("Chatbot Service Adminer")
	db = database.ConnectDB()
	router := mux.NewRouter()
	router.HandleFunc("/", getHandler).Methods("GET")

	log.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", router)
}

//GET handler
func getHandler(writer http.ResponseWriter, r *http.Request){
	//set response type to json
	writer.Header().Set("Content-Type", "application/json")

	queryParams := r.URL.Query()
	question := queryParams.Get("question")

	dbResponse := database.GetDB(question)
	//get from database
	resp := Response{Message: dbResponse}

	//encode and send json
	json.NewEncoder(writer).Encode(resp)
}
