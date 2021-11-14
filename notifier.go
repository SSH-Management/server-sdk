package sdk

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type ClientNotify struct {
	addr   string
	client *http.Client
}

func New(addr string) ClientNotify {
	client := &http.Client{
		Timeout: 100 * time.Millisecond,
	}

	return ClientNotify{
		addr:   addr,
		client: client,
	}
}

func (c ClientNotify) NotifyMasterForNewClient(ctx context.Context, notify NewClientRequest) (uint64, string, []CreateUser, error) {
	body, err := json.Marshal(notify)

	if err != nil {
		return 0, "", nil, err
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/client/new", c.addr), bytes.NewReader(body))

	if err != nil {
		return 0, "", nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	res, err := c.client.Do(req)

	if err != nil {
		return 0, "", nil, err
	}

	if res.StatusCode != http.StatusCreated {
		return 0, "", nil, ErrFailedToCreateClient
	}

	var newClientResponse NewClientResponse

	resBody := res.Body
	defer resBody.Close()

	if err := json.NewDecoder(resBody).Decode(&newClientResponse); err != nil {
		return 0, "", nil, err
	}

	return newClientResponse.Id, newClientResponse.PublicKey, newClientResponse.Users, nil
}
