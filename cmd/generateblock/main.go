package main

import (
	"context"

	arg "github.com/alexflint/go-arg"
	env "github.com/caarlos0/env/v6"
	"github.com/cryptogarageinc/generate-block-for-testing/internal/domain/model"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var argObj *argument
var logger *zap.Logger

func init() {
	argObj = &argument{}
	arg.MustParse(argObj)
	logger = logInit()
}

func logInit() *zap.Logger {
	logOpts := make([]zap.Option, 0, 1)
	logConf := zap.NewProductionConfig()
	logConf.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	logConf.EncoderConfig.TimeKey = "timestamp"
	logConf.EncoderConfig.LevelKey = "status"
	if argObj.Logging {
		logConf.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
		logConf.Encoding = "json"
		logConf.EncoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder
		logOpts = append(logOpts, zap.AddStacktrace(zapcore.WarnLevel), zap.Development())
	} else {
		logConf.Encoding = "console"
		logConf.EncoderConfig.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	}
	logger, err := logConf.Build(logOpts...)
	if err != nil {
		panic(err)
	}
	return logger
}

func main() {
	logger.Debug("start")
	ctx := context.Background()

	envObj := &environment{}
	if err := env.Parse(envObj); err != nil {
		logError("Error while parsing environment", err)
		return
	}

	argObj.setValueFromEnvironment(envObj)
	if err := argObj.Validate(); err != nil {
		logError("Error while validate argument", err)
		return
	}

	var network string
	if argObj.Network != "" {
		network = argObj.Network
	} else {
		network = model.ElementsRegtest.String()
		logger.Debug("set: default network elementsRegTest")
	}

	// dependency
	handle := NewHandler(argObj)

	// execute
	if err := handle.GenerateBlock(
		ctx, network, argObj.FedpegScript, argObj.Pak, argObj.Address); err != nil {
		logError("GenerateBlock fail", err)
	}
	logger.Debug("end")
}

func logError(message string, err error) {
	if argObj.Logging {
		logger.Error(message, zap.Error(err))
	} else {
		// for console
		logger.Sugar().Errorf("%s.\nError: %+v", message, err)
	}
}
