package main

import (
	"fmt"
	"log"
	"net/http"
)
func handler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Hi Beesh, this app version 2")
}
func main() {
	// fmt.Println("hi beesh")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
