// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"errors"
	"github.com/bgadrian/pseudoservice/handlers"
	"github.com/go-openapi/swag"
	"net/http"

	"github.com/bgadrian/pseudoservice/restapi/operations"
	swError "github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
)

//go:generate swagger generate server --target .. --name PseudoService --spec ../swagger.yaml

func configureFlags(api *operations.PseudoServiceAPI) {
	type key struct {
		// apikey tralalala
		Apikey string `long:"api-key" description:"token used to accept incoming requests" env:"APIKEY" default:"SECRET42"`
	}
	values := key{}
	opts := swag.CommandLineOptionsGroup{ShortDescription: "Security", LongDescription: "Security options", Options: &values}
	api.CommandLineOptionsGroups = append(api.CommandLineOptionsGroups, opts)

	api.ApikeyAuth = func(token string) (interface{}, error) {
		if token == values.Apikey {
			return true, nil
		} else {
			return nil, errors.New("invalid token (apikey)")
		}
	}
}

func configureAPI(api *operations.PseudoServiceAPI) http.Handler {
	// configure the api here
	api.ServeError = swError.ServeError
	api.JSONConsumer = runtime.JSONConsumer()
	api.JSONProducer = runtime.JSONProducer()

	h := handlers.MyHandlers{}
	h.Api = api

	api.GetHealthHandler = operations.GetHealthHandlerFunc(h.GetHealthHandler)
	api.GetUsersCountHandler = operations.GetUsersCountHandlerFunc(h.GetUsersHandler)

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
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	handler = handlers.Gzip(handler)
	handler = handlers.CustomHeaders(handler)
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
