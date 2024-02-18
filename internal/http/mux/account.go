package mux

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/silvergama/transations/account"
	_ "github.com/silvergama/transations/docs"
	"github.com/silvergama/transations/pkg/logger"
	"github.com/silvergama/transations/pkg/response"
	"go.uber.org/zap"
)

// AccountHandler is responsible for handling HTTP requests related to accounts
type AccountHandler struct {
	accountService account.UseCase
}

// NewAccountHandler creates a new instance of AccountHandler
func NewAccountHandler(accountService account.UseCase) *AccountHandler {
	return &AccountHandler{
		accountService: accountService,
	}
}

// GetAccount godoc
// @Summary      Get an account
// @Description  get string by ID
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Account ID"
// @Success      200  {object}  account.Account
// @Failure      400  {object}  response.Error
// @Failure      404  {object}  response.Error
// @Failure      500  {object}  response.Error
// @Router       /accounts/{id} [get]
func (h *AccountHandler) GetAccountHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	accountIDStr := params["id"]

	accountID, err := strconv.Atoi(accountIDStr)
	if err != nil {
		logger.Error("failed to convert string to int",
			zap.Error(err),
			zap.String("account_id", accountIDStr),
			zap.Any("request_id", r.Context().Value("request_id")),
		)

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

// CreateAccountHandler handles the creation of a new account.
// @Summary Create a new account
// @Description Create a new account
// @ID create-account
// @Accept json
// @Produce json
// @Param account body account.Account true "Account object to be created"
// @Success      200 {object} response.Response
// @Failure      400  {object}  response.Error
// @Failure      404  {object}  response.Error
// @Failure      500  {object}  response.Error
// @Router /accounts [post]
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

	response.Write(w, resp, http.StatusOK)
}
