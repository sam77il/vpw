package sugar

import (
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Config struct {
	Host    	string
	Cors    	CorsSettings
	Timeout 	time.Duration
	Database	*pgxpool.Pool
}
