package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/silvergama/transations/account"
	"github.com/silvergama/transations/config"
	"github.com/silvergama/transations/infrastructure"
	"github.com/silvergama/transations/pkg/logger"
	"github.com/silvergama/transations/transaction"
	httpSwagger "github.com/swaggo/http-swagger"
	"go.uber.org/zap"

	accountRepo "github.com/silvergama/transations/account/postgres"
	handler "github.com/silvergama/transations/internal/http/mux"
	transactionRepo "github.com/silvergama/transations/transaction/postgres"
)

// RequestIDMiddleware is a middleware that adds a request ID to the context
func RequestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := r.Header.Get("X-Request-ID")
		if requestID == "" {
			requestID = uuid.New().String()
		}

		ctx := context.WithValue(r.Context(), "request_id", requestID)

		w.Header().Set("X-Request-ID", requestID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Start initializes and starts the API server
func Start(cfg config.Config) {
	stringConnection := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.Database.User,
		cfg.Database.Pwd,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Base,
	)

	db, err := infrastructure.NewDatabase(stringConnection)
	if err != nil {
		logger.Error(
			"failed to create an instance of the database",
			zap.Error(err),
		)
		return
	}

	accountRepo := accountRepo.NewAccount(db.Connection)
	accountService := account.NewService(accountRepo)
	accountHandler := handler.NewAccountHandler(accountService)

	transactionRepo := transactionRepo.NewTransaction(db.Connection)
	transactionService := transaction.NewService(transactionRepo)
	transactionHandler := handler.NewTransactionHandler(transactionService)

	router := mux.NewRouter()

	router.Use(RequestIDMiddleware)

	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	router.HandleFunc("/accounts/{id:[0-9]+}", accountHandler.GetAccountHandler).Methods(http.MethodGet)
	router.HandleFunc("/accounts", accountHandler.CreateAccountHandler).Methods(http.MethodPost)

	router.HandleFunc("/transactions", transactionHandler.CreateTransactionHandler).Methods(http.MethodPost)

	srv := &http.Server{
		Handler:      router,
		Addr:         fmt.Sprintf(":%d", cfg.ServerHTTP.Port),
		WriteTimeout: cfg.ServerHTTP.WriteTimeout,
		ReadTimeout:  cfg.ServerHTTP.ReadTimeout,
	}

	logger.Info(fmt.Sprintf("listening on %d", cfg.ServerHTTP.Port))
	logger.Fatal("failed to run the server", zap.Any("server", srv.ListenAndServe()))
}
