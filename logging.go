package main

import (
	"time"

	"github.com/go-kit/kit/log"
)

type loggingMiddleware struct {
	logger log.Logger
	next   TaxonomyService
}

func (mw loggingMiddleware) Taxonomy() ([]Record, error) {
	begin := time.Now()
	output, err := mw.next.Taxonomy()
	_ = mw.logger.Log("Output %v", output)

	_ = mw.logger.Log(
		"method", "taxonomys",
		"input", "",
		"err", err,
		"took", time.Since(begin),
	)

	return output, err
}
