package car

import (
	"car_project/internal/model"
	"encoding/json"
	"io"
	"net/http"
)

func (h *Handler) UpdateCar(w http.ResponseWriter, req *http.Request) {
	regNum := req.URL.Query().Get("regNum")
	if regNum == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	car, err := readUpdateCarRequest(req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.repo.UpdateCar(req.Context(), regNum, &car)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func readUpdateCarRequest(req *http.Request) (model.Car, error) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return model.Car{}, err
	}
	var unm model.Car
	err = json.Unmarshal(body, &unm)
	if err != nil {
		return model.Car{}, err
	}
	return unm, nil
}
