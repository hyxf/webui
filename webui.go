package main

import (
	"fmt"
	"github.com/hyxf/webui/server"
	"github.com/urfave/cli"
	"log"
	"os"
	"time"
)

const (
	Host       = "127.0.0.1"
	Port       = "8089"
	AppName    = "webui"
	AppUsage   = "aria2 web ui"
	AppVersion = "1.0.0"
	FlagHost   = "host"
	FlagPort   = "port"
	FlagDaemon = "signal"
)

var flags = []cli.Flag{
	cli.StringFlag{
		Name:  FlagHost,
		Usage: "server host",
		Value: Host,
	},
	cli.StringFlag{
		Name:  FlagPort,
		Usage: "server port",
		Value: Port,
	},
	cli.StringFlag{
		Name:  fmt.Sprintf("%s, s", FlagDaemon),
		Usage: server.FlagUsage,
	},
}

func RunAction(c *cli.Context) error {
	host := c.String(FlagHost)
	port := c.String(FlagPort)
	daemon := c.String(FlagDaemon)
	if daemon != "" {
		ser, err := server.InitDaemon("webui-service", "aria2 webui service")
		if err != nil {
			return err
		}
		result, runErr := ser.Run(daemon, func() error {
			runErr := server.RunServer(host, port)
			if runErr != nil {
				return runErr
			}
			return nil
		})
		if runErr != nil {
			return runErr
		}
		if result != "" {
			log.Println(result)
		}
	} else {
		runErr := server.RunServer(host, port)
		if runErr != nil {
			return runErr
		}
	}
	return nil
}

func execute() error {
	app := cli.NewApp()
	app.Name = AppName
	app.Usage = AppUsage
	app.Version = AppVersion
	app.Compiled = time.Now()
	app.Flags = flags
	app.Action = RunAction
	err := app.Run(os.Args)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	//os.Args = append(os.Args, "-s", "remove")
	err := execute()
	if err != nil {
		log.Fatal(err)
	}
}
