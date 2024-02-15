package mux

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/silvergama/transations/account"
	_ "github.com/silvergama/transations/cmd/api/docs"
	"github.com/silvergama/transations/pkg/logger"
	"github.com/silvergama/transations/pkg/response"
	"go.uber.org/zap"
)

type AccountHandler struct {
	accountService account.UseCase
}

func NewAccountHandler(accountService account.UseCase) *AccountHandler {
	return &AccountHandler{
		accountService: accountService,
	}
}

func (h *AccountHandler) GetAccountHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	accountIDStr := params["id"]

	// Verificar se o ID é um número válido
	accountID, err := strconv.Atoi(accountIDStr)
	if err != nil {
		logger.Error("failed to convert string to int",
			zap.Error(err),
			zap.String("account_id", accountIDStr))
		response.WriteBadRequest(w, "id parameter is different from expected")
		return
	}

	acc, err := h.accountService.GetByID(r.Context(), accountID)
	if err != nil {
		logger.Warn("failed to get account by account_id",
			zap.Error(err),
			zap.Int("account_id", accountID))
		response.WriteNotFound(w, "Unable to find an account with this account_id")
		return
	}

	response.Write(w, acc, http.StatusOK)
}

func (h *AccountHandler) CreateAccountHandler(w http.ResponseWriter, r *http.Request) {
	var requestAccount account.Account

	if err := json.NewDecoder(r.Body).Decode(&requestAccount); err != nil {
		logger.Error("failed to decoding json", zap.Error(err))
		response.WriteBadRequest(w, "failed to decode payload")
		return
	}

	if requestAccount.DocumentNumber == "" {
		logger.Warn("document number not found")
		response.WriteBadRequest(w, "document number not found")
		return
	}

	accountID, err := h.accountService.Create(r.Context(), &requestAccount)
	if err != nil {
		logger.Error("failed to create account", zap.Error(err), zap.Any("document_id", requestAccount.AccoundID))
		response.WriteServerError(w, "failed to create account")
		return
	}

	resp := response.Response{
		Message: "account created successfully",
		Data: map[string]int{
			"account_id": accountID,
		},
	}

	response.Write(w, resp, http.StatusCreated)
}
