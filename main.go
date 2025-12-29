package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

type User struct {
	Name string `json:"name"`
}

var UserCache = make(map[int]User)

var cacheMutex sync.RWMutex

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/", handleRoot)
	router.HandleFunc("POST /users", createUser)
	router.HandleFunc("GET /users/get/{id}", getUser)
	router.HandleFunc("DELETE /users/{id}", deleteUser)
 
	http.ListenAndServe(":4000", router)
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world from GO")
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	if user.Name == "" {
		http.Error(w, "name field is empty", http.StatusBadRequest)
		return
	}

	cacheMutex.Lock()
	UserCache[len(UserCache) + 1] = user
	cacheMutex.Unlock()

	w.WriteHeader(http.StatusNoContent)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return 
	}

	cacheMutex.RLock()
	user, ok := UserCache[id]
	cacheMutex.RUnlock()

	if !ok {
		http.Error(w, "user not found", http.StatusNotFound)
		return 
	}

	w.Header().Set("Content-Type", "application/json")
	b,err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if _, ok := UserCache[id]; !ok {
		http.Error(w, "user not found in DB", http.StatusNotFound)
		return
	}

	cacheMutex.Lock()
	delete(UserCache, id)
	cacheMutex.Unlock()

	w.WriteHeader(http.StatusNoContent)
}