// package main

// import (
// 	"graphql-demo/graph"
// 	"log"
// 	"net/http"
// 	"os"

// 	"github.com/99designs/gqlgen/graphql/handler"
// 	"github.com/99designs/gqlgen/graphql/handler/extension"
// 	"github.com/99designs/gqlgen/graphql/handler/lru"
// 	"github.com/99designs/gqlgen/graphql/handler/transport"
// 	"github.com/99designs/gqlgen/graphql/playground"
// 	"github.com/vektah/gqlparser/v2/ast"
// )

// const defaultPort = "8080"

// func main() {
// 	port := os.Getenv("PORT")
// 	if port == "" {
// 		port = defaultPort
// 	}

// 	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

// 	srv.AddTransport(transport.Options{})
// 	srv.AddTransport(transport.GET{})
// 	srv.AddTransport(transport.POST{})

// 	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

// 	srv.Use(extension.Introspection{})
// 	srv.Use(extension.AutomaticPersistedQuery{
// 		Cache: lru.New[string](100),
// 	})

// 	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
// 	http.Handle("/query", srv)

//		log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
//		log.Fatal(http.ListenAndServe(":"+port, nil))
//	}
package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"graphql-demo/graph"
	generated "graphql-demo/graph/generated"
	"graphql-demo/graph/model"
	"log"
	"net/http"
)

const defaultPort = "8080"

func main() {
	dsn := "host=localhost user=postgres password=root@123 dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto-migrate schema
	db.AutoMigrate(&model.User{}, &model.Post{})

	resolver := &graph.Resolver{DB: db}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", defaultPort)
	log.Fatal(http.ListenAndServe(":"+defaultPort, nil))
}
