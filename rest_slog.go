//go:build go1.21
// +build go1.21

package rest

import (
	"log/slog"
	"os"
)

func init() {
	h := slog.NewTextHandler(os.Stderr, nil)
	Logger = slog.New(h)
}

// Logger logs information about incoming requests.
var Logger *slog.Logger
