package httpserver

import (
	"net/http"
)

type carHandler interface {
	AddCars(http.ResponseWriter, *http.Request)
	DeleteCar(http.ResponseWriter, *http.Request)
	GetCars(http.ResponseWriter, *http.Request)
	UpdateCar(http.ResponseWriter, *http.Request)
}

type Server struct {
	carHandler carHandler
}

func NewServer(carHandler carHandler) *Server {
	return &Server{
		carHandler: carHandler,
	}
}
