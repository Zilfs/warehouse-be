package config

import (
	"fmt"
	"strconv"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	DB *gorm.DB
}

func (cfg Config) ConnectionPostgres() (*Postgres, error) {
	dbConnString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		cfg.Pgsql.User,
		cfg.Pgsql.Password,
		cfg.Pgsql.Host,
		cfg.Pgsql.Port,
		cfg.Pgsql.DBName,
	)

	db, err := gorm.Open(postgres.Open(dbConnString), &gorm.Config{})
	if err != nil {
		log.Error().Err(err).Msg("[Connection Postgres-1] Failed to connect to database " + cfg.Pgsql.Host)
		return nil, err
	}

	pgsqlDB, err := db.DB()
	if err != nil {
		log.Error().Err(err).Msg("[Connection Postgres-2] Failed to get database instance")
		return nil, err
	}

	maxOpen, _ := strconv.Atoi(cfg.Pgsql.DBMaxOpen)
	maxIdle, _ := strconv.Atoi(cfg.Pgsql.DBMaxIdle)

	pgsqlDB.SetMaxOpenConns(maxOpen)
	pgsqlDB.SetMaxIdleConns(maxIdle)

	return &Postgres{DB: db}, nil
}
