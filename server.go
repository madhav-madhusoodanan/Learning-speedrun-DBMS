package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"rampx/backend/apps/oracle"
	"rampx/backend/apps/ping"
	"rampx/backend/apps/swaps"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
)

// Create context that listens for the interrupt signal from the OS.
func main(){
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	router := gin.Default()

	secureMiddleware := secure.New(secure.Options{
        // AllowedHosts:          []string{"localhost"},
        // AllowedHostsAreRegex:  true,
        HostsProxyHeaders:     []string{"X-Forwarded-Host"},
        SSLRedirect:           false,
        // SSLHost:               "ssl.example.com",
        // SSLProxyHeaders:       map[string]string{"X-Forwarded-Proto": "https"},
        STSSeconds:            31536000,
        STSIncludeSubdomains:  true,
        STSPreload:            true,
        FrameDeny:             true,
        ContentTypeNosniff:    true,
        BrowserXssFilter:      true,
        // ContentSecurityPolicy: "script-src $NONCE",
    })

    _ = func() gin.HandlerFunc {
        return func(c *gin.Context) {
            err := secureMiddleware.Process(c.Writer, c.Request)

            // If there was an error, do not continue.
            if err != nil {
                c.Abort()
                return
            }

            // Avoid header rewrite if response is a redirection.
            if status := c.Writer.Status(); status > 300 && status < 399 {
                c.Abort()
            }
        }
    }()

	// router.Use(secureMiddlewareFunc)
	router.Use(cors.Default())

	swaps.Setup(router.Group("/swap"))
	ping.Setup(router.Group("/ping"))
	oracle.Setup(router.Group("/oracle"))

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Listen for the interrupt signal.
	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	log.Println("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}