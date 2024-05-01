package car

import (
	"go.uber.org/zap"
	"net/http"
)

func (h *Handler) DeleteCar(w http.ResponseWriter, req *http.Request) {
	regNum := req.URL.Query().Get("regNum")
	if regNum == "" {
		logger.Error("Missing regNum in request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	logger.Debug("Attempting to delete car", zap.String("regNum", regNum))
	err := h.repo.DeleteCar(req.Context(), regNum)
	if err != nil {
		logger.Error("Failed to delete car in repository", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	logger.Info("Successfully deleted car", zap.String("regNum", regNum))
	w.WriteHeader(http.StatusOK)
}
