package service

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"
)

func (srv *clickStatService) collect() {
	l := log.WithField("action", "service.collect")
	var m = make(map[int]int)
	for ticker := time.Tick(time.Second); ; {
		select {
		case id := <-srv.query:
			m[id]++
		case <-ticker:
			if len(m) > 0 {
				ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
				if err := srv.repo.AddClicks(ctx, m, time.Now().Round(60*time.Second).UTC()); err != nil {
					l.Errorf("add clicks: %v", err)
				} else {
					m = make(map[int]int)
				}
			}
		}
	}
}
