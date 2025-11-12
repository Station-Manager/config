package config

import "github.com/Station-Manager/types"

const (
	errMsgNilService     = "Config service is nil."
	errMsgNotInitialized = "Config service is not initialized."
)

const (
	ServiceName    = types.ConfigServiceName
	configFileName = "config.json"
	// EnvSmDefaultDB selects the default datastore driver when generating a new config.json.
	// Accepts: "sqlite" (default), "postgres", and common aliases like "postgresql" or "pg".
	EnvSmDefaultDB = "SM_DEFAULT_DB"
)
