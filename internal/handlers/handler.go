package handlers

import (
    "net/http"
    "github.com/gorilla/mux"
    "thefinisher-goapp/internal/services"
)

func SetupRouter() *mux.Router {
    router := mux.NewRouter()

    // Routes
    router.HandleFunc("/users", GetUsers).Methods("GET")
    router.HandleFunc("/user/{id}", GetUserByID).Methods("GET")

    return router
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
    users := services.GetAllUsers()
    w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type", "application/json")
    w.Write([]byte(users))
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    user := services.GetUser(id)
    w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type", "application/json")
    w.Write([]byte(user))
}

