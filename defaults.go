package config

import (
	"github.com/Station-Manager/types"
)

var defaultConfig = types.AppConfig{
	DatastoreConfigs: []types.DatastoreConfig{
		{
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
		},
	},
}
