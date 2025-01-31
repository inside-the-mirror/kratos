package blademaster

import (
	"fmt"
	"sync/atomic"
	"time"

	"github.com/inside-the-mirror/kratos/pkg/log"
	limit "github.com/inside-the-mirror/kratos/pkg/ratelimit"
	"github.com/inside-the-mirror/kratos/pkg/ratelimit/bbr"
	"github.com/inside-the-mirror/kratos/pkg/stat/prom"
)

const (
	_statName = "go_http_bbr"
)

var (
	bbrStats = prom.New().WithState("go_http_bbr", []string{"url"})
)

// RateLimiter bbr middleware.
type RateLimiter struct {
	group   *bbr.Group
	logTime int64
}

// New return a ratelimit middleware.
func NewRateLimiter(conf *bbr.Config) (s *RateLimiter) {
	return &RateLimiter{
		group:   bbr.NewGroup(conf),
		logTime: time.Now().UnixNano(),
	}
}

func (b *RateLimiter) printStats(routePath string, limiter limit.Limiter) {
	now := time.Now().UnixNano()
	if now-atomic.LoadInt64(&b.logTime) > int64(time.Second*3) {
		atomic.StoreInt64(&b.logTime, now)
		log.Info("http.bbr path:%s stat:%+v", routePath, limiter.(*bbr.BBR).Stat())
	}
}

// Limit return a bm handler func.
func (b *RateLimiter) Limit() HandlerFunc {
	return func(c *Context) {
		uri := fmt.Sprintf("%s://%s%s", c.Request.URL.Scheme, c.Request.Host, c.Request.URL.Path)
		limiter := b.group.Get(uri)
		done, err := limiter.Allow(c)
		if err != nil {
			bbrStats.Incr(_statName, uri)
			c.JSON(nil, err)
			c.Abort()
			return
		}
		defer func() {
			done(limit.DoneInfo{Op: limit.Success})
			b.printStats(uri, limiter)
		}()
		c.Next()
	}
}
