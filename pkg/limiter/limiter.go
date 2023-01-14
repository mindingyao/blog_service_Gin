package limiter

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

type LimiterIface interface {
	Key(c *gin.Context) string
	GetBucket(key string) (*ratelimit.Bucket, bool)
	AddBuckets(rules ...LimiterBuckerRule) LimiterIface
}

type Limiter struct {
	limiterBuckets map[string]*ratelimit.Bucket
}

type LimiterBuckerRule struct {
	Key          string
	FillInterval time.Duration
	Capacity     int64
	Quantum      int64
}
