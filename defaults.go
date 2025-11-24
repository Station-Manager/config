package config

import (
	"github.com/Station-Manager/cat/enums/ans"
	"github.com/Station-Manager/cat/enums/cmd"
	"github.com/Station-Manager/types"
	"time"
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
	RequiredConfigs: defaultRequiredConfigs,
	RigConfigs:      defaultRigConfigs,
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

var defaultRequiredConfigs = types.RequiredConfigs{
	DefaultRigID: 1,
}

var defaultRigConfigs = []types.RigConfig{
	{
		ID:    1,
		Name:  "FTdx10",
		Model: "Yaesu FTdx10",
		SerialConfig: types.SerialConfig{
			PortName:       "/dev/ttyUSB0",
			BaudRate:       38400,
			DataBits:       8,
			Parity:         0,
			StopBits:       0,
			ReadTimeoutms:  150, // Used for context timeout as well as setting on the port
			WriteTimeoutms: 300,
			RTS:            true,
			DTR:            true,
			LineDelimiter:  ';',
		},
		CatCommands: []types.CatCommand{
			{
				Name: cmd.Init.String(),
				Cmd:  "AI1;ID;",
			},
			{
				Name: cmd.Read.String(),
				Cmd:  "FA;FB;ST;VS;MD0;MD1;PC;",
			},
			{
				Name: cmd.PlayBack.String(),
				Cmd:  "PB0%s;",
			},
		},
		CatStates: []types.CatState{
			{
				Prefix: "ID",
				Markers: []types.Marker{
					{
						Tag:    ans.Identity.String(),
						Index:  2,
						Length: 4,
						ValueMappings: []types.ValueMapping{
							{
								Key:   "0761",
								Value: "FTdx10",
							},
						},
					},
				},
			},
			{
				Prefix: "FA",
				Markers: []types.Marker{
					{
						Tag:    ans.VfoAFreq.String(),
						Index:  2,
						Length: 9,
					},
				},
			},
			{
				Prefix: "FB",
				Markers: []types.Marker{
					{
						Tag:    ans.VfoBFreq.String(),
						Index:  2,
						Length: 9,
					},
				},
			},
			{
				Prefix: "ST",
				Markers: []types.Marker{
					{
						Tag:    ans.Split.String(),
						Index:  2,
						Length: 1,
						ValueMappings: []types.ValueMapping{
							{
								Key:   "0",
								Value: "OFF",
							},
							{
								Key:   "1",
								Value: "ON",
							},
							{
								Key:   "2",
								Value: "ON+",
							},
						},
					},
				},
			},
			{
				Prefix: "VS",
				Markers: []types.Marker{
					{
						Tag:    ans.Select.String(),
						Index:  2,
						Length: 1,
						ValueMappings: []types.ValueMapping{
							{
								Key:   "0",
								Value: "VFO-A",
							},
							{
								Key:   "1",
								Value: "VFO-B",
							},
						},
					},
				},
			},
			{
				Prefix: "MD0",
				Markers: []types.Marker{
					{
						Tag:    ans.MainMode.String(),
						Index:  3,
						Length: 1,
						ValueMappings: []types.ValueMapping{
							{
								Key:   "1",
								Value: "LSB",
							},
							{
								Key:   "2",
								Value: "USB",
							},
							{
								Key:   "3",
								Value: "CW-U",
							},
							{
								Key:   "4",
								Value: "FM",
							},
							{
								Key:   "5",
								Value: "AM",
							},
							{
								Key:   "6",
								Value: "RTTY-L",
							},
							{
								Key:   "7",
								Value: "CW-L",
							},
							{
								Key:   "8",
								Value: "DATA-L",
							},
							{
								Key:   "9",
								Value: "RTTY-U",
							},
							{
								Key:   "A",
								Value: "DATA-FM",
							},
							{
								Key:   "B",
								Value: "FM-N",
							},
							{
								Key:   "C",
								Value: "DATA-U",
							},
							{
								Key:   "D",
								Value: "AM-N",
							},
							{
								Key:   "E",
								Value: "PSK",
							},
							{
								Key:   "F",
								Value: "DATA-FM-N",
							},
						},
					},
				},
			},
			{
				Prefix: "MD1",
				Markers: []types.Marker{
					{
						Tag:    ans.SubMode.String(),
						Index:  3,
						Length: 1,
						ValueMappings: []types.ValueMapping{
							{
								Key:   "1",
								Value: "LSB",
							},
							{
								Key:   "2",
								Value: "USB",
							},
							{
								Key:   "3",
								Value: "CW-U",
							},
							{
								Key:   "4",
								Value: "FM",
							},
							{
								Key:   "5",
								Value: "AM",
							},
							{
								Key:   "6",
								Value: "RTTY-L",
							},
							{
								Key:   "7",
								Value: "CW-L",
							},
							{
								Key:   "8",
								Value: "DATA-L",
							},
							{
								Key:   "9",
								Value: "RTTY-U",
							},
							{
								Key:   "A",
								Value: "DATA-FM",
							},
							{
								Key:   "B",
								Value: "FM-N",
							},
							{
								Key:   "C",
								Value: "DATA-U",
							},
							{
								Key:   "D",
								Value: "AM-N",
							},
							{
								Key:   "E",
								Value: "PSK",
							},
							{
								Key:   "F",
								Value: "DATA-FM-N",
							},
						},
					},
				},
			},
			{
				Prefix: "PC",
				Markers: []types.Marker{
					{
						Tag:    ans.TxPwr.String(),
						Index:  2,
						Length: 3,
					},
				},
			},
		},
		CatConfig: types.CatConfig{
			RateLimiterInterval:      20 * time.Millisecond,
			ReadBufferSize:           1024,
			CmdChannelSize:           1000,            // Max allowed by validation: 1000
			ReplyChannelSize:         1000,            // Max allowed by validation: 1000
			StatusChannelSize:        1,               // Setting this to '0' will cause the channel to be unbuffered!
			CommandTimeout:           5 * time.Second, // Timeout for command responses (0 = no timeout)
			RateLimiterCmdsPerSecond: 10,              // Max commands per second (min 1, max 20 by validation)
		},
	},
}
