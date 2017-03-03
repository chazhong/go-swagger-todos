package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	graceful "github.com/tylerb/graceful"

	"github.com/chazhong/go-swagger-todos/server/restapi/controllers"
	"github.com/chazhong/go-swagger-todos/server/restapi/operations"
	"github.com/chazhong/go-swagger-todos/server/restapi/operations/todos"
)

// This file is safe to edit. Once it exists it will not be overwritten

//go:generate swagger generate server --target ../gen --name TodoList --spec ../swagger.yml

func configureFlags(api *operations.TodoListAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.TodoListAPI) http.Handler {
	api.ServeError = errors.ServeError

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// Controllers
	todoListController := controllers.TodoListController{}

	// APIs
	api.TodosGetTodosHandler = todos.GetTodosHandlerFunc(todoListController.GetTodoList)
	api.TodosAddOneHandler = todos.AddOneHandlerFunc(todoListController.AddTodoItem)

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *graceful.Server, scheme string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
