package mux

import (
	"encoding/json"
	"net/http"

	"github.com/silvergama/transations/pkg/logger"
	"github.com/silvergama/transations/transaction"
	"go.uber.org/zap"
)

type TransactionHandler struct {
	transactionService transaction.UseCase
}

func NewTransactionHandler(transactionService transaction.UseCase) *TransactionHandler {
	return &TransactionHandler{
		transactionService: transactionService,
	}
}

func (h *TransactionHandler) CreateTransactionHandler(w http.ResponseWriter, r *http.Request) {
	var requestTransaction transaction.Transaction
	if err := json.NewDecoder(r.Body).Decode(&requestTransaction); err != nil {
		logger.Error("failed to decoding json", zap.Error(err), zap.Any("context", r.Context()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if !isValidOperationType(requestTransaction.OperationTypeID) {
		logger.Error("type of operation not permitted",
			zap.Any("operation_type", requestTransaction.OperationTypeID),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	transactionID, err := h.transactionService.Create(r.Context(), &requestTransaction)
	if err != nil {
		logger.Error("failed to create transaction",
			zap.Error(err),
			zap.Any("transaction", requestTransaction),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"account_id": transactionID,
		"message":    "Conta criada com sucesso",
	}

	JSONResponse(w, http.StatusCreated, response)
}

// Função de validação básica no handler
func isValidOperationType(opType transaction.OperationType) bool {
	switch opType {
	case transaction.Purchase, transaction.Installment, transaction.Withdrawal, transaction.Payment:
		return true
	default:
		return false
	}
}
