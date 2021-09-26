package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/nathan-osman/pratl/db"
	"github.com/nathan-osman/pratl/server"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "pratl"
	app.Usage = "self-contained chat server"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "db-host",
			Value:  "postgres",
			EnvVar: "DB_HOST",
			Usage:  "PostgreSQL database host",
		},
		cli.IntFlag{
			Name:   "db-port",
			Value:  5432,
			EnvVar: "DB_PORT",
			Usage:  "PostgreSQL database port",
		},
		cli.StringFlag{
			Name:   "db-name",
			Value:  "postgres",
			EnvVar: "DB_NAME",
			Usage:  "PostgreSQL database name",
		},
		cli.StringFlag{
			Name:   "db-user",
			Value:  "postgres",
			EnvVar: "DB_NAME",
			Usage:  "PostgreSQL database user",
		},
		cli.StringFlag{
			Name:   "db-password",
			Value:  "postgres",
			EnvVar: "DB_PASSWORD",
			Usage:  "PostgreSQL database password",
		},
		cli.StringFlag{
			Name:   "server-addr",
			Value:  ":http",
			EnvVar: "SERVER_ADDR",
			Usage:  "address for server to listen on",
		},
	}
	app.Action = func(ctx *cli.Context) error {

		// Connect to the database
		d, err := db.New(&db.Config{
			Host:     ctx.String("db-host"),
			Port:     ctx.Int("db-port"),
			Name:     ctx.String("db-name"),
			User:     ctx.String("db-user"),
			Password: ctx.String("db-password"),
		})
		if err != nil {
			return err
		}
		defer d.Close()

		// Run migrations
		if err := d.Migrate(); err != nil {
			return err
		}

		// Start the server
		s, err := server.New(&server.Config{
			Addr: ctx.String("server-addr"),
			Conn: d,
		})
		if err != nil {
			return err
		}
		defer s.Close()

		// Wait for SIGINT or SIGTERM
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		<-sigChan

		return nil
	}
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "fatal: %s\n", err.Error())
	}
}
