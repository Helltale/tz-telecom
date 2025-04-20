package httpdelivery

import (
	"encoding/json"
	"net/http"

	"github.com/Helltale/tz-telecom/internal/domain"
	"github.com/Helltale/tz-telecom/internal/models"
	"github.com/Helltale/tz-telecom/internal/usecase"
)

type UserHandler struct {
	userUC usecase.UserUseCaseInterface
}

func NewUserHandler(uc usecase.UserUseCaseInterface) *UserHandler {
	return &UserHandler{userUC: uc}
}

func (h *UserHandler) RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	var req models.UserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	// simple req validation
	if req.FirstName == "" || req.LastName == "" {
		http.Error(w, "first and last name required", http.StatusBadRequest)
		return
	}

	user := &domain.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Age:       req.Age,
		IsMarried: req.IsMarried,
		Password:  req.Password,
	}

	if err := h.userUC.RegisterUser(r.Context(), user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
