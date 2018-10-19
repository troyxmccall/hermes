package core

import (
	"github.com/spf13/afero"
	"os"
	"os/signal"
	"syscall"
)

var (
	appFs afero.Fs
)

func Setup() {
	appFs = afero.NewOsFs()
	sigChannel := make(chan os.Signal, 2)
	signal.Notify(sigChannel, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		for sig := range sigChannel {
			if sig == syscall.SIGINT {
				ExitWithErrorMessage(red("Keyboard Interrupt"))
			} else if sig == syscall.SIGTERM {
				Exit(0)
			} else {
				ExitWithErrorMessage("Unknown Signal: " + sig.String())
			}
		}
	}()
}
