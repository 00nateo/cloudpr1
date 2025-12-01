package main
import (
	"database/sql"
	"encoding/json"
	"net/http"
	"log"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Response struct {
	Message string `json:"message"`
}

type DeleteRequest struct {
	Question string `json:"question"`
}

//main 
func main(){
	log.Println("DELETE Microservice")
	db = connectDB()
	router := mux.NewRouter()
	router.HandleFunc("/", deleteHandler).Methods("DELETE")

	log.Println("Server listening on port 8083")
	http.ListenAndServe(":8083", router)
}

//DELETE handler
func deleteHandler(writer http.ResponseWriter, r *http.Request){
	//set response type to json
	writer.Header().Set("Content-Type", "application/json")
	
	var deleteReq DeleteRequest
	err := json.NewDecoder(r.Body).Decode(&deleteReq)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		resp := Response{Message: "Invalid JSON format"}
		json.NewEncoder(writer).Encode(resp)
		return
	}

	dbResponse := deleteDB(deleteReq)

	//response from database
	resp := Response{Message: dbResponse}

	//encode and send json
	json.NewEncoder(writer).Encode(resp)
}

func deleteDB(body DeleteRequest) string{
	log.Println("=====Database DELETE=====")
	question := body.Question

	//delete from database
	query := "DELETE FROM chatbot_hints WHERE question = ?"
	result, err := db.Exec(query, question)
	if err != nil {
		log.Println(err)
		return "Did not delete data"
	}
	
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return "Error deleting"
	}
	
	if rowsAffected == 0 {
		return "Query not found"
	}
	
	return "Data deleted "
}

func connectDB() *sql.DB{
	//Kubernetes service connection
	connectionString := "root:ChangeMe@tcp(nateodb:3306)/chatbot"
	
	db, error := sql.Open("mysql", connectionString)
	if error != nil {
		log.Println("Failed to open database connection:", error)
		return nil
	}

	fmt.Println("Pinging....")
	error = db.Ping()
	if error != nil {
		log.Fatalln("Failed to ping database:", error)
		return nil
	} else{
		log.Println("Connected to database!")
	}
	return db
}

