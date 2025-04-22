package idempotencypostgres

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/auth/pkg/idempotency"
	"github.com/Doremi203/couply/backend/auth/pkg/postgres"
	"github.com/goccy/go-json"
)

func NewRepo(db postgres.Client) *repo {
	return &repo{
		db: db,
	}
}

type repo struct {
	db postgres.Client
}

func (r *repo) Create(ctx context.Context, key idempotency.Key) error {
	const query = `
		INSERT INTO idempotency_requests (idempotency_key)
		VALUES ($1)
		ON CONFLICT (idempotency_key) DO NOTHING;
	`
	res, err := r.db.Exec(ctx, query, key)
	if err != nil {
		return errors.WrapFailf(err, "create idempotency request for %v", errors.Token("key", key))
	}
	if res.RowsAffected() == 0 {
		return idempotency.ErrAlreadyProcessed
	}

	return nil
}

func (r *repo) UpdateData(ctx context.Context, key idempotency.Key, data any) error {
	const query = `
		UPDATE idempotency_requests
		SET result = $2
		WHERE idempotency_key = $1
	`

	jsonData, err := json.Marshal(data)
	if err != nil {
		return errors.WrapFailf(err, "marshal json")
	}

	res, err := r.db.Exec(ctx, query, key, jsonData)
	if err != nil {
		return errors.WrapFailf(err, "create idempotency request for %v", errors.Token("key", key))
	}
	if res.RowsAffected() == 0 {
		return idempotency.ErrNotBeingProcessed
	}

	return nil
}

func (r *repo) GetData(ctx context.Context, key idempotency.Key, data any) error {
	const query = `
		SELECT 
		    result
		FROM idempotency_requests
		WHERE idempotency_key = $1
	`

	var result []byte

	res := r.db.QueryRow(ctx, query, key)
	err := res.Scan(&result)
	if err != nil {
		return errors.WrapFailf(err, "get data for %v", errors.Token("key", key))
	}

	err = json.Unmarshal(result, &data)
	if err != nil {
		return errors.WrapFailf(err, "unmarshal json")
	}

	return nil
}
