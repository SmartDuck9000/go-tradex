package config

import "time"

type DBConfig struct {
	URL             string
	MaxIdleConn     int
	MaxOpenConn     int
	ConnMaxLifetime time.Duration
}

type ServiceConfig struct {
	DB   DBConfig
	Host string
	Port string
}
