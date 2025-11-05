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
}

var initOnce sync.Once

func (s *Service) Initialize() error {
	const op errors.Op = "config.Service.Initialize"
	if s == nil {
		return errors.New(op).Msg(errMsgNilService)
	}

	if s.isInitialized.Load() {
		return nil // Exit gracefully
	}

	initOnce.Do(func() {
		s.isInitialized.Store(true)
	})

	return nil
}

func (s *Service) LoggingConfig() (types.LoggingConfig, error) {
	const op errors.Op = "config.Service.LoggingConfig"
	if s == nil {
		return types.LoggingConfig{}, errors.New(op).Msg(errMsgNilService)
	}

	return s.AppConfig.LoggingConfig, nil
}
