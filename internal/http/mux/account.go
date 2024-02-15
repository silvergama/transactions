package mux

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/silvergama/transations/account"
	"github.com/silvergama/transations/pkg/logger"
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
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	acc, err := h.accountService.GetByID(r.Context(), accountID)
	if err != nil {
		logger.Warn("failed to get account by account_id",
			zap.Error(err),
			zap.Int("account_id", accountID))
		w.WriteHeader(http.StatusNotFound)
		return
	}

	JSONResponse(w, http.StatusCreated, acc)
}

func (h *AccountHandler) CreateAccountHandler(w http.ResponseWriter, r *http.Request) {
	var requestAccount account.Account

	if err := json.NewDecoder(r.Body).Decode(&requestAccount); err != nil {
		logger.Error("failed to decoding json", zap.Error(err), zap.Any("context", r.Context()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if requestAccount.DocumentNumber == "" {
		logger.Warn("document number not found", zap.Any("context", r.Context()))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	accountID, err := h.accountService.Create(r.Context(), &requestAccount)
	if err != nil {
		logger.Error("failed to create account", zap.Error(err), zap.Any("document_id", requestAccount.AccoundID))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := map[string]int{
		"account_id": accountID,
	}

	JSONResponse(w, http.StatusCreated, response)
}
