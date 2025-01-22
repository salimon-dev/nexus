package websocket

import (
	"encoding/json"
	"fmt"
)

type BasePayload struct {
	Action      string `json:"action,omitempty"`
	Body        string `json:"body,omitempty"`
	AccessToken string `json:"access_token,omitempty"`
	Entity      string `json:"entity,omitempty"`
}

func Router(data []byte) {
	var payload BasePayload
	err := json.Unmarshal(data, &payload)
	if err != nil {
		fmt.Println(err.Error())
	}
}
