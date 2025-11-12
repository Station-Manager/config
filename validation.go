package config

import (
	"fmt"
	"github.com/Station-Manager/errors"
	"github.com/Station-Manager/types"
)

// validateAppConfig performs minimal validation to avoid obviously bad configs while remaining permissive.
func validateAppConfig(cfg *types.AppConfig) error {
	const op errors.Op = "config.validateAppConfig"
	if cfg == nil {
		return errors.New(op).Msg("AppConfig is nil")
	}

	db := cfg.DatastoreConfig
	switch db.Driver {
	case types.SqliteServiceName:
		if db.Path == "" {
			return errors.New(op).Msg("sqlite path is required")
		}
	case types.PostgresServiceName:
		// Be permissive here; database service will perform full validation later.
		// Just accept the driver selection.
		_ = db
	default:
		return errors.New(op).Msg(fmt.Sprintf("unsupported driver: %s", db.Driver))
	}

	// Minimal check for logging config
	if cfg.LoggingConfig.Level == "" {
		return errors.New(op).Msg("logging level must be set")
	}
	return nil
}
