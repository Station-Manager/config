package config

import (
	"github.com/Station-Manager/errors"
	"sync/atomic"
)

type Service struct {
	isInitialized atomic.Bool
}

func (s *Service) Initialize() error {
	const op errors.Op = "config.Service.Initialize"
	if s == nil {
		return errors.New(op).Msg(errMsgNilService)
	}

	if s.isInitialized.Load() {
		return nil // Exit gracefully
	}

	s.isInitialized.Store(true)

	return nil
}
