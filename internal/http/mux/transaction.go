package mux

import (
	"encoding/json"
	"net/http"

	"github.com/silvergama/transations/pkg/logger"
	"github.com/silvergama/transations/pkg/response"
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
		logger.Error("failed to decoding json", zap.Error(err))
		response.WriteBadRequest(w, "failed to decode payload")
		return
	}

	if !isValidOperationType(requestTransaction.OperationTypeID) {
		logger.Error("type of operation not permitted",
			zap.Any("operation_type", requestTransaction.OperationTypeID),
		)
		response.WriteBadRequest(w, "invalid operation type")
		return
	}

	transactionID, err := h.transactionService.Create(r.Context(), &requestTransaction)
	if err != nil {
		logger.Error("failed to create transaction",
			zap.Error(err),
			zap.Any("transaction", requestTransaction),
		)
		response.WriteServerError(w, "failed to create transaction")
		return
	}

	resp := response.Response{
		Message: "transaction created successfully",
		Data: map[string]interface{}{
			"transaction_id": transactionID,
		},
	}
	response.Write(w, resp, http.StatusCreated)

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
