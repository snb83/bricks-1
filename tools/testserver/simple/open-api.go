// Code generated by github.com/pace/bricks DO NOT EDIT.
package simple

import (
	"context"
	errors1 "errors"
	mux "github.com/gorilla/mux"
	opentracing "github.com/opentracing/opentracing-go"
	errors "github.com/pace/bricks/maintenance/errors"
	metrics "github.com/pace/bricks/maintenance/metric/jsonapi"
	"net/http"
)

/*
GetTestHandler handles request/response marshaling and validation for
 Get /beta/test
*/
func GetTestHandler(service Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer errors.HandleRequest("GetTestHandler", w, r)

		// Trace the service function handler execution
		handlerSpan, ctx := opentracing.StartSpanFromContext(r.Context(), "GetTestHandler")
		defer handlerSpan.Finish()

		// Setup context, response writer and request type
		writer := getTestResponseWriter{
			ResponseWriter: metrics.NewMetric("simple", "/beta/test", w, r),
		}
		request := GetTestRequest{
			Request: r.WithContext(ctx),
		}

		// Scan and validate incoming request parameters

		// Invoke service that implements the business logic
		err := service.GetTest(ctx, &writer, &request)
		select {
		case <-ctx.Done():
			if ctx.Err() != nil {
				// Context cancellation should not be reported if it's the request context
				w.WriteHeader(499)
				if err != nil && !(errors1.Is(err, context.Canceled) || errors1.Is(err, context.DeadlineExceeded)) {
					// Report unclean error handling (err != context err) to sentry
					errors.Handle(ctx, err)
				}
			}
		default:
			if err != nil {
				errors.HandleError(err, "GetTestHandler", w, r)
			}
		}
	})
}

/*
GetTestResponseWriter is a standard http.ResponseWriter extended with methods
to generate the respective responses easily
*/
type GetTestResponseWriter interface {
	http.ResponseWriter
	OK()
}
type getTestResponseWriter struct {
	http.ResponseWriter
}

// OK responds with empty response (HTTP code 200)
func (w *getTestResponseWriter) OK() {
	w.Header().Set("Content-Type", "application/vnd.api+json")
	w.WriteHeader(200)
}

/*
GetTestRequest is a standard http.Request extended with the
un-marshaled content object
*/
type GetTestRequest struct {
	Request *http.Request `valid:"-"`
}

// Service interface for all handlers
type Service interface {
	// GetTest Test
	GetTest(context.Context, GetTestResponseWriter, *GetTestRequest) error
}

/*
Router implements: PACE Payment API

Welcome to the PACE Payment API documentation.
This API is responsible for managing payment methods for users as well as authorizing payments on behalf of PACE services.
*/
func Router(service Service) *mux.Router {
	router := mux.NewRouter()
	// Subrouter s1 - Path: /pay
	s1 := router.PathPrefix("/pay").Subrouter()
	s1.Methods("GET").Path("/beta/test").Handler(GetTestHandler(service)).Name("GetTest")
	return router
}