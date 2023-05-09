package routes

import (
    "github.com/gorilla/mux"
)

func Router() *mux.Router {
    // initialize router
    router := mux.NewRouter()

    // set up routes
    //router.HandleFunc("/", handlers.Home).Methods("GET")
    //router.HandleFunc("/login", handlers.Login).Methods("POST")
    //router.HandleFunc("/protected", handlers.Protected).Methods("GET")

    return router
}
