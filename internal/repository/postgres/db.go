package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

func NewDB(ctx context.Context, host, user, password, dbname string, port int) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		host, user, password, dbname, port,
	)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	pingCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err = db.PingContext(pingCtx); err != nil {
		_ = db.Close()
		return nil, err
	}
	fmt.Println("Connected to database")
	return db, nil
}
