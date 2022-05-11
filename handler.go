package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func timer(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello world")
	message := "This HTTP triggered function executed successfully. Pass a name in the query string for a personalized response.\n"
	name := r.URL.Query().Get("name")
	if name != "" {
		message = fmt.Sprintf("Hello, %s. This HTTP triggered function executed successfully.\n", name)
	}
	fmt.Fprint(w, message)
}

func main() {
	listenAddr := ":8080"
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		listenAddr = ":" + val
	}
	http.HandleFunc("/TimerExample", timer)
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}
