package postgresuow

import "github.com/jackc/pgx/v5/pgxpool"

func New(db *pgxpool.Pool) {

}

type postgresUnitOfWork struct{}
