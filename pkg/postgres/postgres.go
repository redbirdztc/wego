package postgres

import (
	"errors"
	"fmt"
	"time"

	"github.com/redbirdztc/wego/pkg/db"
	"github.com/redbirdztc/wego/pkg/loglevel"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var _ db.ConnectionKeeper = (*PostgresDB)(nil)

type PostgresDB struct {
	DSN  string
	Gorm *gorm.DB
}

func (db *PostgresDB) GetConnection() *gorm.DB {
	if db.Gorm != nil {
		return db.Gorm
	}
	return nil
}

func NewPostgresDB(dsn string) *PostgresDB {
	if dsn == "" {
		panic(errors.New("PostgresDSN is not set"))
	}

	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // Use singular table names
		},
	})
	if err != nil {
		panic(fmt.Errorf("failed to connect to Postgres: %w", err))
	}

	sqlDB, err := gormDB.DB()
	if err != nil {
		panic(fmt.Errorf("failed to get DB from gormDB: %w", err))
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	gormDB.Logger = gormDB.Logger.LogMode(3)
	if loglevel.LogLevel() == loglevel.LogLevelDebug {
		gormDB = gormDB.Debug()
	}

	if err := sqlDB.Ping(); err != nil {
		panic(fmt.Errorf("failed to ping Postgres: %w", err))
	}

	return &PostgresDB{DSN: dsn, Gorm: gormDB}
}
