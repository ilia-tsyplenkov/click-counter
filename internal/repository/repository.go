package repository

import (
	"context"
	"time"

	"github.com/ilia-tsyplenkov/click-counter/internal/model"
)

type Repository interface {
	AddClicks(ctx context.Context, clicksMap map[int]int, timestamp time.Time) error
	GetStats(ctx context.Context, id int, from, to time.Time) ([]*model.BannerStat, error)
}
