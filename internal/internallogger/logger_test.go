package internallogger

import (
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"testing"
)

func TestNewLogger(t *testing.T) {
	logger := NewLogger()

	// Check if internallogger is not nil
	assert.NotNil(t, logger)

	// Check if internallogger is of type *slog.Logger
	assert.IsType(t, zap.SugaredLogger{}, *logger)
}
