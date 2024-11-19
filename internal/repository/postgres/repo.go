package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/ilia-tsyplenkov/click-counter/internal/model"
	"github.com/ilia-tsyplenkov/click-counter/internal/repository"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	log "github.com/sirupsen/logrus"
)

const (
	insertClicksQuery = `INSERT INTO core.banner_clicks (id, clicks, created_at) VALUES ($1, $2, $3)`
	selectStatsQuery  = `SELECT id, SUM(clicks), created_at FROM core.banner_clicks WHERE id = $1 AND created_at BETWEEN $2 AND $3 GROUP BY id, created_at`
)

var _ repository.Repository = &repo{}

type repo struct {
	conn *pgxpool.Pool
}

func New(conn *pgxpool.Pool) repository.Repository {
	return &repo{
		conn: conn,
	}
}

func (r *repo) AddClicks(ctx context.Context, clicksMap map[int]int, timestamp time.Time) error {
	batch := &pgx.Batch{}
	for k, v := range clicksMap {
		batch.Queue(insertClicksQuery, k, v, timestamp)
	}
	br := r.conn.SendBatch(ctx, batch)
	defer br.Close()
	_, err := br.Exec()
	return err
}

func (r *repo) GetStats(ctx context.Context, id int, from, to time.Time) ([]*model.BannerStat, error) {
	var stats []*model.BannerStat
	log.Infof("get stat: id: %d from: %s to: %s", id, from, to)
	rows, err := r.conn.Query(ctx, selectStatsQuery, id, from, to)
	if err != nil {
		return nil, fmt.Errorf("query: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		bs := model.BannerStat{}
		if err := rows.Scan(&bs.ID, &bs.Clicks, &bs.Timestamp); err != nil {
			return nil, fmt.Errorf("scan: %v", err)
		}
		stats = append(stats, &bs)
	}

	return stats, nil
}
