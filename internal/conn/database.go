package conn

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"ddr13/riddlercore/internal/config"
)

// DB holds the database instance
var db *gorm.DB

// Ping tests if db connection is alive
func Ping() error {
	if db == nil {
		return fmt.Errorf("database connection is not initialized")
	}
	return db.Exec("SELECT 'DBD::Pg ping test';").Error
}

// Connect sets the db client of database using configuration cfg
func Connect(cfg *config.Database) error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", cfg.Host, cfg.Username, cfg.Password, cfg.Name, cfg.Port)
	if cfg.Options != nil {
		for k, v := range cfg.Options {
			dsn += fmt.Sprintf(" %s=%s", k, v[0])
		}
	}
	// open a database connection using gorm ORM
	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	db = d

	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	if cfg.MaxIdleConn != 0 {
		sqlDB.SetMaxIdleConns(cfg.MaxIdleConn)
	}
	if cfg.MaxOpenConn != 0 {
		sqlDB.SetMaxOpenConns(cfg.MaxOpenConn)
	}
	if cfg.MaxConnLifetime.Seconds() != 0 {
		sqlDB.SetConnMaxLifetime(cfg.MaxConnLifetime)
	}

	return nil
}

// DefaultDB returns default db
func DefaultDB() *gorm.DB {
	if db == nil {
		return nil
	}
	return db.Debug()
}

// CloseDB closes the db connection
func CloseDB() error {
	if db == nil {
		return nil
	}
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	return sqlDB.Close()
}

// ConnectDB sets the db client of database using default configuration file
func ConnectDB() error {
	cfg := config.DB()
	connectionRenew() //start a connection re-newer
	return Connect(cfg)
}

func connectionRenew() {
	ticker := time.NewTicker(30 * time.Second)
	go func() {
		for t := range ticker.C {
			if err := Ping(); err != nil {
				log.Printf("error: %v [re-connecting database]", err.Error())
				Connect(config.DB())
				_ = t
			}
		}
	}()
}
