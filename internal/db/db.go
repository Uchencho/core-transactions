package db

import "github.com/Uchencho/core-transactions/internal"

type GetTransactionFunc func(accountId string) (internal.Transaction, error)
