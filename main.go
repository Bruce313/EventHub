package main

import(
	"net/http"
	"log"
)

func mainServer(res http.) {
	
}

func main() {
	http.HandleFunc("/", mainServer)
	err := http.ListenAndServe(":4000", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}	
}
