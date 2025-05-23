package postgres

const (
	userBlocksTableName       = "user_blocks"
	userBlockReasonsTableName = "user_block_reasons"

	idColumn        = "id"
	blockedIdColumn = "blocked_id"
	messageColumn   = "message"
	createdAtColumn = "created_at"
	statusColumn    = "status"

	blockIdColumn = "block_id"
	reasonColumn  = "reason"
)

var (
	userBlocksColumns       = []string{idColumn, blockedIdColumn, messageColumn, createdAtColumn, statusColumn}
	userBlockReasonsColumns = []string{blockIdColumn, reasonColumn}
)
