package server

import (
	"github.com/takama/daemon"
	"log"
	"os"
	"os/signal"
	"syscall"
)

const FlagUsage = "signal: install | remove | start | stop | status"

type Service struct {
	srv daemon.Daemon
}

func (d *Service) Run(command string, entry func() error) (string, error) {
	switch command {
	case "install":
		return d.srv.Install()
	case "start":
		return d.srv.Start()
	case "remove":
		return d.srv.Remove()
	case "stop":
		return d.srv.Stop()
	case "status":
		return d.srv.Status()
	default:
		return FlagUsage, nil
	}
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, os.Kill, syscall.SIGTERM)
	//----
	err := entry()
	if err != nil {
		return FlagUsage, err
	}
	//----
	for {
		select {
		case killSignal := <-interrupt:
			log.Println("Got signal:", killSignal)
			if killSignal == os.Interrupt {
				return "Daemon was interruped by system signal", nil
			}
			return "Daemon was killed", nil
		}
	}
	return FlagUsage, nil
}

func InitDaemon(serviceName, description string, dependencies ...string) (*Service, error) {
	srv, err := daemon.New(serviceName, description, dependencies...)
	if err != nil {
		return nil, err
	}
	d := &Service{
		srv: srv,
	}
	return d, nil
}
