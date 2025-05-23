package postgres

const (
	paymentsTableName = "payments"

	idColumn             = "id"
	userIDColumn         = "user_id"
	subscriptionIDColumn = "subscription_id"
	amountColumn         = "amount"
	currencyColumn       = "currency"
	statusColumn         = "status"
	gatewayIDColumn      = "gateway_id"
	createdAtColumn      = "created_at"
	updatedAtColumn      = "updated_at"
)

var paymentsColumns = []string{
	idColumn, userIDColumn, subscriptionIDColumn, amountColumn, currencyColumn, statusColumn, gatewayIDColumn,
	createdAtColumn, updatedAtColumn,
}
