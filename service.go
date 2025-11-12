package config

import (
	"github.com/Station-Manager/errors"
	"github.com/Station-Manager/types"
	"github.com/Station-Manager/utils"
	"sync"
	"sync/atomic"
)

type Service struct {
	WorkingDir    string `di.inject:"workingdir"`
	AppConfig     types.AppConfig
	isInitialized atomic.Bool
	initOnce      sync.Once
	initErr       error
}

// Initialize initializes the config service.
func (s *Service) Initialize() error {
	const op errors.Op = "config.Service.Initialize"
	if s == nil {
		return errors.New(op).Msg(errMsgNilService)
	}

	if s.isInitialized.Load() {
		return nil // Exit gracefully
	}

	s.initOnce.Do(func() {
		var err error
		if s.WorkingDir, err = utils.WorkingDir(s.WorkingDir); err != nil {
			s.initErr = errors.New(op).Err(err).Msg(errMsgWorkingDir)
			return
		}

		if err = s.loadConfigFile(); err != nil {
			s.initErr = errors.New(op).Err(err)
			return
		}

		// Early validation of loaded configuration
		if err = validateAppConfig(&s.AppConfig); err != nil {
			s.initErr = errors.New(op).Err(err)
			return
		}

		s.isInitialized.Store(true)
	})

	return s.initErr
}

// DatastoreConfig returns the datastore configuration.
func (s *Service) DatastoreConfig() (types.DatastoreConfig, error) {
	const op errors.Op = "config.Service.DatastoreConfig"
	emptyRetVal := types.DatastoreConfig{}
	if s == nil {
		return emptyRetVal, errors.New(op).Msg(errMsgNilService)
	}
	if !s.isInitialized.Load() {
		return emptyRetVal, errors.New(op).Msg(errMsgNotInitialized)
	}

	return s.AppConfig.DatastoreConfig, nil
}

// LoggingConfig returns the logging configuration.
func (s *Service) LoggingConfig() (types.LoggingConfig, error) {
	const op errors.Op = "config.Service.LoggingConfig"
	emptyRetVal := types.LoggingConfig{}
	if s == nil {
		return emptyRetVal, errors.New(op).Msg(errMsgNilService)
	}
	if !s.isInitialized.Load() {
		return emptyRetVal, errors.New(op).Msg(errMsgNotInitialized)
	}

	return s.AppConfig.LoggingConfig, nil
}
