package service

import (
	"context"
	"time"

	"github.com/ilia-tsyplenkov/click-counter/internal/model"
	"github.com/ilia-tsyplenkov/click-counter/internal/repository"
)

type Service interface {
	Inc(id int)
	Stat(ctx context.Context, id int, from, to time.Time) ([]*model.BannerStat, error)
}

func New(repo repository.Repository) Service {
	srv := &clickStatService{
		repo:  repo,
		query: make(chan int, 100_000),
	}
	go srv.collect()
	return srv
}

type clickStatService struct {
	repo  repository.Repository
	query chan int
}

func (srv *clickStatService) Inc(id int) {
	srv.query <- id
}

func (srv *clickStatService) Stat(ctx context.Context, id int, from, to time.Time) ([]*model.BannerStat, error) {
	return srv.repo.GetStats(ctx, id, from.UTC(), to.UTC())
}
