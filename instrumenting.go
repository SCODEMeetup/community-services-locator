package main

import (
	"fmt"
	"time"

	"github.com/go-kit/kit/metrics"
)

type instrumentingMiddleware struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	next           TaxonomyService
}

func (mw instrumentingMiddleware) Taxonomy() ([]Record, error) {
	begin := time.Now()
	taxonomy, err := mw.next.Taxonomy()

	lvs := []string{"method", "taxonomys", "error", fmt.Sprint(err != nil)}
	mw.requestCount.With(lvs...).Add(1)
	mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	return taxonomy, err
}
