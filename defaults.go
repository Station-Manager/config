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
	MaxOpenConns:              10, // General-purpose default; tune per deployment
	MaxIdleConns:              5,
	ConnMaxLifetime:           30, // Minutes
	ConnMaxIdleTime:           5,  // Minutes
	ContextTimeout:            5,  // Seconds
	TransactionContextTimeout: 15, // Seconds
}

var sqliteConfig = types.DatastoreConfig{
	Driver:                    "sqlite",
	Path:                      "db/data.db",
	Options:                   map[string]string{"mode": "rwc", "_foreign_keys": "on", "_journal_mode": "WAL", "_busy_timeout": "5000"},
	Debug:                     false,
	MaxOpenConns:              4, // Readers benefit; writes still serialized
	MaxIdleConns:              4,
	ConnMaxLifetime:           0,  // No forced recycle for file-backed DB
	ConnMaxIdleTime:           5,  // Minutes
	ContextTimeout:            5,  // Seconds
	TransactionContextTimeout: 10, // Seconds
}

var defaultDesktopConfig = types.AppConfig{
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
		ShutdownTimeoutWarning: false,
	},
	RequiredConfigs: types.RequiredConfigs{},
}

var defaultServerConfig = types.AppConfig{
	DatastoreConfig: postgresConfig,
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
		ShutdownTimeoutWarning: false,
	},
	RequiredConfigs: types.RequiredConfigs{},
	ServerConfig: types.ServerConfig{
		Name:         "Station Manager",
		Port:         3000,
		ReadTimeout:  5,       // Seconds
		WriteTimeout: 10,      // Seconds
		IdleTimeout:  60,      // Seconds
		BodyLimit:    2097152, // 1024 * 1024 * 2 = 2MB
	},
}
