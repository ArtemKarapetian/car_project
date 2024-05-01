package car

import (
	"encoding/json"
	"io"
	"net/http"
)

func (h *Handler) AddCars(w http.ResponseWriter, req *http.Request) {
	cars, err := readAddCarsRequest(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = h.repo.AddCars(req.Context(), cars)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func readAddCarsRequest(req *http.Request) ([]string, error) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}
	var unm []string
	if err = json.Unmarshal(body, &unm); err != nil {
		return nil, err
	}
	return unm, nil
}
