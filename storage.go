package main

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Storage struct {
	Cfg  Config
	Psql *sqlx.DB
	Gorm *gorm.DB
}

func NewStorage(cfg Config) Storage {
	return Storage{
		Cfg: cfg,
	}
}

func (s *Storage) InitPsql() {
	db, err := sqlx.Connect("postgres", s.Cfg.PSQL_URL)
	if err != nil {
		log("sqlx connect postgres failed")
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		log("psql ping failed")
		panic(err)
	}
	log("psql ping success")
	s.Psql = db
}

func (s *Storage) InitGorm() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect psql using gorm")
	}
	return db
}
