package proxy

import (
	"context"
	"encoding/json"

	"github.com/centrifugal/centrifuge"
)

// ConnectRequest ...
type ConnectRequest struct {
	ClientID  string
	Transport centrifuge.TransportInfo
	Data      centrifuge.Raw
}

// ConnectCredentials ...
type ConnectCredentials struct {
	UserID     string          `json:"user"`
	ExpireAt   int64           `json:"exp"`
	Info       json.RawMessage `json:"info"`
	Base64Info string          `json:"b64info"`
}

// ConnectReply ...
type ConnectReply struct {
	Result     *ConnectCredentials    `json:"result"`
	Error      *centrifuge.Error      `json:"error"`
	Disconnect *centrifuge.Disconnect `json:"disconnect"`
}

// ConnectProxy allows to proxy connect requests to application backend to
// authenticate client connection.
type ConnectProxy interface {
	ProxyConnect(context.Context, ConnectRequest) (*ConnectReply, error)
	// Protocol for metrics and logging.
	Protocol() string
}
