package types

type AuthEvent struct {
	Action      string `json:"action"`
	AccessToken string `json:"access_token"`
}

type MessageEvent struct {
	Action string `json:"action"`
	Body   string `json:"body"`
}

type ConnectEvent struct {
	Action string `json:"action"`
	Entity string `json:"entity"`
}
