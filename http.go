package sdk

type (
	NewClientRequest struct {
		Name     string `json:"name,omitempty"`
		Ip       string `json:"ip,omitempty"`
		PublicIp string `json:"public_ip,omitempty"`
		Group    string `json:"group,omitempty"`
	}

	NewClientResponse struct {
		Id        uint64       `json:"id,omitempty"`
		PublicKey string       `json:"public_key,omitempty"`
		Users     []CreateUser `json:"users,omitempty"`
	}
)
