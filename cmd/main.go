package main

import (
	"context"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	log "github.com/sirupsen/logrus"

	"github.com/ilia-tsyplenkov/click-counter/config"
	"github.com/ilia-tsyplenkov/click-counter/internal/handler"
	pgRepo "github.com/ilia-tsyplenkov/click-counter/internal/repository/postgres"
	"github.com/ilia-tsyplenkov/click-counter/internal/service"
	"github.com/labstack/echo/v4"
)

func main() {

	cfg, err := config.New()
	if err != nil {
		log.Fatalf("config error: %v", err)
	}

	pgxCfg, err := pgxpool.ParseConfig(cfg.ConnectionString)
	if err != nil {
		log.Fatalf("parse connection string: %v", err)
	}
	pgxpool, err := pgxpool.NewWithConfig(context.Background(), pgxCfg)
	if err != nil {
		log.Fatalf("failed connect to db: %v", err)
	}
	repo := pgRepo.New(pgxpool)
	srv := service.New(repo)
	bannerHandlers := handler.New(srv)
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/counter/:banner_id", bannerHandlers.GetBannerClick)
	e.POST("/stat/:banner_id", bannerHandlers.GetBannerStats)
	e.Logger.Fatal(e.Start(cfg.ServerAddr))
}
