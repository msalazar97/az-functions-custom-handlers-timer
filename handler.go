package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func TimerHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Println("Timer triggered!")
	w.Write([]byte("{}"))
}

func main() {
	customHandlerPort, exists := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT")
	if exists {
		fmt.Println("FUNCTIONS_CUSTOMHANDLER_PORT: " + customHandlerPort)
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/TimerExample", TimerHandler)
	log.Fatal(http.ListenAndServe(":"+customHandlerPort, mux))
}
