package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/silvergama/transations/config"
	"github.com/silvergama/transations/internal/account"
	v1 "github.com/silvergama/transations/internal/http/mux/v1"
	"github.com/silvergama/transations/internal/transaction"
	"github.com/silvergama/transations/pkg/logger"
	httpSwagger "github.com/swaggo/http-swagger"
	"go.uber.org/zap"
)

type Services struct {
	Account     account.UseCase
	Transaction transaction.UseCase
}

func StartServer(cfg config.Config, services Services) {

	router := mux.NewRouter()

	router.Use(RequestIDMiddleware)

	routesV1 := router.PathPrefix("/v1").Subrouter()
	routesV1.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	accountHandler := v1.NewAccountHandler(services.Account)
	routesV1.HandleFunc("/accounts/{id}", accountHandler.GetAccount).Methods(http.MethodGet)

	routesV1.HandleFunc("/accounts", accountHandler.Create).Methods(http.MethodPost)

	trHandler := v1.NewTransactionHandler(services.Transaction)
	routesV1.HandleFunc("/transactions", trHandler.Create).Methods(http.MethodPost)

	srv := &http.Server{
		Handler:      router,
		Addr:         fmt.Sprintf(":%d", cfg.ServerHTTP.Port),
		WriteTimeout: cfg.ServerHTTP.WriteTimeout,
		ReadTimeout:  cfg.ServerHTTP.ReadTimeout,
	}

	logger.Info(fmt.Sprintf("listening on %d", cfg.ServerHTTP.Port))
	logger.Fatal("failed to run the server", zap.Any("server", srv.ListenAndServe()))
}
