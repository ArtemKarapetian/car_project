package car

import "net/http"

func (h *Handler) DeleteCar(w http.ResponseWriter, req *http.Request) {
	regNum := req.URL.Query().Get("regNum")
	if regNum == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err := h.repo.DeleteCar(req.Context(), regNum)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
