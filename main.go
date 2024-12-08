package main

import (
	"net/http"
	"os"
	"time"

	"github.com/Biskwit/CoTify/controllers"
	"github.com/Biskwit/CoTify/middlewares"

	ratelimit "github.com/JGLTechnologies/gin-rate-limit"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func main() {
	godotenv.Load()

	logLevel, _ := log.ParseLevel(os.Getenv("LOG_LEVEL"))
	log.SetLevel(logLevel)

	r := gin.New()
	r.Use(gin.Recovery())

	// Rate limit middleware
	store := ratelimit.InMemoryStore(&ratelimit.InMemoryOptions{
		Rate:  time.Second,
		Limit: 20,
	})
	ratelimit := ratelimit.RateLimiter(store, &ratelimit.Options{
		ErrorHandler: func(c *gin.Context, info ratelimit.Info) {

			c.AbortWithStatus(http.StatusTooManyRequests)
		},
		KeyFunc: func(c *gin.Context) string {
			return c.ClientIP()
		},
	})
	r.Use(ratelimit)

	// CORS middleware
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true                                                               // Allow all origins
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}          // You can also use cors.Default() to allow all standard methods
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"} // Or use cors.Default() to allow all standard headers
	config.AllowCredentials = true                                                              // Allow credentials
	r.Use(cors.New(config))

	//Auth middleware
	AuthRoutes := r.Group("/")
	AuthRoutes.Use(middlewares.AuthMiddleware())

	//Forwarding
	AuthRoutes.Any("*path", controllers.Forward)

	//Server
	log.Info("CoT Proxy is running on port " + os.Getenv("PORT"))
	r.Run(os.Getenv("HOST") + ":" + os.Getenv("PORT")) // listen and serve on
}
