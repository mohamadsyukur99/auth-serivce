package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/mohamadsyukur99/auth-service/graph"
	"github.com/mohamadsyukur99/auth-service/graph/generated"
	database "github.com/mohamadsyukur99/auth-service/internal/pkg/db/mysql"
	setting "github.com/mohamadsyukur99/auth-service/internal/pkg/setting"
)

// const defaultPort = "8080"

func main() {
	setting, _ := setting.LoadConfiguration("setting.json")
	port := os.Getenv("PORT")
	if port == "" {
		port = setting.Listen.Port
	}

	router := chi.NewRouter() // menggunakan chi router

	// router.Use(auth.Middleware())

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	database.InitDB()

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router)) // use chi router here

	// http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	// http.Handle("/query", srv)
	// log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	// log.Fatal(http.ListenAndServe(":"+port, nil))
}
