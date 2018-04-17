package distributor

import (
	"errors"
	"log"
	"os"
	"os/signal"

	"github.com/deckarep/golang-set"
)

var gSignal = mapset.NewSet()

var SigAlreadyRegisted = errors.New("sig already registed")

func RegisterSignal(sig os.Signal, process func()) error {
	sigChan, err := register(sig)
	if err != nil {
		return err
	}
	go func() {
		msg := <-sigChan
		log.Println("MOSS |sig", msg, "received")
		process()
	}()
	return nil
}

func register(sig os.Signal) (chan os.Signal, error) {
	if gSignal.Contains(sig) {
		log.Fatal("MOSS |signal", sig, "already registed")
		return nil, SigAlreadyRegisted
	}
	gSignal.Add(sig)
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, sig)
	return sigChan, nil
}

func RegisterContinueSignal(sig os.Signal, process func()) error {
	sigChan, err := register(sig)
	if err != nil {
		return err
	}
	go func() {
		for {
			msg := <-sigChan
			log.Println("MOSS |sig", msg, "receveid")
			process()
		}
	}()
	return nil
}

func SignalProcessed(sig os.Signal) bool {
	return gSignal.Contains(sig)
}
