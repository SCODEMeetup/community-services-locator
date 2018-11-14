package main

import (
	"time"

	"github.com/go-kit/kit/log"
)

type loggingMiddleware struct {
	logger log.Logger
	next   FoodPantryService
}

func (mw loggingMiddleware) Providers() ([]Provider, error) {
	begin := time.Now()
	output, err := mw.next.Providers()

	_ = mw.logger.Log(
		"method", "providers",
		"input", "",
		"err", err,
		"took", time.Since(begin),
	)

	return output, err
}
