package service

import (
	"fmt"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
)

const (
	queryParamKey = "id"
)

var logger, _ = zap.NewProduction()

func NewRouter(s *Server) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/info", func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodPost:
			logger.Debug("Attempting to add cars")
			s.carHandler.AddCars(w, req)
		case http.MethodGet:
			logger.Debug("Attempting to get cars")
			s.carHandler.GetCars(w, req)
		default:
			logger.Error("Invalid method")
		}
	})

	router.HandleFunc(fmt.Sprintf("/info/{%s:[0-9]+}", queryParamKey), func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodGet:
			logger.Debug("Attempting to delete car")
			s.carHandler.DeleteCar(w, req)
		case http.MethodPut:
			logger.Debug("Attempting to update car")
			s.carHandler.UpdateCar(w, req)
		default:
			logger.Error("Invalid method")
		}
	})

	logger.Info("Router successfully created")
	return router
}
