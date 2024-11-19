package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/ilia-tsyplenkov/click-counter/internal/handler/request"
	"github.com/ilia-tsyplenkov/click-counter/internal/service"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type Counter struct {
	service service.Service
}

func New(service service.Service) *Counter {
	return &Counter{
		service: service,
	}
}

// GetBannerClick
// @Title 		GetBannerClick
// @Description GET click on banner
// @Tags        counter
// @Param 		banner_id path integer true "banner id"
// @Success 	200
// @Failure 	400
// @Router 		/counter/{banner_id} [get]
func (ctr *Counter) GetBannerClick(c echo.Context) error {
	l := log.WithField("action", "handler.GetBannerClick")
	bannerID := c.Param("banner_id")
	id, err := strconv.Atoi(bannerID)
	if err != nil {
		l.Errorf("parsing banner id: %v", err)
		return fmt.Errorf("error parsing banner id: %+v", err)
	}
	ctr.service.Inc(id)
	return c.NoContent(http.StatusOK)
}

// BannerStats
// @Title 		GetBannerStats
// @Description POST clicks stats for a banner
// @Tags        stats
// @Param 		banner_id path integer true "banner id"
// @Param       tsFrom query int true "unix by UTC"
// @Param       tsTo query int true "unix by UTC"
// @Success 	200 {object}
// @Failure 	400
// @Router 		/stats/{banner_id} [post]
func (ctr *Counter) GetBannerStats(c echo.Context) error {
	l := log.WithField("action", "handler.GetBannerStats")
	bannerID := c.Param("banner_id")
	id, err := strconv.Atoi(bannerID)
	if err != nil {
		l.Errorf("error parsing banner id: %v", err)
		return fmt.Errorf("error parsing banner id: %+v", err)
	}
	var req request.GetStatRequest
	l.Infof("query: %v", c.Request().URL.Query())
	// c.Bind() doesn't work for post
	// here is a workaround
	if err = (&echo.DefaultBinder{}).BindQueryParams(c, &req); err != nil {
		l.Errorf("invalid request: %v", err)
		return fmt.Errorf("invalid request: %v", err)
	}
	l.Infof("from: %d to: %d", req.From, req.To)
	stat, err := ctr.service.Stat(c.Request().Context(), id, time.Unix(req.From, 0), time.Unix(req.To, 0))
	if err != nil {
		l.Errorf("failed to get stat: %v", err)
		return fmt.Errorf("failed to get stat: %+v", err)
	}
	return c.JSON(http.StatusOK, stat)
}
