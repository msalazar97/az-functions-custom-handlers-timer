package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func simpleTimer(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello world")
	message := "TIMER TRIGGER TESTING.\n"
	name := r.URL.Query().Get("name")
	if name != "" {
		message = fmt.Sprintf("Hello, %s. This HTTP triggered function executed successfully.\n", name)
	}
	fmt.Fprint(w, message)
}

func main() {
	customHandlerPort, exists := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT")
	if exists {
		fmt.Println("FUNCTIONS_CUSTOMHANDLER_PORT: " + customHandlerPort)
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/TimerExample", simpleTimer)
	log.Fatal(http.ListenAndServe(":"+customHandlerPort, mux))
}
