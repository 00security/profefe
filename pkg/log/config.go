package log

import (
	"flag"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	Level  string
	Format string
}

func (conf *Config) RegisterFlags(f *flag.FlagSet) {
	f.StringVar(&conf.Level, "log.level", "error", "log level")
	f.StringVar(&conf.Format, "log.format", "console", "(TODO) log formatter")
}

func (conf *Config) Build() (*Logger, error) {
	encConf := zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	configLevel, err := zap.ParseAtomicLevel(conf.Level)
	if err != nil {
		return nil, err
	}
	zapConf := zap.Config{
		Level:             configLevel,
		Development:       true,
		DisableStacktrace: true,
		Encoding:          conf.Format,
		EncoderConfig:     encConf,
		OutputPaths:       []string{"stderr"},
		ErrorOutputPaths:  []string{"stderr"},
	}

	// add caller skip 1 to prevent zap from reporting logger as the source of the log line (caller)
	baseLogger, err := zapConf.Build(zap.AddCallerSkip(1))
	if err != nil {
		return nil, err
	}

	zap.RedirectStdLog(baseLogger)

	return New(baseLogger), nil
}
