package logger

import (
	"go.uber.org/zap/zapcore"
	"io"
	"os"
	"phanes/config"
)

func Init() func() {
	var (
		l       = config.Conf.Collect.Log
		writers = make([]io.Writer, 0, 0)
	)
	if l.FileName == "" && l.RedisKey == "" {
		panic("no log storage target")
	}

	if l.FileName != "" {
		writers = append(writers, FileOutputWriter("./logs", l.FileName, 50, 3))
	}
	if l.RedisKey != "" {
		writers = append(writers, RedisOutputWriter(config.KV, l.RedisKey))
	}
	writers = append(writers, os.Stderr)

	logger := ZapLog(zapcore.DebugLevel, writers...)
	InitLogger(logger)

	return func() {}
}
