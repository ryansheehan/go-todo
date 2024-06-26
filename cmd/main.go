package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"

	"github.com/ryansheehan/go-todo/internal/todo"
)

func main() {
	esrv := echo.New()
	esrv.Logger.SetLevel(log.DEBUG)

	//flag parsing
	port := flag.String("port", "4000", "port for app")
	flag.Parse()

	//--------------------------
	//middleware provided by echo
	//--------------------------
	//panic recover
	esrv.Use(middleware.Recover())
	//body limit
	esrv.Use(middleware.BodyLimit("35K"))
	//secure header
	esrv.Use(middleware.Secure())
	//timeout
	esrv.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Timeout: 5 * time.Second,
	}))
	//logger
	esrv.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `{"time":"${time_rfc3339_nano}","remote_ip":"${remote_ip}",` +
			`"host":"${host}","method":"${method}","uri":"${uri}",` +
			`"status":${status},"error":"${error}","latency_human":"${latency_human}"` +
			`` + "\n",
		CustomTimeFormat: "2006-01-02 15:04:05.00000",
	}))

	//static
	esrv.Static("/public", "public")
	esrv.Static("/static", "static")

	//init app modules
	todoModule := todo.NewTodoModule()

	//mount module routes
	todoModule.Mount(esrv, "")

	//graceful shutdown
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	go func() {
		if err := esrv.Start(":" + *port); err != nil && err != http.ErrServerClosed {
			esrv.Logger.Fatal("shutting down server")
		}

	}()

	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := esrv.Shutdown(ctx); err != nil {
		esrv.Logger.Fatal(err)
	}
}
