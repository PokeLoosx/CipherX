package initialize

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"path"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"CipherX/utils"
)

const director = "logs"

var FileRotateLogs = new(fileRotateLogs)

type fileRotateLogs struct{}

// GetWriteSyncer Get zapcore.WriteSyncer
func (r *fileRotateLogs) GetWriteSyncer(level string) (zapcore.WriteSyncer, error) {
	fileWriter, err := rotatelogs.New(
		path.Join(director, "%Y-%m-%d", level+".log"),
		rotatelogs.WithClock(rotatelogs.Local),
		rotatelogs.WithMaxAge(30*24*time.Hour), // Log retention time
		rotatelogs.WithRotationTime(time.Hour*24),
	)

	return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter)), err
}

type _zap struct{}

// GetEncoder Get zapcore.Encoder
func (z *_zap) GetEncoder() zapcore.Encoder {
	return zapcore.NewJSONEncoder(z.GetEncoderConfig())
}

// GetEncoderConfig Get zapcore.EncoderConfig
func (z *_zap) GetEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     z.CustomTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

// GetEncoderCore Get Encoder's zapcore.Core
func (z *_zap) GetEncoderCore(l zapcore.Level, level zap.LevelEnablerFunc) zapcore.Core {
	writer, err := FileRotateLogs.GetWriteSyncer(l.String()) // Use file-rotatelogs for log splitting
	if err != nil {
		fmt.Printf("Get Write Syncer Failed err:%v", err.Error())
		return nil
	}

	return zapcore.NewCore(z.GetEncoder(), writer, level)
}

// CustomTimeEncoder Custom log output time format
func (z *_zap) CustomTimeEncoder(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
	encoder.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

// TransportLevel converts string to zapcore.Level
func (z *_zap) TransportLevel() zapcore.Level {
	switch gin.Mode() {
	case gin.DebugMode:
		return zapcore.DebugLevel
	case gin.ReleaseMode:
		return zapcore.InfoLevel
	default:
		return zapcore.DebugLevel
	}
}

// GetZapCores Get []zapcore.Core based on the Level in the configuration file
func (z *_zap) GetZapCores() []zapcore.Core {
	cores := make([]zapcore.Core, 0, 7)
	for level := z.TransportLevel(); level <= zapcore.FatalLevel; level++ {
		cores = append(cores, z.GetEncoderCore(level, z.GetLevelPriority(level)))
	}
	return cores
}

// GetLevelPriority Get zap.LevelEnablerFunc based on zapcore.Level
func (z *_zap) GetLevelPriority(level zapcore.Level) zap.LevelEnablerFunc {
	switch level {
	case zapcore.DebugLevel:
		return func(level zapcore.Level) bool { // Debug level
			return level == zap.DebugLevel
		}
	case zapcore.InfoLevel:
		return func(level zapcore.Level) bool { // Info level
			return level == zap.InfoLevel
		}
	case zapcore.WarnLevel:
		return func(level zapcore.Level) bool { // Warning level
			return level == zap.WarnLevel
		}
	case zapcore.ErrorLevel:
		return func(level zapcore.Level) bool { // Error level
			return level == zap.ErrorLevel
		}
	case zapcore.DPanicLevel:
		return func(level zapcore.Level) bool { // DPanic level
			return level == zap.DPanicLevel
		}
	case zapcore.PanicLevel:
		return func(level zapcore.Level) bool { // Panic level
			return level == zap.PanicLevel
		}
	case zapcore.FatalLevel:
		return func(level zapcore.Level) bool { // Fatal level
			return level == zap.FatalLevel
		}
	default:
		return func(level zapcore.Level) bool { // Debug level
			return level == zap.DebugLevel
		}
	}
}

// Zap Initialize logger
func Zap() (logger *zap.Logger) {
	if ok, _ := utils.PathExists(director); !ok { // Check if Director folder exists
		fmt.Printf("create %v directory\n", director)
		_ = os.Mkdir(director, os.ModePerm)
	}
	var z = new(_zap)
	cores := z.GetZapCores()
	logger = zap.New(zapcore.NewTee(cores...))
	logger = logger.WithOptions(zap.AddCaller())

	fmt.Println("Zap initialized successfully")
	return logger
}
