package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"go-chi-rest-api-boilerplate/controller"
	"go-chi-rest-api-boilerplate/pkg/utils"

	"github.com/joho/godotenv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading environment variables: %s", err.Error())
	}

	if os.Getenv("SM") != "" {
		var env map[string]string
		json.Unmarshal([]byte(os.Getenv("SM")), &env)
		for k, v := range env {
			os.Setenv(k, v)
		}
	}

	r := chi.NewRouter()
	port := os.Getenv("GO_PORT")

	if port == "" {
		log.Fatal("Port not set in .env file")
	}

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"*"},
		AllowCredentials: false,
		MaxAge:           300,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	}))

	r.Use(setCORSHeaders)
	r.Use(middleware.Recoverer)
	r.Get("/health", apiToGetHealth)
	r.Route("/user", func(r chi.Router) {
		r.Post("/create", controller.CreateUser)
		r.Post("/update/{userId}", controller.UpdateUser)
	})
	r.Get("/users", controller.GetAllUsers)

	addr := fmt.Sprintf(":%s", port)
	log.Printf("Server started on port: %v\n", port)
	if err := http.ListenAndServe(addr, r); err != nil {
		utils.LogError("Server error: %s", "ServerError", 5, err, nil)
		log.Fatalf("Server error: %s", err)
	}
}

func setCORSHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "*")
			w.Header().Set("Access-Control-Max-Age", "300")
			w.Header().Set("Access-Control-Expose-Headers", "*")
			w.Header().Set("Access-Control-Allow-Credentials", "false")
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func apiToGetHealth(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
