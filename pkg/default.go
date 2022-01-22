package pkg

// Transaction is a representation of a transaction
type Transaction struct {
	ID                   string `json:"id"`
	AccountID            string `json:"accountId"`
	Amount               int64  `json:"amount"`
	Status               string `json:"status"`
	CreatedTimestamp     string `json:"createdTimestamp"`
	LastUpdatedTimestamp string `json:"lastUpdatedTimestamp"`
}
