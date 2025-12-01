package main
import (
	"net/http"
	"net/http/httputil"
	"log"
	"net/url"
	"github.com/gorilla/mux"
)

//main 
func main(){
	log.Println("Gateway Service")
	router := mux.NewRouter()
	router.HandleFunc("/", requestHandler)

	log.Println("Server listening on port 8084")
	http.ListenAndServe(":8084", router)
}

func requestHandler(writer http.ResponseWriter, r *http.Request){
	var target string 

	switch r.Method{
	case http.MethodGet:
		log.Println("Routing to GET service")
		target = "http://nateo-get-service:8080/"
	case http.MethodPost:
		log.Println("Routing to POST service")
		target = "http://nateo-post-service:8081/"
	case http.MethodPut:
		log.Println("Routing to PUT service")
		target = "http://nateo-put-service:8082/"
	case http.MethodDelete:
		log.Println("Routing to DELETE service")
		target = "http://nateo-delete-service:8083/"
	default:
		log.Printf("Unsupported HTTP Method: %s", r.Method)
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}


	proxyURL, err := url.Parse(target)
	if err != nil{
		log.Printf("Error parsing URL %s: %v", target, err)
		http.Error(writer, "Internal server error", http.StatusInternalServerError)
		return
	}
	proxy := httputil.NewSingleHostReverseProxy(proxyURL)
	proxy.ServeHTTP(writer, r)
}
