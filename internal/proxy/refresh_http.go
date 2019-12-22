package proxy

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/centrifugal/centrifugo/internal/middleware"
)

// RefreshRequestHTTP ...
type RefreshRequestHTTP struct {
	ClientID  string `json:"client"`
	UserID    string `json:"user"`
	Transport string `json:"transport"`
	Protocol  string `json:"protocol"`
	Encoding  string `json:"encoding"`
}

// HTTPRefreshProxy ...
type HTTPRefreshProxy struct {
	httpCaller HTTPCaller
}

// NewHTTPRefreshProxy ...
func NewHTTPRefreshProxy(endpoint string, httpClient *http.Client) *HTTPRefreshProxy {
	return &HTTPRefreshProxy{
		httpCaller: NewHTTPCaller(endpoint, httpClient),
	}
}

// ProxyRefresh proxies refresh to application backend.
func (p *HTTPRefreshProxy) ProxyRefresh(ctx context.Context, req RefreshRequest) (*RefreshReply, error) {
	httpRequest := middleware.HeadersFromContext(ctx)

	refreshHTTPReq := RefreshRequestHTTP{
		ClientID:  req.ClientID,
		UserID:    req.UserID,
		Transport: req.Transport.Name(),
		Protocol:  string(req.Transport.Protocol()),
		Encoding:  string(req.Transport.Encoding()),
	}

	data, err := json.Marshal(refreshHTTPReq)
	if err != nil {
		return nil, err
	}

	respData, err := p.httpCaller.CallHTTP(ctx, getProxyHeader(httpRequest), data)
	if err != nil {
		return nil, err
	}

	var res RefreshReply
	err = json.Unmarshal(respData, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// Protocol ...
func (p *HTTPRefreshProxy) Protocol() string {
	return "http"
}
