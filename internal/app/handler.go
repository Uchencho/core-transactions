package app

import (
	"net/http"

	"github.com/Uchencho/commons/httputils"
	"github.com/Uchencho/core-proto/generated/accounts"
	"github.com/Uchencho/core-transactions/internal/db"
	"github.com/Uchencho/core-transactions/internal/workflow"
)

const (
	accountIdParam = "accountId"
)

// GetTransactionHandler is the handler in charge of getting a account's transaction
func GetTransactionHandler(accountClient accounts.ClientClient,
	getTrasnaction db.GetTransactionFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		accountId, err := httputils.RetrieveUUIDResource(r, accountIdParam)
		if err != nil {
			httputils.ServeError(err, w)
			return
		}

		wf := workflow.GetTransaction(accountClient, getTrasnaction)
		resp, err := wf(accountId)
		if err != nil {
			httputils.ServeError(err, w)
			return
		}
		httputils.ServeGeneralJSON(resp, w, http.StatusOK)
	}
}
