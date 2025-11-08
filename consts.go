package config

import "github.com/Station-Manager/types"

const (
	errMsgNilService       = "Config service is nil."
	errMsgNotInitialized   = "Config service is not initialized."
	errMsgEmptyServiceName = "Service name parameter cannot be empty."
)

const (
	ServiceName = types.ConfigServiceName
)
