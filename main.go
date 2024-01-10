package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

type apiConfig struct {
	apiKey  string
	baseUrl string
}

func main() {
	godotenv.Load("Key.env")
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatalln("PORT enviroment is not set")
	}

	baseUrl := os.Getenv("BASE_URL")
	apiKey := os.Getenv("API_KEY")

	cfg := apiConfig{
		apiKey:  apiKey,
		baseUrl: baseUrl,
	}

	router := chi.NewRouter()
	v1Router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	router.Handle("/", http.FileServer(http.Dir(".")))

	v1Router.Get("/validate", cfg.handlerValidateCC)
	v1Router.Get("/accountrng", cfg.handlerAccountRange)

	router.Mount("/v1", v1Router)

	srv := http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	log.Printf("Serving on port: %s\n", port)
	log.Fatal(srv.ListenAndServe())
}
