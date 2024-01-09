package main

import (
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

type apiConfig struct {
	pathToCerts string
	consumerKey string
	keyPass     string
	url         *url.URL
}

func main() {
	godotenv.Load("Key.env")
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatalln("PORT enviroment is not set")
	}

	pthCerts := os.Getenv("CERTS_PATH")
	cnsmrKey := os.Getenv("CONSUMER_KEY")
	keyPass := os.Getenv("KEY_PASSWORD")
	strUrl := os.Getenv("URL")
	parsedUrl, _ := url.Parse(strUrl)

	cfg := apiConfig{
		pathToCerts: pthCerts,
		consumerKey: cnsmrKey,
		keyPass:     keyPass,
		url:         parsedUrl,
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

	router.Mount("/v1", v1Router)

	srv := http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	log.Printf("Serving on port: %s\n", port)
	log.Fatal(srv.ListenAndServe())
}
