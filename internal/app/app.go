package app

import (
	"car_project/internal/config"
	"car_project/internal/db/postgres"
	"car_project/internal/repository/car"
	"car_project/internal/service"
	carservice "car_project/internal/service/car"
	"context"
	"log"
	"net"
	"net/http"
)

var (
	env  = ".env"
	host = "localhost"
	port = "8080"
)

type App struct {
	serviceProvider *serviceProvider
	server          *service.Server
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Run() error {
	return a.runServer()
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initConfig(_ context.Context) error {
	return config.Load(env)
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

func (a *App) initServer(ctx context.Context) error {
	a.serviceProvider.db, _ = postgres.NewDb(ctx, env)
	a.serviceProvider.repo = car.NewRepository(a.serviceProvider.db)
	a.serviceProvider.handler = carservice.NewHandler(a.serviceProvider.repo)
	a.server = service.NewServer(a.serviceProvider.handler)
	return nil
}

func (a *App) runServer() error {
	http.Handle("/", service.NewRouter(a.server))
	if err := http.ListenAndServe(net.JoinHostPort(host, port), nil); err != nil {
		log.Fatal(err)
	}
	return nil
}
