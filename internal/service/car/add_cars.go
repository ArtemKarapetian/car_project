package car

import (
	"encoding/json"
	"go.uber.org/zap"
	"io"
	"net/http"
)

func (h *Handler) AddCars(w http.ResponseWriter, req *http.Request) {
	cars, err := readAddCarsRequest(req)
	if err != nil {
		logger.Error("Failed to read add cars request", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	logger.Debug("Attempting to add cars", zap.Strings("cars", cars))
	err = h.repo.AddCars(req.Context(), cars)
	if err != nil {
		logger.Error("Failed to add cars in repository", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	logger.Info("Successfully added cars", zap.Strings("cars", cars))
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
