package server_notify

import "errors"

var (
	ErrFailedToCreateClient = errors.New("failed to register client on master SSH server")
)
