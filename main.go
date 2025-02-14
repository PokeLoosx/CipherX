package main

import (
	"CipherX/config"
	"CipherX/initialize"
	"CipherX/internal"
	"CipherX/utils"
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
)

func main() {
	// Defining Command Line Parameters
	configPath := flag.String("config", "config.yaml", "Configuration file path")
	mode := flag.String("mode", gin.ReleaseMode, "Application running mode")
	host := flag.String("host", "0.0.0.0", "Application running host")
	port := flag.Int("port", 8080, "Application running port")

	// Parsing command line arguments
	flag.Parse()

	// Initialize Mode
	initializeMode(*mode)

	// Initialize Logger
	config.GinLOG = initialize.Zap()

	// Check if the configuration file exists
	if utils.PathFileExists(*configPath) {
		// Initialize Configuration
		config.GinVP = initialize.Viper(*configPath)

		// Initialize Database
		config.GinDB = initialize.DB()
		if config.GinDB == nil {
			fmt.Println("Failed to initialize database...")
			return
		}

		// Initialize Redis
		config.GinRedis = initialize.Redis()
		if config.GinRedis == nil {
			fmt.Println("Failed to initialize redis...")
			return
		}
	} else {
		fmt.Println("Configuration file not found, installation required")
	}

	// Initializing routes
	router := internal.Routers()
	if router == nil {
		fmt.Println("Failed to initialize route...")
		return
	}

	// Creating a service
	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", *host, *port),
		Handler: router,
	}

	// Startup
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Listten: %s\n", err)
		}
	}()

	// Wait for the terminal signal to gracefully shut down the server, setting a 10-second timeout for shutting down the server
	quit := make(chan os.Signal, 1) // Create a channel to receive the signal

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // It doesn't block here
	<-quit                                               // Block here and only continue execution when both signals are received
	config.GinLOG.Info("Service ready to shut down")

	// Create a Context with a 10-second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// Gracefully shut down the service within 10 seconds (stop the service after processing pending requests), and time out after 10 seconds
	if err := srv.Shutdown(ctx); err != nil {
		config.GinLOG.Fatal("Service timed out has been shut down: ", zap.Error(err))
	}

	config.GinLOG.Info("Service has been shut down")
}

// Initialize Mode
func initializeMode(mode string) {
	switch mode {
	case gin.DebugMode:
		gin.SetMode(gin.DebugMode)
	case gin.TestMode:
		gin.SetMode(gin.TestMode)
	case gin.ReleaseMode:
		gin.SetMode(gin.ReleaseMode)
	default:
		gin.SetMode(gin.ReleaseMode)
	}
}
