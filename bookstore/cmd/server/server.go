package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
	"github.com/janapc/graphql-books-go/bookstore/graph"
	"github.com/janapc/graphql-books-go/bookstore/internal/infra/database"
	"github.com/janapc/graphql-books-go/bookstore/internal/infra/service"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	port := os.Getenv("PORT_BOOKSTORE")

	db, err := database.Connection()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	bookDB := database.NewBookDatabase(db)
	authorDB := database.NewAuthorDatabase(db)
	router := chi.NewRouter()
	router.Use(cors.AllowAll().Handler)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		BookDB:   bookDB,
		AuthorDB: authorDB,
	}}))

	router.Handle("/", playground.Handler("bookstore playground", "/bookstore"))
	router.Route("/bookstore", func(r chi.Router) {
		r.Use(service.MiddlewareAuth)
		r.Handle("/", srv)
	})

	log.Printf("connect to http://localhost:%s/bookstore for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
