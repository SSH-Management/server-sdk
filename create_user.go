package server_notify

type (
	CreateUser interface {
		GetUser() User
		GetPublicKey() string
	}

	User interface {
		GetName() string
		GetSurname() string
		GetUsername() string
		GetEmail() string
		GetPassword() string
		GetShell() string
		GetRole() string
		GetGroups() []string
		GetSystemGroups() []string
	}
)
