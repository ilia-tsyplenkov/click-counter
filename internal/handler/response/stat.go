package response

import "github.com/ilia-tsyplenkov/click-counter/internal/model"

type BannerStatistics struct {
	Stat []*model.BannerStat `json:"banner_stat"`
}
