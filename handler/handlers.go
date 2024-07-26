package handler

import (
	"jobPostGraphQl/driver"
	"jobPostGraphQl/graph"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
)

func GraphqlHandler() gin.HandlerFunc {
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		DB: driver.DBInstance,
	}}))

	return func(ctx *gin.Context) {
		srv.ServeHTTP(ctx.Writer, ctx.Request)
	}
}

func PlaygroundHandler() gin.HandlerFunc {
	handler := playground.Handler("GraphQL playground", "/query")

	return func(ctx *gin.Context) {
		handler.ServeHTTP(ctx.Writer, ctx.Request)
	}
}
