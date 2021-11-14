package sdk

import "context"

type (
	Interface interface {
		NotifyMasterForNewClient(context.Context, NewClientRequest) (uint64, string, []CreateUser, error)
	}
)
