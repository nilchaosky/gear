package logz

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type zapCore struct {
	level zapcore.Level
	zapcore.Core
}

func newZapCore(level zapcore.Level) *zapCore {
	entity := &zapCore{level: level}
	syncer := entity.writeSyncer()
	levelEnabler := zap.LevelEnablerFunc(func(l zapcore.Level) bool {
		return l == level
	})
	entity.Core = zapcore.NewCore(Config.encoder(), syncer, levelEnabler)
	return entity
}

func (z *zapCore) writeSyncer(formats ...string) zapcore.WriteSyncer {
	cutter := NewCutter(
		Config.Director,
		z.level.String(),
		Config.RetentionDay,
		CutterWithLayout(time.DateOnly),
		CutterWithFormats(formats...),
	)
	if Config.LogInConsole {
		multiSyncer := zapcore.NewMultiWriteSyncer(os.Stdout, cutter)
		return zapcore.AddSync(multiSyncer)
	}
	return zapcore.AddSync(cutter)
}

func (z *zapCore) enabled(level zapcore.Level) bool {
	return z.level == level
}

func (z *zapCore) with(fields []zapcore.Field) zapcore.Core {
	return z.Core.With(fields)
}

func (z *zapCore) check(entry zapcore.Entry, check *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	if z.Enabled(entry.Level) {
		return check.AddCore(entry, z)
	}
	return check
}

func (z *zapCore) write(entry zapcore.Entry, fields []zapcore.Field) error {
	for i := 0; i < len(fields); i++ {
		if fields[i].Key == "business" || fields[i].Key == "folder" || fields[i].Key == "directory" {
			syncer := z.writeSyncer(fields[i].String)
			z.Core = zapcore.NewCore(Config.encoder(), syncer, z.level)
		}
	}
	return z.Core.Write(entry, fields)
}

func (z *zapCore) sync() error {
	return z.Core.Sync()
}
