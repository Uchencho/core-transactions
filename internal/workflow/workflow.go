package workflow

import (
	"context"

	"github.com/Uchencho/commons/commonerror"
	"github.com/Uchencho/commons/uuid"
	"github.com/Uchencho/core-proto/generated/accounts"
	"github.com/Uchencho/core-transactions/internal/db"
	"github.com/Uchencho/core-transactions/pkg"
	"github.com/pkg/errors"
)

func GetTransaction(accountClient accounts.ClientClient,
	getTrasnaction db.GetTransactionFunc) GetTransactionFunc {
	return func(accountId uuid.V4) (pkg.Transaction, error) {

		req := &accounts.GetAccountRequest{AccountId: string(accountId)}

		if _, err := accountClient.GetAccount(context.Background(), req); err != nil {
			return pkg.Transaction{}, commonerror.NewErrorParams("accountId", "invalid account id").ToBadRequest()
		}

		transaction, err := getTrasnaction(string(accountId))
		if err != nil {
			return pkg.Transaction{}, errors.Wrap(err, "unable to get transaction")
		}

		return pkg.Transaction{
			ID:                   string(transaction.ID),
			AccountID:            string(transaction.AccountID),
			Amount:               transaction.Amount,
			Status:               transaction.Status,
			CreatedTimestamp:     transaction.CreatedTimestamp.String(),
			LastUpdatedTimestamp: transaction.LastUpdatedTimestamp.String(),
		}, nil
	}
}
