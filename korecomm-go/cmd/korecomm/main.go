package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/hegemone/kore-poc/korecomm-go/pkg/comm"
	"github.com/hegemone/kore-poc/korecomm-go/pkg/mock"
	log "github.com/sirupsen/logrus"
)

const FatalExitCode = 1

func main() {
	var err error

	// TODO: Should be configurable
	//log.SetLevel(log.DebugLevel)
	log.SetLevel(log.InfoLevel)

	log.Info("============================================================")
	log.Info("                 Kore::Comm Golang POC")
	log.Info("============================================================")

	// Some amount of startup validation like confirming the location of config file and
	// configured extension directores..
	if err = validateStartup(); err != nil {
		log.Fatalf("Startup validation failed with the following error, exiting:")
		log.Fatal(err.Error())
		os.Exit(FatalExitCode)
	}

	// Start the demux early
	demux := mock.StdinDemuxInstance()
	demux.Listen()

	engine := comm.NewEngine()

	if err = engine.LoadExtensions(); err != nil {
		log.Fatalf("Fatal error occurred loading extensions during startup: ", err.Error())
		log.Fatal(err.Error())
	}

	go func() {
		log.Info("Starting engine")
		engine.Start()
	}()

	// Listen for shutdown signals
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs,
		syscall.SIGINT,
		syscall.SIGQUIT,
		syscall.SIGTERM,
		syscall.SIGKILL,
	)

	go func() {
		sig := <-sigs
		log.Infof("Received shutdown signal: %v", sig)
		// Possibly do some kind of shutdown cleanup
		// engine.Shutdown()
		done <- true
	}()

	<-done // Keep alive until done
	log.Info("Kore::Comm finished.")
}

func validateStartup() error {
	log.Info("Performing startup validations")
	return nil // TODO: Faking out
}
