package db

import (
	"github.com/SmartDuck9000/go-tradex/src/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

type RepoDB interface {
	Open() error
	configureConnectionPools() error

	CreateStat(stat *SavedStat) error
	GetStat(fromDate string, toDate string, orderedBy string) *ResultStat
	DeleteStat()
}

type RepoPostgres struct {
	url             string
	maxIdleConn     int
	maxOpenConn     int
	connMaxLifetime time.Duration
	conn            *gorm.DB
}

func CreateRepoPostgres(conf config.DBConfig) *RepoPostgres {
	return &RepoPostgres{
		url:             conf.URL,
		maxIdleConn:     conf.MaxIdleConn,     // maximum number of connections in the idle connection pool
		maxOpenConn:     conf.MaxOpenConn,     // maximum number of open connections to the database
		connMaxLifetime: conf.ConnMaxLifetime, // maximum amount of time a connection may be reused
		conn:            nil,
	}
}

func (db *RepoPostgres) Open() error {
	var err error
	db.conn, err = gorm.Open(postgres.Open(db.url), &gorm.Config{})
	if err == nil {
		err = db.configureConnectionPools()
	}
	return err
}

func (db RepoPostgres) configureConnectionPools() error {
	sqlDB, err := db.conn.DB()
	if err != nil {
		return err
	}

	sqlDB.SetMaxIdleConns(db.maxIdleConn)
	sqlDB.SetMaxOpenConns(db.maxOpenConn)
	sqlDB.SetConnMaxLifetime(db.connMaxLifetime)

	return nil
}

func (db RepoPostgres) CreateStat(stat *SavedStat) error {
	res := db.conn.Select("Date", "Views", "Clicks", "Cost").Create(stat)
	return res.Error
}

func (db RepoPostgres) GetStat(fromDate string, toDate string, orderBy string) *ResultStat {
	var stat ResultStat

	db.conn.
		Table("users").
		Select("date, views, clicks, cost, cost / clicks AS cpc, cost / views * 1000 AS cpm").
		Where("? <= date AND date <= ?", fromDate, toDate).Order(orderBy).Scan(&stat)

	return &stat
}

func (db RepoPostgres) DeleteStat() {
	db.conn.Exec("TRUNCATE statistics RESTART IDENTITY")
}