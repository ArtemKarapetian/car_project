package service

import (
	"net/http"
)

type CarHandler interface {
	AddCars(http.ResponseWriter, *http.Request)
	DeleteCar(http.ResponseWriter, *http.Request)
	GetCars(http.ResponseWriter, *http.Request)
	UpdateCar(http.ResponseWriter, *http.Request)
}

type Server struct {
	carHandler CarHandler
}

func NewServer(carHandler CarHandler) *Server {
	return &Server{
		carHandler: carHandler,
	}
}
