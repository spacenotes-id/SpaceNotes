package postgres

import (
	"context"

	"github.com/gofiber/fiber/v2/log"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/spacenotes-id/SpaceNotes/server/config"
	"github.com/spacenotes-id/SpaceNotes/server/database/postgres/sqlc"
)

var (
	Pool *pgxpool.Pool
	err  error
)

func init() {
	Pool, err = pgxpool.New(context.Background(), config.PostgresURL)
	if err != nil {
		log.Fatal(err)
	}
}

func GetPostgresSQLCQuerier() (*sqlc.Queries, error) {
	db := sqlc.New(Pool)

	log.Info("Connected to DB")

	return db, nil
}
