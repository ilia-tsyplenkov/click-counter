package model

import "time"

type BannerStat struct {
	ID        int32
	Clicks    int64
	Timestamp time.Time
}
