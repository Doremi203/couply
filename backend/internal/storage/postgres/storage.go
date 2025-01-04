package postgres

type PgStorage struct {
	txManager TransactionManager
}

func NewPgStorage(txManager TransactionManager) *PgStorage {
	return &PgStorage{txManager: txManager}
}
