package config

import (
	"database/sql"
	"fmt"
	"net/url"
	"sync"
	"time"

	_ "github.com/lib/pq"
)

type Database struct {
	DB           *sql.DB
	Mutex        *sync.RWMutex
	Driver       string
	QueryTimeout time.Duration
	ExecTimout   time.Duration
}

type PgConnString struct {
	Host            string
	User            string
	Password        string
	Name            string
	SSLModeDisabled string
}

const (
	PostgresDriver = "postgres"
)

func NewDatabase(driver string) *Database {
	return &Database{Mutex: &sync.RWMutex{}, Driver: driver}
}

func (db *Database) Connect(connStr string) error {
	AppLogger.Infof("Locking connection to %s database", db.Driver)
	db.Mutex.Lock()
	defer func() {
		db.Mutex.Unlock()
		AppLogger.Infof("Unlocking connection to %s database", db.Driver)
	}()
	if db.DB != nil {
		db.DB.Close()
	}
	var err error
	if db.DB, err = sql.Open(db.Driver, connStr); err != nil {
		AppLogger.Error(err)
		return err
	}
	AppLogger.Infof("Connection to %s database established", db.Driver)
	db.DB.SetMaxOpenConns(10)
	db.DB.SetMaxIdleConns(5)
	return db.DB.Ping()
}

func (c *PgConnString) BuildConnectionString() string {
	if c.Host == "" {
		return ""
	}
	s := fmt.Sprintf("postgres://%s:%s@%s/%s",
		url.QueryEscape(c.User), url.QueryEscape(c.Password), url.QueryEscape(c.Host), url.QueryEscape(c.Name))
	if c.SSLModeDisabled == "1" {
		s = s + "?sslmode=disable"
	}
	return s
}
