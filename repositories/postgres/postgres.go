package postgres

import (
	"NewProj1/internal/config"
	"NewProj1/repositories"
	"context"
	"database/sql"
	"fmt"
)


type postgres struct {
	db *sql.DB
}

func dsn(cfg config.Postgres) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBName, cfg.SSL,
	)
}

func New(ctx context.Context, cfg config.Postgres) (repositories.DB, error) {
	db, err := sql.Open("postgres", dsn(cfg))
	if err != nil {
		return nil, fmt.Errorf("failed to open the database driver: %v", err)
	}

	err = db.PingContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("can not ping the database: %v", err)
	}

	return &postgres{db: db}, nil
}


func (p *postgres) Close(ctx context.Context) error {
	err := p.db.Close()
	if err != nil {
		return fmt.Errorf("can not close the database: %v", err)
	}

	return nil
}