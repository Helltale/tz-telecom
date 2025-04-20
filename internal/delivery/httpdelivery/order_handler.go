package httpdelivery

import (
	"encoding/json"
	"net/http"

	"github.com/Helltale/tz-telecom/internal/models"
	"github.com/Helltale/tz-telecom/internal/usecase"
)

type OrderHandler struct {
	worker *usecase.OrderWorker
}

func NewOrderHandler(w *usecase.OrderWorker) *OrderHandler {
	return &OrderHandler{worker: w}
}

func (h *OrderHandler) CreateOrderHandler(wr http.ResponseWriter, r *http.Request) {
	var req models.OrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(wr, "invalid input", http.StatusBadRequest)
		return
	}

	h.worker.Enqueue(usecase.OrderJob{
		UserID: req.UserID,
		Items:  req.ToDomainItems(),
	})

	wr.WriteHeader(http.StatusAccepted)
}
