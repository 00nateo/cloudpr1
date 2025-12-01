package main
import (
	"database/sql"
	"encoding/json"
	"net/http"
	"log"
	"fmt"
	"net/url"
	"github.com/gorilla/mux"
	_ "github.com/go-sql-driver/mysql"
)
var db *sql.DB
type Response struct {
	Message string `json:"message"`
}

//main 
func main(){
	log.Println("GET Microservice")
	db = connectDB()
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

	dbResponse := getDB(queryParams)
	fmt.Println(queryParams)

	//get from database
	resp := Response{Message: dbResponse}

	//encode and send json
	json.NewEncoder(writer).Encode(resp)
}

func getDB(body url.Values) string{
	log.Println("=====Database GET=====")
	question := body.Get("question")
	var reply string

	//query database
	query := "SELECT reply FROM chatbot_hints WHERE question = ?"
	err := db.QueryRow(query, question).Scan(&reply)
	if err != nil {
		if err == sql.ErrNoRows {
			return "Sorry not be able to understand you"
		} 
	}
	
	return reply
	
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
