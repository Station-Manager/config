package config

import (
	"github.com/Station-Manager/errors"
	"github.com/Station-Manager/utils"
	"github.com/goccy/go-json"
	"os"
	"path/filepath"
	"strings"
)

func (s *Service) loadConfigFile() error {
	const op errors.Op = "config.Service.loadConfigFile"
	if s == nil {
		return errors.New(op).Msg(errMsgNilService)
	}

	filePath := filepath.Join(s.WorkingDir, configFileName)
	exists, err := utils.PathExists(filePath)
	if err != nil {
		return errors.New(op).Err(err)
	}

	if !exists {
		if err = s.generateDefaultConfig(); err != nil {
			return errors.New(op).Err(err)
		}
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return errors.New(op).Err(err)
	}

	if err = json.Unmarshal(data, &s.AppConfig); err != nil {
		return errors.New(op).Err(err)
	}

	return nil
}

func (s *Service) generateDefaultConfig() error {
	const op errors.Op = "config.Service.generateDefaultConfig"
	if s == nil {
		return errors.New(op).Msg(errMsgNilService)
	}

	// Decide which datastore config to embed based on env
	selected := defaultConfig // start with sqlite default
	if dbSel := strings.ToLower(strings.TrimSpace(os.Getenv(EnvSmDefaultDB))); dbSel != "" {
		switch dbSel {
		case "postgres", "postgresql", "pg":
			selected.DatastoreConfig = postgresConfig
		case "sqlite", "sqlite3":
			// already sqlite; explicitly set for clarity
			selected.DatastoreConfig = sqliteConfig
		default:
			// Unknown selector: leave default (sqlite). Could log later if logging available.
		}
	}

	// Pretty-print selected configuration for readability
	data, err := json.MarshalIndent(selected, "", "  ")
	if err != nil {
		return errors.New(op).Err(err)
	}

	return writeDataToFile(data, filepath.Join(s.WorkingDir, configFileName))
}
