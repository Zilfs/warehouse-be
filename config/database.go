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

func (cfg Config) ConnetionPostgres() (*Postgres, error) {
	dbConnString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		cfg.PgsqlDB.User,
		cfg.PgsqlDB.Password,
		cfg.PgsqlDB.Host,
		cfg.PgsqlDB.Port,
		cfg.PgsqlDB.DBName,
	)

	db, err := gorm.Open(postgres.Open(dbConnString), &gorm.Config{})
	if err != nil {
		log.Error().Err(err).Msg("[Connection Postgres-1] Failed to connect to database " + cfg.PgsqlDB.Host)
		return nil, err
	}

	pgsqlDB, err := db.DB()
	if err != nil {
		log.Error().Err(err).Msg("[Connection Postgres-2] Failed to get database instance")
		return nil, err
	}

	maxOpen, _ := strconv.Atoi(cfg.PgsqlDB.DBMaxOpen)
	maxIdle, _ := strconv.Atoi(cfg.PgsqlDB.DBMaxIdle)

	pgsqlDB.SetMaxOpenConns(maxOpen)
	pgsqlDB.SetMaxIdleConns(maxIdle)

	return &Postgres{DB: db}, nil
}
