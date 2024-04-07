package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sAnjum55/server/internal/database"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load()
	var port = os.Getenv("PORT")
	if port == "" {
		log.Fatal("Port does not exist")
	}

	var dbUrl = os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatal("No DB url found")
	}

	conn, error := sql.Open("postgres", dbUrl)
	if error != nil {
		log.Fatal("Cannot coonect to databses", error)
	}

	queries := database.New(conn)

	apiCfg := apiConfig{
		DB: queries,
	}

	var router = chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()

	v1Router.Get("/healthz", handleRoutes)
	v1Router.Get("/error", handleError)
	v1Router.Post("/users", apiCfg.createUserHandler)
	v1Router.Get("/users", apiCfg.middlewareAuth(apiCfg.getUserHandler))
	v1Router.Post("/feed", apiCfg.middlewareAuth(apiCfg.createFeedHandler))
	v1Router.Get("/feed", apiCfg.getFeedsHandler)
	v1Router.Post("/followfeed", apiCfg.middlewareAuth(apiCfg.createFollowFeedHandler))
	v1Router.Get("/followfeed", apiCfg.middlewareAuth(apiCfg.getFollowFeedHandler))
	v1Router.Delete("/followfeed/{feedFollowID}", apiCfg.middlewareAuth(apiCfg.deleteFollowFeedForUser))
	router.Mount("/v1", v1Router)

	server := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}
	fmt.Printf("Server is running at port: %v", port)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
