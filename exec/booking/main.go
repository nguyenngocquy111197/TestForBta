package main

import (
	"context"
	"fmt"

	"net/http"
	"os"
	"os/signal"
	"time"

	"example.com/m/v2/configuration"
	"example.com/m/v2/libs/logger"
	"example.com/m/v2/source/apis/bookingService"
	"example.com/m/v2/source/gateway"
	"example.com/m/v2/source/mongodb"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	os.Setenv("CONFIG", "booking.yaml")
	var (
		setting         = configuration.NewService("./").Init()
		entry           = logger.NewModel(setting.LogToFile)
		entryLogGateway = logger.NewModel(setting.LogGatewayToFile)
		dbstore         = mongodb.New(setting.MongoDB.ParseURI(), setting.MongoDB.DBName, setting.MongoDB.Timeout)
		interrupt       = make(chan os.Signal)
		ctx             context.Context
		cancel          context.CancelFunc

		root = gin.New()
		//service gateway

		getway_service gateway.Service = *gateway.New(setting.ProviderGateways, entryLogGateway)

		bookingServices = bookingService.NewServiceAsLogger(entry, dbstore, getway_service)
	)

	fmt.Println(setting, setting.HostServer)
	{
		handlerFunc := cors.New(cors.Config{
			AllowAllOrigins:  true,
			AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodOptions},
			AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "*"},
			AllowCredentials: false,
			AllowWebSockets:  true,
			MaxAge:           12 * time.Hour,
		})

		loggerFunc := gin.LoggerWithWriter(
			entry.Writer(),
		)
		root.Use(handlerFunc, loggerFunc).Use(gin.Recovery())
		root.NoRoute(func(c *gin.Context) { c.AbortWithStatus(http.StatusNotFound) })

	}

	{
		bookingService.NewHandler(bookingServices).AddRoutes(root)

	}

	server := &http.Server{
		Addr:         setting.HostServer.ParseAddr(),
		Handler:      root,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {

		entry.Println("[   Server] Server Version 4 start on ", server.Addr)
		if err := server.ListenAndServe(); err != nil {
			entry.Error(err)
		}
		entry.Println("[   Server] Server has shutdown !")
		interrupt <- os.Interrupt
	}()

	{
		signal.Notify(interrupt, os.Interrupt)
		<-interrupt

		ctx, cancel = context.WithTimeout(context.TODO(), 5*time.Second)
		defer cancel()
		server.Shutdown(ctx)
		os.Exit(0)

	}
}
