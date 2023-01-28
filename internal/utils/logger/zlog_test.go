package logger

import (
	"github.com/rs/zerolog/log"
	"testing"
)

func TestZLog(t *testing.T) {

	log.Debug().Caller().
		Str("Scale", "833 cents").
		Float64("Interval", 833.09).
		Msg("Fibonacci is everywhere")

	log.Debug().
		Str("Name", "Tom").
		Send()
}
