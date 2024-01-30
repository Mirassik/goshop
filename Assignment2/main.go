package main

import (
	"Assignment2/handlers"
	"Assignment2/middlewares"
	"Assignment2/models"
	"Assignment2/routes"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var templatesDir = filepath.Join("html", "templates")

func main() {
	db = SetupDatabase()

	handlers.SetupAuthHandler(db)

	router := mux.NewRouter()
	routes.SetupRoutes(router)
	middlewares.InitHTMLRendering(templatesDir)

	router.Use(middlewares.HTMLRenderingMiddleware)

	// START HTML FILES ROUTING
	router.HandleFunc("/login.html", func(w http.ResponseWriter, r *http.Request) {
		err := middlewares.Templates.ExecuteTemplate(w, "login.html", nil)
		if err != nil {
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
			return
		}
	}).Methods("GET")
	router.HandleFunc("/register.html", func(w http.ResponseWriter, r *http.Request) {
		err := middlewares.Templates.ExecuteTemplate(w, "register.html", nil)
		if err != nil {
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
			return
		}
	}).Methods("GET")

	router.HandleFunc("/index.html", func(w http.ResponseWriter, r *http.Request) {
		err := middlewares.Templates.ExecuteTemplate(w, "index.html", nil)
		if err != nil {
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
			return
		}
	}).Methods("GET")

	router.HandleFunc("/logout.html", func(w http.ResponseWriter, r *http.Request) {
		err := middlewares.Templates.ExecuteTemplate(w, "logout.html", nil)
		if err != nil {
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
			return
		}
	}).Methods("GET")
	// END HTML FILES ROUTING

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	fmt.Println("Server is running on port", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func SetupDatabase() *gorm.DB {
	dsn := "user=postgres password=1234 dbname=goshop sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		return nil
	}

	return db
}
