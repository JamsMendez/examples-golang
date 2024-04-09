package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

const API_KEY = "YOUR_API_KEY"

func main() {
	apiGatewayURL := "http://localhost:8080"
	microserviceOneURL := "http://localhost:8081"
	microserviceTwoURL := "http://localhost:8082"

	proxyOne := createReverseProxy(microserviceOneURL)
	proxyTwo := createReverseProxy(microserviceTwoURL)

	r := gin.Default()

	r.Use(authenticate())
	r.Use(rateLimit())

	r.Any("/service_one/*path", proxyOne)
	r.Any("/service_two/*path", proxyTwo)

	// run how other microservices
	go runMicroserviceOne()
	go runMicroserviceTwo()

	log.Printf("Starting server API Gateway on %s", apiGatewayURL)
	log.Fatal(r.Run(":8080"))
}

func createReverseProxy(targetURL string) func(*gin.Context) {
	target, err := url.Parse(targetURL)
	if err != nil {
		log.Fatalf("Error parsing url: %v", err)
	}

	proxy := httputil.NewSingleHostReverseProxy(target)

	return func(c *gin.Context) {
		log.Printf("Proxying request: %s", c.Request.URL)

		c.Request.URL.Path = strings.TrimPrefix(c.Request.URL.Path, "/service_one")
		c.Request.URL.Path = strings.TrimPrefix(c.Request.URL.Path, "/service_two")

		proxy.ServeHTTP(c.Writer, c.Request)
	}
}

func authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Header.Get("X-API-KEY") != API_KEY {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{"error": "unauthorized"},
			)

			return
		}

		c.Next()
	}
}

func rateLimit() gin.HandlerFunc {
	limiter := rate.NewLimiter(rate.Every(1*time.Second), 5)

	return func(c *gin.Context) {
		if !limiter.Allow() {
			c.AbortWithStatusJSON(
				http.StatusTooManyRequests,
				gin.H{"error": "too many requests"},
			)

			return
		}

		c.Next()
	}
}
