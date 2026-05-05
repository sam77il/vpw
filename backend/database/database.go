package database

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect() *pgxpool.Pool {
	pgUrl := os.Getenv("POSTGRES")
	config, err := pgxpool.ParseConfig(pgUrl)
	if err != nil {
		log.Fatalf("Error parsing pgx config: %v", err)
	}
	config.MinConns = 0
	config.MaxConns = 10
	config.MaxConnLifetime = time.Hour
	config.MaxConnIdleTime = time.Minute * 30

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("Error creating pgx: %v", err)
	}

	if err := pool.Ping(context.Background()); err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	log.Println("Successfully connected to postgres database")

	return pool
}
