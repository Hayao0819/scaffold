package log

import (
	"log/slog"

	"github.com/m-mizutani/clog"
)

func init() {
	handler := clog.New(clog.WithColor(true))
	logger := slog.New(handler)
	slog.SetDefault(logger)
}
