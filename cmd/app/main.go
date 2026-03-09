package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"
	"time"

	"github.com/patrickkdev/tcptunnelsmanager/configs"
	"github.com/patrickkdev/tcptunnelsmanager/internal/application"
	"github.com/patrickkdev/tcptunnelsmanager/internal/infrastructure/db"
)

func main() {
	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	defer stop()

	dbConn, err := db.Connect(configs.DBConfig)
	if err != nil {
		panic(err)
	} else {
		log.Println("connected to database")
	}
	defer dbConn.Close()

	tunnelRowsRepo := db.NewTunnelRowsRepo(dbConn)
	tunnelLogsRepo := db.NewTunnelLogsRepo(dbConn)

	manager := application.NewManager(tunnelRowsRepo, tunnelLogsRepo)
	manager.Run(ctx, time.Duration(configs.ReconcileIntervalSeconds)*time.Second)
}
