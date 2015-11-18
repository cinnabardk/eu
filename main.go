package main
import (
	"net/http"
	"log"
)


const port = "8080"
func main(){
	log.Println("Server now starting on port:", port)
	log.Fatal(http.ListenAndServe(":"+port, http.FileServer(http.Dir("assets"))))
}