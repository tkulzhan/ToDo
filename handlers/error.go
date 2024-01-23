package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, status int, err error) {
	w.WriteHeader(status)
	log.Println("ERROR status: " + strconv.Itoa(status))
	// Increment error count in prometheus metric
	HttpRequestErrors.WithLabelValues("ERROR").Inc()

	if status == http.StatusNotFound {
		fmt.Fprint(w, "404 page not found")
	} else if status == http.StatusBadRequest {
		fmt.Fprint(w, "400 bad request")
	} else if status == http.StatusInternalServerError {
		fmt.Fprint(w, "500 internal server error")
	} else if status == http.StatusMethodNotAllowed {
		fmt.Fprint(w, "405 method not allowed")
	}
}

// HttpRequestErrors is the prometheus counter metric
var HttpRequestErrors = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "x_http_request_errors_total",
		Help: "Total number of HTTP requests resulting in errors",
	},
	[]string{"type"},
)
