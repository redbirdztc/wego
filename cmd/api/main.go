// @title Wego API
// @version 1.0
// @description API service for Wego application
// @host localhost:8080
// @BasePath /api/v1
package main

import (
	"time"

	"github.com/redbirdztc/wego/internal/conf"
	"github.com/redbirdztc/wego/internal/httpservice"
	"github.com/redbirdztc/wego/pkg/db"
	"github.com/redbirdztc/wego/pkg/postgres"
)

func main() {
	postgres := postgres.NewPostgresDB(conf.GetPostgresDSN())
	db.SetConnectionKeeper(postgres)

	svc := httpservice.New()
	err := svc.Start(":" + conf.GetPort())
	if err != nil {
		panic(err)
	}
	for {
		time.Sleep(time.Hour)
	}
}
