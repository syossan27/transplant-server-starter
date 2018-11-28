package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "transplant-server-starter"
	app.Usage = ""
	app.Flags = []cli.Flag {
		cli.StringFlag{
			Name: "dir",
			Usage: "working directory, start_server do chdir to before exec (optional)",
		},
		cli.Int64Flag{
			Name: "interval",
			Value: 1,
			Usage: "minimum interval (in seconds) to respawn the server program",
		},
		cli.StringSliceFlag{
			Name: "port",
			Usage: "TCP port to listen to (if omitted, will not bind to any ports)",
		},
		cli.StringSliceFlag{
			Name: "path",
			Usage: "path at where to listen using unix socket (optional)",
		},
		cli.StringFlag{
			Name: "signal-on-hup",
			Value: "TERM",
			Usage: "name of the signal to be sent to the server process when start_server\nreceives a SIGHUP. If you use this option, be sure to\nalso use '--signal-on-term' below.",
		},
		cli.StringFlag{
			Name: "signal-on-term",
			Value: "TERM",
			Usage: "name of the signal to be sent to the server process when start_server\nreceives a SIGTERM",
		},
		cli.StringFlag{
			Name: "pid-file",
			Usage: "if set, writes the process id of the start_server process to the file",
		},
		cli.StringFlag{
			Name: "status-file",
			Usage: "if set, writes the status of the server process(es) to the file",
		},
		cli.StringFlag{
			Name: "envdir",
			Usage: "directory that contains environment variables to the server processes.\nIt is intended for use with \"envdir\" in \"daemontools\". This can be\noverwritten by environment variable \"ENVDIR\".",
		},
		cli.BoolFlag{
			Name: "enable-auto-restart",
			Usage: "enables automatic restart by time. This can be overwritten by\nenvironment variable \"ENABLE_AUTO_RESTART\".",
		},
		cli.Int64Flag{
			Name: "auto-restart-interval",
			Value: 360,
			Usage: "automatic restart interval. It is used with\n\"--enable-auto-restart\" option. This can be overwritten by environment\nvariable \"AUTO_RESTART_INTERVAL\".",
		},
		cli.Int64Flag{
			Name: "kill-old-delay",
			Usage: "time to suspend to send a signal to the old worker. The default value is\n5 when \"--enable-auto-restart\" is set, 0 otherwise. This can be\noverwritten by environment variable \"KILL_OLD_DELAY\".",
		},
		cli.BoolFlag{
			Name: "restart",
			Usage: "this is a wrapper command that reads the pid of the start_server process\nfrom --pid-file, sends SIGHUP to the process and waits until the\nserver(s) of the older generation(s) die by monitoring the contents of\nthe --status-file",
		},
	}
	app.Action = func(c *cli.Context) error {
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}