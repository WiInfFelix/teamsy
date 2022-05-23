package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"log"
	"net/http"
	"os"
	"teamsy/graph"
	"teamsy/graph/generated"
	"teamsy/internal/pkg/db"
	"teamsy/internal/pkg/logger"
	"teamsy/internal/pkg/organisations"
	"teamsy/internal/pkg/users"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	//router := chi.NewRouter()
	//router.Use(auth.AuthMiddleware())

	err := logger.InitLogger()
	if err != nil {
		log.Fatalf("There was an error initializing the logger: %s", err)
	}

	err = db.InitDB()
	if err != nil {
		log.Fatalf("Error connecting to the db: %s", err)
	}

	log.Println("Attempting table migration...")
	err = db.Db.AutoMigrate(&organisations.Organisation{}, &users.User{})
	if err != nil {
		log.Fatalf("There was an error migrating the tables: %s", err)
	}
	log.Println("Table migration successful!")

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
