package main

import (
	"fmt"
	"time"

	"github.com/go-kit/kit/metrics"
)

type instrumentingMiddleware struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	next           FoodPantryService
}

func (mw instrumentingMiddleware) Providers() ([]Provider, error) {
	begin := time.Now()
	providers, err := mw.next.Providers()

	lvs := []string{"method", "providers", "error", fmt.Sprint(err != nil)}
	mw.requestCount.With(lvs...).Add(1)
	mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	return providers, err
}
