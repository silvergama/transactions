package v1

import (
	"encoding/json"
	"net/http"

	"github.com/silvergama/transations/internal/domain"
	"github.com/silvergama/transations/internal/usecase/transaction"
	"github.com/silvergama/transations/pkg/logger"
	"github.com/silvergama/transations/pkg/response"
	"go.uber.org/zap"
)

// TransactionHandler is responsible for handling HTTP requests related to transactions
type transactionHandler struct {
	transactionService transaction.UseCase
}

// NewTransactionHandler creates a new instance of TransactionHandler
func NewTransactionHandler(transactionService transaction.UseCase) *transactionHandler {
	return &transactionHandler{
		transactionService: transactionService,
	}
}

// Create handles the creation of a new transaction.
// @Summary Create a new transaction
// @Description Create a new transaction
// @ID create-transaction
// @Accept json
// @Produce json
// @Param transaction body transaction.Transaction true "Transaction object to be created"
// @Success      200 {object} response.Response
// @Failure      400  {object}  response.Error
// @Failure      404  {object}  response.Error
// @Failure      500  {object}  response.Error
// @Router /transaction [post]
func (h *transactionHandler) Create(w http.ResponseWriter, r *http.Request) {
	var requestTransaction domain.Transaction
	if err := json.NewDecoder(r.Body).Decode(&requestTransaction); err != nil {
		logger.Warn("failed to decoding json", zap.Error(err))
		response.WriteBadRequest(w, "failed to decode payload")
		return
	}

	if !isValidOperationType(requestTransaction.OperationTypeID) {
		logger.Warn("type of operation not permitted",
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
	response.Write(w, resp, http.StatusOK)

}

// isValidOperationType checks if the provided operation type is valid
func isValidOperationType(opType domain.OperationType) bool {
	switch opType {
	case domain.Purchase, domain.Installment, domain.Withdrawal, domain.Payment:
		return true
	default:
		return false
	}
}
