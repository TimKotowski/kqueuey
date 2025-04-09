package main

import (
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap/zapcore"

	"kqueuey"
)

func main() {
	done := make(chan struct{}, 1)

	flagOpts := kqueuey.FlagOpts{}
	flagOpts.RegisterFlags()
	flagOpts.Parse()

	logger := flagOpts.Logging.NewLogger()
	zapcore.NewTee
	//c, err := kqueuey.LoadConfiguration(flagOpts, logger)
	//if err != nil {
	//	logger.Error(err.Error())
	//	panic(err)
	//}

	logger.Info("config file set up")

	go awaitTerminated(done)
	<-done
}

func awaitTerminated(done chan struct{}) {
	killSignal := make(chan os.Signal, 1)
	signal.Notify(killSignal, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	<-killSignal
	done <- struct{}{}
}
