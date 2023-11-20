package main

import (
	"ToDo/handlers"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	uptimeStartTime = time.Now()
	uptime          = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "x_uptime",
			Help: "Uptime of the application in seconds",
		},
	)

	HttpRequestCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "x_http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method"},
	)

	HttpRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "x_http_request_duration_seconds",
			Help:    "Duration of HTTP requests in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method"},
	)

	HttpRequestErrors = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "x_http_request_errors_total",
			Help: "Total number of HTTP requests resulting in errors",
		},
		[]string{"method"},
	)

	HttpRequestLatency = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "x_http_request_latency_seconds",
			Help:    "Latency of HTTP requests in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method"},
	)

	xUp = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "x_up",
			Help: "Indicates whether the service is up (1) or down (0).",
		},
	)

	throughputMutex sync.Mutex
	throughput      = 0
	throughputTimer = time.NewTicker(1 * time.Second)
	errorRate       = 0.1
)

func prometheusHandler() gin.HandlerFunc {
	h := promhttp.Handler()

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func init() {
	prometheus.MustRegister(uptime)
	prometheus.MustRegister(HttpRequestCounter)
	prometheus.MustRegister(HttpRequestDuration)
	prometheus.MustRegister(HttpRequestErrors)
	prometheus.MustRegister(HttpRequestLatency)
	prometheus.MustRegister(xUp)
}

func recordUptime() {
	for {
		uptime.Set(time.Since(uptimeStartTime).Seconds())
		time.Sleep(1 * time.Second)
	}
}

func updateThroughput() {
	for range throughputTimer.C {
		throughputMutex.Lock()

		// Introduce 1% error rate
		if rand.Float64() < errorRate/100 {
			// Introduce errors
			HttpRequestCounter.WithLabelValues("GET").Inc()
			HttpRequestErrors.WithLabelValues("ERROR").Inc()
		}

		throughput = 0
		throughputMutex.Unlock()
	}
}

func getThroughput() int {
	throughputMutex.Lock()
	defer throughputMutex.Unlock()
	return throughput
}

func updateLatencyMetrics() {
	for range time.Tick(5 * time.Second) {
		// Simulate latency for demonstration purposes
		latency := rand.Float64() * 0.5 // Adjust for your application's actual latency
		HttpRequestLatency.WithLabelValues("GET").Observe(latency)
	}
}

func main() {
	// Setup Gin router
	router := gin.Default()
	go recordUptime()
	go updateThroughput()
	go updateLatencyMetrics()
	go func() {
		for {
			// Simulate some health check logic here.
			isServiceUp := performHealthCheck()

			// Set the value of the "x_up" metric based on the health check result.
			if isServiceUp {
				xUp.Set(1)
			} else {
				xUp.Set(0)
			}

			// Sleep for a duration before the next health check.
			time.Sleep(1 * time.Minute)
		}
	}()
	router.Use(prometheusMiddleware())
	router.GET("/metrics", prometheusHandler())
	router.Static("/static", "./static/")
	router.GET("/", handlers.HomePage)
	// Authentication and Authorization
	router.GET("/logout", handlers.Logout)
	router.GET("/login", handlers.LoginPage)
	router.GET("/register", handlers.RegistrationPage)
	router.POST("register", handlers.Register)
	router.POST("/login", handlers.Login)
	// ToDo operations
	router.GET("/todo", handlers.ToDoPage)
	router.GET("/add", handlers.AddToDoPage)
	router.GET("/read/:id", handlers.ReadPage)
	router.GET("/edit/:id", handlers.EditPage)
	router.GET("/delete/:id", handlers.DeleteToDo)
	router.POST("/add", handlers.AddToDo)
	router.POST("/edit/:id", handlers.EditToDo)
	// Seacrh, Group, Sort
	router.GET("/search", handlers.ToDoPage)
	router.GET("/sort", handlers.ToDoPage)
	// Admin
	router.GET("/admin/search", handlers.AdminPage)
	router.GET("/admin/sort", handlers.AdminPage)
	router.GET("/admin", handlers.AdminPage)
	router.GET("admin/delete/:user", handlers.DeleteUser)
	router.Run("localhost:3000")
}

func performHealthCheck() bool {
	serviceURL := "http://localhost:3000"

	// Create an HTTP client with a timeout.
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	// Perform a GET request to the service endpoint.
	response, err := client.Get(serviceURL)
	if err != nil {
		// An error occurred, consider the service as down.
		return false
	}
	defer response.Body.Close()

	// Check if the response status is 200.
	return response.StatusCode == http.StatusOK
}

func prometheusMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start).Seconds()

		HttpRequestDuration.WithLabelValues(c.Request.Method).Observe(duration)
		HttpRequestCounter.WithLabelValues(c.Request.Method).Inc()
	}
}
