package main

import (
	"context"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"logstation/config"
	"logstation/routers"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	var err error
	// init config
	err = config.InitConfig("config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()
	store := memstore.NewStore([]byte(uuid.New().String()))
	router.Use(sessions.Sessions("LogStationSession", store))
	routers.InitRouter(router)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", config.Config.Server.Port.HttpPort),
		Handler: router,
	}

	go func() {
		if err = srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err = srv.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
