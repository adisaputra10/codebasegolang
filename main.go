package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"gitlab.playcourt.id/new-mypertamina/myptm-external-insurance-service/src/helper/wrapper"
	"gitlab.playcourt.id/new-mypertamina/myptm-external-insurance-service/src/helper/errors"
	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload"
	"github.com/rs/cors"

)

func main() {
	port := os.Getenv("PORT")
	appName := os.Getenv("APP_NAME")
	

	// init auth
	

	router := mux.NewRouter()


	// set default headers
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Server", appName)
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods",
				"GET, PUT, POST, DELETE")
			w.Header().Set("Access-Control-Allow-Headers",
				"Origin, X-Requested-With, Content-Type")
			next.ServeHTTP(w, r)
		})
	})

	// set not found resource handler
	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		result := new(wrapper.Result)
		result.Err = errors.ErrNotFound
		result.Message = "Resource is not found"

		wrapper.HTTPResponseJSON(w, result, http.StatusNotFound, "")
	})

	// set method not allowed handler
	router.MethodNotAllowedHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		result := new(wrapper.Result)
		result.Err = errors.ErrMethodNotAllowed
		result.Message = "Method is not allowed"

		wrapper.HTTPResponseJSON(w, result, http.StatusMethodNotAllowed, "")
	})

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		result := new(wrapper.Result)
		wrapper.HTTPResponseJSON(w, result, http.StatusOK, "Application is running properly")
	}).Methods("GET")


	// set cors handler
	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"X-Requested-With", "Origin", "Content-Type", "Authorization"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
	})

	generalHandler := corsHandler.Handler(router)

	log.Printf("Application is running on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), generalHandler))
}


