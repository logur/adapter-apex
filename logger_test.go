package apex

import (
	"testing"

	"logur.dev/logur"
	"logur.dev/logur/conformance"
)

// nolint: gochecknoglobals
var levelMap = map[logur.Level]string{
	logur.Trace: "trace",
	logur.Debug: "debug",
	logur.Info:  "info",
	logur.Warn:  "warn",
	logur.Error: "error",
}

func TestLogger(t *testing.T) {
	t.Skip("implement me")

	suite := conformance.TestSuite{
		LoggerFactory: func(level logur.Level) (logur.Logger, conformance.TestLogger) {
			_ = levelMap

			return New(), conformance.TestLoggerFunc(func() []logur.LogEvent {
				return []logur.LogEvent{}
			})
		},
	}

	suite.Run(t)
}
