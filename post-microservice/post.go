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

type PostRequest struct {
	Question string `json:"question"`
	Reply    string `json:"reply"`
}

//main 
func main(){
	log.Println("POST Microservice")
	db = connectDB()
	router := mux.NewRouter()
	router.HandleFunc("/", postHandler).Methods("POST")

	log.Println("Server listening on port 8081")
	http.ListenAndServe(":8081", router)
}

//POST handler
func postHandler(writer http.ResponseWriter, r *http.Request){
	//set response type to json
	writer.Header().Set("Content-Type", "application/json")
	
	var postReq PostRequest
	err := json.NewDecoder(r.Body).Decode(&postReq)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		resp := Response{Message: "Invalid JSON format"}
		json.NewEncoder(writer).Encode(resp)
		return
	}

	dbResponse := postDB(postReq)

	//response from database
	resp := Response{Message: dbResponse}

	//encode and send json
	json.NewEncoder(writer).Encode(resp)
}

func postDB(body PostRequest) string{
	log.Println("=====Database POST=====")
	question := body.Question
	reply := body.Reply

	//insert into database
	query := "INSERT INTO chatbot_hints (question, reply) VALUES (?, ?)"
	_, err := db.Exec(query, question, reply)
	if err != nil {
		log.Println(err)
		return "Didnot to insert data"
	}
	
	return "Data inserted"
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
