package config

import (
	"github.com/Station-Manager/types"
)

var postgresConfig = types.DatastoreConfig{
	Driver:                    "postgres",
	Host:                      "localhost",
	Port:                      5432,
	Database:                  "station_manager",
	User:                      "smuser",
	Password:                  "",
	SSLMode:                   "disable",
	Debug:                     false,
	MaxOpenConns:              2,
	MaxIdleConns:              2,
	ConnMaxLifetime:           10, // Minutes
	ConnMaxIdleTime:           5,  // Minutes
	ContextTimeout:            5,  // Seconds
	TransactionContextTimeout: 5,  // Seconds
}

var sqliteConfig = types.DatastoreConfig{
	Driver:                    "sqlite",
	Path:                      "db/data.db",
	Options:                   map[string]string{"mode": "rwc", "_foreign_keys": "on", "_journal_mode": "WAL", "_busy_timeout": "5000"},
	Debug:                     false,
	MaxOpenConns:              2,
	MaxIdleConns:              2,
	ConnMaxLifetime:           10, // Minutes
	ConnMaxIdleTime:           5,  // Minutes
	ContextTimeout:            5,  // Seconds
	TransactionContextTimeout: 5,  // Seconds
}

var defaultConfig = types.AppConfig{
	DatastoreConfig: sqliteConfig,
	LoggingConfig: types.LoggingConfig{
		Level:                  "info",
		WithTimestamp:          true,
		ConsoleLogging:         false,
		FileLogging:            true,
		RelLogFileDir:          "logs",
		SkipFrameCount:         3,
		LogFileMaxSizeMB:       100,
		LogFileMaxAgeDays:      30,
		LogFileMaxBackups:      5,
		ShutdownTimeoutMS:      10000,
		ShutdownTimeoutWarning: true,
	},
}
