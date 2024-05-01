package car

import (
	"car_project/internal/model"
	"encoding/json"
	"go.uber.org/zap"
	"io"
	"net/http"
)

func (h *Handler) UpdateCar(w http.ResponseWriter, req *http.Request) {
	regNum := req.URL.Query().Get("regNum")
	if regNum == "" {
		logger.Error("Missing regNum in request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	car, err := readUpdateCarRequest(req)
	if err != nil {
		logger.Error("Failed to read update car request", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	logger.Debug("Attempting to update car", zap.String("regNum", regNum), zap.Any("car", car))
	err = h.repo.UpdateCar(req.Context(), regNum, &car)
	if err != nil {
		logger.Error("Failed to update car in repository", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	logger.Info("Successfully updated car", zap.String("regNum", regNum))
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
