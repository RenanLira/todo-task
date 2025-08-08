package main

import (
	"net/http"
	"todo-tasks/graph"

	"todo-tasks/graph/directives"
	"todo-tasks/graph/generated"
	"todo-tasks/internal/domain/auth/middlewares"
	auth_service "todo-tasks/internal/domain/auth/services"
	todo_service "todo-tasks/internal/domain/todos/services"
	user_service "todo-tasks/internal/domain/users/services"
	"todo-tasks/internal/internalerrors"
	customvalidators "todo-tasks/internal/internalerrors/custom-validators"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/vektah/gqlparser/v2/ast"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
}

func init() {
	internalerrors.Validate.RegisterValidation("is_unique_field", customvalidators.IsUniqueField)
}

func graphQLHandler() http.Handler {

	c := generated.Config{Resolvers: &graph.Resolver{
		TodoService:  todo_service.NewTodoService(),
		UserResolver: user_service.NewUserService(),
		AuthResolver: auth_service.NewAuthService(),
	}}

	c.Directives.Authenticated = directives.AuthDirective

	srv := handler.New(generated.NewExecutableSchema(c))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	return srv
}

func main() {
	r := chi.NewRouter()

	r.Route("/api", func(r chi.Router) {
		r.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello, World!"))
		})
	})

	r.Route("/query", func(r chi.Router) {
		r.Use(middlewares.AuthMiddleware())
		r.Handle("/playground", playground.Handler("GraphQL playground", "/query"))
		r.Handle("/", graphQLHandler())
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}

}
