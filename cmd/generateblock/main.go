package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	arg "github.com/alexflint/go-arg"
	env "github.com/caarlos0/env/v6"
	"github.com/cryptogarageinc/generate-block-for-testing/internal/domain/model"
	"github.com/cryptogarageinc/generate-block-for-testing/internal/handler"
	pkgerror "github.com/cryptogarageinc/generate-block-for-testing/internal/pkg/errors"
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
	if argObj.PollingTime == time.Duration(0) {
		if err := handle.GenerateBlock(
			ctx,
			network,
			argObj.FedpegScript,
			argObj.Pak,
			argObj.Address,
			argObj.GenerateCount,
			argObj.IgnoreEmptyMempool,
		); err != nil {
			if err == pkgerror.ErrEmptyMempoolTx {
				logger.Debug("empty mempool tx. skip generate block.")
			} else {
				logError("GenerateBlock fail", err)
			}
		}
	} else {
		// process mode
		if err := run(handle, network, argObj.PollingTime); err != nil {
			logError("GenerateBlock fail", err)
		}
	}
	logger.Debug("end", zap.Uint("generateCount", argObj.GenerateCount))
}

func run(handle handler.Handler, network string, pollingTime time.Duration) error {
	ctx, stop := signal.NotifyContext(
		context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT)
	defer stop()

	ticker := time.NewTicker(pollingTime)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			logger.Debug("call GenerateBlock")
			if err := handle.GenerateBlock(
				ctx,
				network,
				argObj.FedpegScript,
				argObj.Pak,
				argObj.Address,
				argObj.GenerateCount,
				argObj.IgnoreEmptyMempool,
			); err != nil {
				if err == pkgerror.ErrEmptyMempoolTx {
					logger.Debug("empty mempool tx. skip generate block.")
				} else {
					logError("GenerateBlock fail", err)
				}
			}
		case <-ctx.Done():
			return nil
		}
	}
}

func logError(message string, err error) {
	if argObj.Logging {
		logger.Error(message, zap.Error(err))
	} else {
		// for console
		logger.Sugar().Errorf("%s.\nError: %+v", message, err)
	}
}
