package workflow

import (
	"github.com/Uchencho/commons/uuid"
	"github.com/Uchencho/core-transactions/pkg"
)

type GetTransactionFunc func(accountId uuid.V4) (pkg.Transaction, error)
