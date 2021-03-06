package readers

import (
	"errors"

	"github.com/mainflux/mainflux"
)

// ErrNotFound indicates that requested entity doesn't exist.
var ErrNotFound = errors.New("entity not found")

// MessageRepository specifies message reader API.
type MessageRepository interface {
	// ReadAll skips given number of messages for given channel and returns next
	// limited number of messages.
	ReadAll(uint64, uint64, uint64) []mainflux.Message
}
