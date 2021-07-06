package main

import (
	"fmt"
	"net/http"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/weaveworks/metal-janitor-action/action"
)

func main() {
	for _, env := range os.Environ() {
		fmt.Printf("%s\n", env)
	}

	logger, err := configureLogger()
	if err != nil {
		fmt.Printf("failed to configure logger: %s\n", err.Error())
		os.Exit(1)
	}
	defer logger.Sync() //nolint: errcheck

	input, err := action.NewInput()
	if err != nil {
		logger.Sugar().Fatalw("failed parsing action input", "error", err.Error())
	}
	action, err := action.New(input.APIKey, logger, http.DefaultClient)
	if err != nil {
		logger.Sugar().Fatalw("failed to created action", "error", err.Error())
	}

	if err := action.Cleanup(input.Projects, input.DryRun); err != nil {
		logger.Sugar().Fatalw("failed to cleanup projects", "error", err.Error())
	}
}

func configureLogger() (*zap.Logger, error) {
	conf := zap.NewProductionConfig()

	conf.Encoding = "console"
	conf.EncoderConfig.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	conf.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	conf.Level.SetLevel(zap.InfoLevel)

	return conf.Build()
}
