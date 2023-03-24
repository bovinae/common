package logger

import (
	"testing"
	"time"

	"go.uber.org/zap"
)

func TestLogger(t *testing.T) {
	InitLogger()
	Info("test", zap.Reflect("now", time.Now()))
}
