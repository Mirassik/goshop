package routes

import (
	"Assignment2/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func SetupRoutes(r *mux.Router) {
	// localhost:8080/api/register
	apiRouter := r.PathPrefix("/api").Subrouter()

	apiRouter.HandleFunc("/register", handlers.Register).Methods("POST")
	apiRouter.HandleFunc("/login", handlers.Login).Methods("POST")
	apiRouter.HandleFunc("/logout", handlers.Logout).Methods("GET")
	apiRouter.HandleFunc("/users/{id}", handlers.UpdateUser).Methods("PUT")
	apiRouter.HandleFunc("/users/{id}", handlers.DeleteUser).Methods("DELETE")

	apiRouter.Handle("/protected", handlers.TokenVerificationMiddleware(http.HandlerFunc(protectedEndpoint))).Methods("GET")
}

func protectedEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is a protected endpoint"))
}
