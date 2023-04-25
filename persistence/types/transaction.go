package types

// TODO implement a transaction type for SQL interaction
const (
	TransactionTableName   = "transaction"
	TransactionTableSchema = `(
		key text,
		value text
	)`
)
