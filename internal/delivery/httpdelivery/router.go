package httpdelivery

import (
	"net/http"

	"github.com/Helltale/tz-telecom/internal/delivery/httpdelivery/middleware"
	"github.com/Helltale/tz-telecom/internal/usecase"
)

func NewRouter(userUC usecase.UserUseCaseInterface, orderWorker *usecase.OrderWorker) http.Handler {
	mux := http.NewServeMux()

	userHandler := NewUserHandler(userUC)
	orderHandler := NewOrderHandler(orderWorker)

	mux.Handle("/users/register", middleware.Chain(http.HandlerFunc(userHandler.RegisterUserHandler), middleware.Logging, middleware.Recover))
	mux.Handle("/orders", middleware.Chain(http.HandlerFunc(orderHandler.CreateOrderHandler), middleware.Logging, middleware.Recover))
	return mux
}
