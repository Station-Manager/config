package config

import (
	"github.com/Station-Manager/errors"
	"github.com/Station-Manager/types"
	"sync"
	"sync/atomic"
)

type Service struct {
	WorkingDir    string
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
		// Do some initialization here
		if s.initErr == nil {
			s.isInitialized.Store(true)
		}
	})

	return s.initErr
}

// DatastoreConfig returns the datastore configuration.
func (s *Service) DatastoreConfig() (types.DatastoreConfig, error) {
	const op errors.Op = "config.Service.DatastoreConfig"
	emptryRetVal := types.DatastoreConfig{}
	if s == nil {
		return emptryRetVal, errors.New(op).Msg(errMsgNilService)
	}
	if !s.isInitialized.Load() {
		return emptryRetVal, errors.New(op).Msg(errMsgNotInitialized)
	}

	return s.AppConfig.DatastoreConfig, nil
}

// LoggingConfig returns the logging configuration.
func (s *Service) LoggingConfig() (types.LoggingConfig, error) {
	const op errors.Op = "config.Service.LoggingConfig"
	if s == nil {
		return types.LoggingConfig{}, errors.New(op).Msg(errMsgNilService)
	}
	if !s.isInitialized.Load() {
		return types.LoggingConfig{}, errors.New(op).Msg(errMsgNotInitialized)
	}

	return s.AppConfig.LoggingConfig, nil
}
