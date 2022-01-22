package internal

import (
	"github.com/Uchencho/commons/ctime"
	"github.com/Uchencho/commons/uuid"
)

// Transaction is a representation of a transaction
type Transaction struct {
	ID                   uuid.V4     `json:"id"`
	AccountID            uuid.V4     `json:"accountId"`
	Amount               int64       `json:"amount"`
	Status               string      `json:"status"`
	CreatedTimestamp     ctime.Epoch `json:"createdTimestamp"`
	LastUpdatedTimestamp ctime.Epoch `json:"lastUpdatedTimestamp"`
}
