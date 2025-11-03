package config

import (
	cfgtypes "github.com/Station-Manager/types/config"
	dbtypes "github.com/Station-Manager/types/database"
)

var defaultConfig = cfgtypes.Config{
	DatastoreConfigs: []dbtypes.Config{
		{
			Driver:          "postgres",
			Host:            "localhost",
			Port:            5432,
			Database:        "station_manager",
			User:            "smuser",
			Password:        "",
			SSLMode:         "disable",
			Debug:           false,
			MaxOpenConns:    2,
			MaxIdleConns:    2,
			ConnMaxLifetime: 10, // Minutes
			ConnMaxIdleTime: 5,  // Minutes
			ContextTimeout:  5,  // Seconds
		},
	},
}
