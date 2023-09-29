package main

import (
	"encoding/json"
	"net/http"
)

var message = "Hello, World!"

func getMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": message})
}

func setMessage(w http.ResponseWriter, r *http.Request) {
	var newMessage map[string]string
	if err := json.NewDecoder(r.Body).Decode(&newMessage); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	message = newMessage["message"]
	w.WriteHeader(http.StatusNoContent)
}

func main() {
	http.HandleFunc("/message", getMessage)
	http.HandleFunc("/message/set", setMessage)

	http.ListenAndServe(":8080", nil)
}
