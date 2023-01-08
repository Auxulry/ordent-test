// Package constants is describe all constants used in this source
package constants

import (
	"time"
)

const (
	Green        = "\u001B[32m"
	Orange       = "\u001B[0;33m"
	Blue         = "\033[34m"
	Red          = "\033[34m"
	DefaultColor = "\033[0m"
)

const (
	ReadTimeout    = 10 * time.Second
	WriteTimeout   = 10 * time.Second
	MaxHeaderBytes = 1 << 20
)

const Salt = 4

const DefaultErrCode = 599
