package httpdelivery

import (
	"net/http"

	"github.com/Helltale/tz-telecom/internal/delivery/httpdelivery/middleware"
	"github.com/Helltale/tz-telecom/internal/usecase"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func NewRouter(userUC usecase.UserUseCaseInterface, orderWorker *usecase.OrderWorker) http.Handler {
	mux := http.NewServeMux()

	userHandler := NewUserHandler(userUC)
	orderHandler := NewOrderHandler(orderWorker)

	mux.Handle("/users/register", middleware.Chain(
		otelhttp.NewHandler(http.HandlerFunc(userHandler.RegisterUserHandler), "RegisterUser"),
		middleware.Logging, middleware.Recover))

	mux.Handle("/orders", middleware.Chain(
		otelhttp.NewHandler(http.HandlerFunc(orderHandler.CreateOrderHandler), "CreateOrder"),
		middleware.Logging, middleware.Recover))

	return mux
}
