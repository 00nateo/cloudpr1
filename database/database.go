package database

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	_ "github.com/go-sql-driver/mysql"
)
var db *sql.DB
func Checkdb() string{
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

	return "checkdb"

}
func GetDB(question string) string{
	fmt.Println("=====Database GET=====")
	db=ConnectDB()

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
func PostDB(body url.Values) {
	fmt.Println("=====Database POST=====")
	question := body.Get("question")
	reply := body.Get("reply")

	fmt.Println(question)
	fmt.Println(reply)
	//add to database
	
}

func PutDB(body url.Values) {
	fmt.Println("=====Database PUT=====")
	question := body.Get("question")
	reply := body.Get("reply")

	fmt.Println(question)
	fmt.Println(reply)
	//update database
	
}
func DeleteDB(body url.Values) {
	fmt.Println("=====Database DELETE=====")
	question := body.Get("question")
	reply := body.Get("reply")

	fmt.Println(question)
	fmt.Println(reply)
	//delete from database
	
}

func ConnectDB() *sql.DB{
	//connect with Unix socket 
	// connectionString := "root@unix(/run/mysqld/mysqld.sock)/chatbot"
	//Local
	// connectionString := "root:ChangeMe@tcp(nateodb:3306)/chatbot"
	connectionString := "root:ChangeMe@tcp(localhost:3306)/chatbot"
	
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

