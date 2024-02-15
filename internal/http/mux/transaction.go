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

// CreateAccountHandler handles the creation of a new transaction.
// @Summary Create a new transaction
// @Description Create a new transaction
// @ID create-transaction
// @Accept json
// @Produce json
// @Param transaction body transaction.Transaction true "Transaction object to be created"
// @Success 201 {object} response.Response
// @Failure      400  {object}  response.Error
// @Failure      404  {object}  response.Error
// @Failure      500  {object}  response.Error
// @Router /transaction [post]
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

func isValidOperationType(opType transaction.OperationType) bool {
	switch opType {
	case transaction.Purchase, transaction.Installment, transaction.Withdrawal, transaction.Payment:
		return true
	default:
		return false
	}
}
