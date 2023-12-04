package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/go-sql-driver/mysql"
	"github.com/janapc/graphql-books-go/user/graph"
	"github.com/janapc/graphql-books-go/user/internal/infra/database"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT_USER")

	db, err := database.Connection()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userDB := database.NewUserDatabase(db)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		UserDB: userDB,
	}}))
	http.Handle("/", playground.Handler("user api playground", "/user"))
	http.Handle("/user", srv)

	log.Printf("connect to http://localhost:%s/user", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
