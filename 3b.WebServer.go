//Web Server

package main
import (
	"fmt"
 	"net/http"
	"log"
)

func main(){
	//Define the route handler function
	handler := func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w,"Hello World!!")
	}

	//Cretae a new server mux
	mux := http.NewServeMux()

	//Register the route handler function to handle all requests to "/"
	mux.HandleFunc("/",handler)

	//Create a new HTTP server
	server := http.Server{
		Addr: ":8080",
		Handler: mux,
	}

	//Start the HTTP server
	log.Println("Server listening on port 8080....")
	log.Fatal(server.ListenAndServe())
}

// output:
// 2023/09/10 05:58:54 Server listening on port 8080....
