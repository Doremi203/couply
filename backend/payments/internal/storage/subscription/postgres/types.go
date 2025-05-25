package postgres

const (
	subscriptionsTableName = "subscriptions"

	idColumn        = "id"
	userIDColumn    = "user_id"
	planColumn      = "plan"
	statusColumn    = "status"
	autoRenewColumn = "auto_renew"
	startDateColumn = "start_date"
	endDateColumn   = "end_date"
)

var subscriptionsColumns = []string{
	idColumn, userIDColumn, planColumn, statusColumn, autoRenewColumn, startDateColumn, endDateColumn,
}
