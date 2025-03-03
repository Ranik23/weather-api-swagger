// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"context"
	"crypto/tls"
	"log"
	"log/slog"
	"net"
	"net/http"
	"time"

	"weatherbot/api/restapi/operations"
	"weatherbot/config"
	"weatherbot/internal/handlers"
	"weatherbot/internal/middlewares"
	"weatherbot/internal/usecase"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

//go:generate swagger generate server --target ../../api --name WeatherAPI --spec ../../docs/api.json --principal interface{}

func configureFlags(api *operations.WeatherAPIAPI) {
	api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{
		
	}
}

func configureAPI(api *operations.WeatherAPIAPI) http.Handler {
	
	defer func() {
		if err := recover(); err != nil {
			slog.Error("panic caught", slog.Any("error", err))
		}
	}()


	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	// api.Logger = slog.Debug

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	usecase := usecase.NewUseCaseImpl(cfg)

	requests := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "endpoint"},
	)

	duration := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_duration_seconds",
			Help:    "Histogram of HTTP request durations",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "endpoint"},
	)

	prometheus.MustRegister(requests)
	prometheus.MustRegister(duration)

	api.GetWeatherHandler = handlers.NewWeatherForecastHandler(usecase, duration, requests)

	//api.HandlerFor("/metrics", promhttp.Handler())

		
	api.PreServerShutdown = func() {
		api.Logger("Shutting Down The Server!")
	}

	api.ServerShutdown = func() {
		api.Logger("Server Is ShutDown! Done Handling All The Connections")
	}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
	s.ConnContext = func(ctx context.Context, c net.Conn) context.Context {
		ctxTimeout, cancel := context.WithTimeout(ctx, 10 * time.Second)
		defer cancel()
		return ctxTimeout
	}
}



// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/metrics", promhttp.Handler())
	mux.Handle("/", middlewares.PanicMiddleWare(middlewares.LoggingMiddleware(handler)))
	
	return mux
}

