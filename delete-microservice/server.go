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
	log.Println("Chatbot Service Adminer")
	db = connectDB()
	defer db.Close()
	checkdb()
	
	router := mux.NewRouter()
	router.HandleFunc("/", getHandler).Methods("GET")

	log.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", router)
}


//querys database and returns response
func getDatabase(question string) string{
	log.Println(question)
	var reply string
	query := "SELECT reply FROM chatbot_hints WHERE question = ?"
	
	err := db.QueryRow(query, question).Scan(&reply)
	log.Println(reply)
	if err != nil {
		if err == sql.ErrNoRows {
			return "Sorry not be able to understand you"
		} 
	}
	
	return reply
}
func checkdb(){
	log.Println("Checking database....")
	var count int
	var test string 
	err := db.QueryRow("SELECT COUNT(*) FROM chatbot_hints").Scan(&count)
	if err != nil {
		log.Println("Error counting rows:", err)
	}
	fmt.Println("Number of rows:", count)
	db.QueryRow("SELECT reply FROM chatbot_hints WHERE question = 'HI'").Scan(&test)
	fmt.Println("Test: ", test)

}



//GET handler
func getHandler(writer http.ResponseWriter, r *http.Request){
	//set response type to json
	queryParams := r.URL.Query()
	question := queryParams.Get("question")

	writer.Header().Set("Content-Type", "application/json")
	//get from database
	dbResponse := getDatabase(question)
	resp := Response{Message: dbResponse}

	//encode and send json
	json.NewEncoder(writer).Encode(resp)
}

func connectDB() *sql.DB{
	//connect with Unix socket 
	// connectionString := "root@unix(/run/mysqld/mysqld.sock)/chatbot"
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
	}

	
	log.Println("Connected to database!")
	return db
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
