//go:build !go1.21
// +build !go1.21

package rest

import log "github.com/inconshreveable/log15"

// Logger logs information about incoming requests.
var Logger log.Logger = log.New()
