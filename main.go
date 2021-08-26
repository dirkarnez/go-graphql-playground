package main

import (
	"log"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/dirkarnez/go-graphql-playground/graph"
	"github.com/dirkarnez/go-graphql-playground/graph/generated"
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()
	
	app.Any("/graphql", iris.FromStd(playground.Handler("GraphQL playground", "/query")))
	app.Any("/query", iris.FromStd(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))))
	
	app.HandleDir("/", "./public")

	err := app.Run(
		// Start the web server at localhost:8080
		iris.Addr(":9000"),
		// skip err server closed when CTRL/CMD+C pressed:
		iris.WithoutServerError(iris.ErrServerClosed),
		// enables faster json serialization and more:
		iris.WithOptimizations,
	)

	if err != nil {
		log.Println(err.Error())
	}
}