package car

import (
	"car_project/internal/model"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

func (h *Handler) GetCars(w http.ResponseWriter, req *http.Request) {
	filter, err := readGetCarsRequest(req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	limit, offset, err := getLimitOffset(req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	cars, err := h.repo.GetCars(req.Context(), &filter, limit, offset)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	body, err := json.Marshal(cars)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = w.Write(body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func readGetCarsRequest(req *http.Request) (model.Car, error) {
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

func getLimitOffset(req *http.Request) (int, int, error) {
	limitString := req.URL.Query().Get("limit")
	offsetString := req.URL.Query().Get("offset")
	if limitString == "" {
		limitString = "10"
	}
	limit, err := strconv.Atoi(limitString)
	if err != nil {
		return 0, 0, err
	}
	offset, err := strconv.Atoi(offsetString)
	if err != nil {
		return 0, 0, err
	}

	return limit, offset, nil
}
