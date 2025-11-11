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
package database 
import (
	"database/sql"
	"log"
	"fmt"
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

func Test(yo string) string{
	return yo
}


func ConnectDB() *sql.DB{
	//connect with Unix socket 
	// connectionString := "root@unix(/run/mysqld/mysqld.sock)/chatbot"
	//Local
	// connectionString := "root:ChangeMe@tcp(nateodb:3306)/chatbot"
	connectionString := "root:ChangeMe@tcp(mariadb_container:3306)/chatbot"
	
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

