package config

import (
	"github.com/Station-Manager/types"
)

var defaultConfig = types.AppConfig{
	DatastoreConfig: types.DatastoreConfig{
		Driver:                    "postgres",
		Path:                      "db/data.db",
		Options:                   "?_busy_timeout=5000&_journal_mode=WAL&_foreign_keys=on&cache=shared&_txlock=immediate",
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
	},
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
