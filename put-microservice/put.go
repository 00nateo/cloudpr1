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

type PutRequest struct {
	Question string `json:"question"`
	Reply    string `json:"reply"`
}

//main 
func main(){
	log.Println("PUT Microservice")
	db = connectDB()
	router := mux.NewRouter()
	router.HandleFunc("/", putHandler).Methods("PUT")

	log.Println("Server listening on port 8082")
	http.ListenAndServe(":8082", router)
}

//PUT handler
func putHandler(writer http.ResponseWriter, r *http.Request){
	//set response type to json
	writer.Header().Set("Content-Type", "application/json")
	
	var putReq PutRequest
	err := json.NewDecoder(r.Body).Decode(&putReq)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		resp := Response{Message: "Invalid JSON format"}
		json.NewEncoder(writer).Encode(resp)
		return
	}

	dbResponse := putDB(putReq)

	//response from database
	resp := Response{Message: dbResponse}

	//encode and send json
	json.NewEncoder(writer).Encode(resp)
}

func putDB(body PutRequest) string{
	log.Println("=====Database PUT=====")
	question := body.Question
	reply := body.Reply

	//update database
	query := "UPDATE chatbot_hints SET reply = ? WHERE question = ?"
	result, err := db.Exec(query, reply, question)
	if err != nil {
		log.Println(err)
		return "Did not update data"
	}
	
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return "Error checking update result"
	}
	
	if rowsAffected == 0 {
		return "Input params not found"
	}
	
	return "Data updated "
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

