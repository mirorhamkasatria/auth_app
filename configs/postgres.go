package configs

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func GORMOpenConn(c *Config) (*gorm.DB, error) {

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.DbHost, c.DbPort, c.DbUsername, c.DbPassword, c.DbName)

	// init GORM configuration
	dbConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	}
	// open connection to database using driver PostgreSQL
	db, err := gorm.Open(postgres.Open(dsn), dbConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Set database connection settings.
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to configure database connection pool: %w", err)
	}

	sqlDB.SetMaxOpenConns(c.MaxConn)
	sqlDB.SetMaxIdleConns(c.MaxIdleConn)
	sqlDB.SetConnMaxLifetime(time.Duration(c.MaxLifetimeConn))

	// Try to ping database.
	err = sqlDB.Ping()
	if err != nil {
		defer sqlDB.Close() // close database connection
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil
}
