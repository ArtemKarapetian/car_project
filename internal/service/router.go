package httpserver

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const (
	queryParamKey = "id"
)

func NewRouter(s *Server) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/info", func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodPost:
			s.carHandler.AddCars(w, req)
		case http.MethodGet:
			s.carHandler.GetCars(w, req)
		default:
			log.Println("error")
		}
	})

	router.HandleFunc(fmt.Sprintf("/info/{%s:[0-9]+}", queryParamKey), func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodGet:
			s.carHandler.DeleteCar(w, req)
		case http.MethodPut:
			s.carHandler.UpdateCar(w, req)
		default:
			log.Println("error")
		}
	})

	return router
}
