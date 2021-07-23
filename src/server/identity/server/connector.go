package server

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/dexidp/dex/connector"
	"github.com/dexidp/dex/pkg/log"
)

const localhostIdentityServer = "localhost:30658"

var (
	_ connector.CallbackConnector = &placeholder{}
)

// placeholder is a fake Dex connector which redirects the user to a static page with instructions.
// This is necessary because Dex won't start the web server unless a connector is configured.
type placeholderConfig struct {
	LocalhostIssuer bool
}

func (p placeholderConfig) Open(id string, logger log.Logger) (connector.Connector, error) {
	return &placeholder{LocalhostIssuer: p.LocalhostIssuer}, nil
}

type placeholder struct {
	LocalhostIssuer bool
}

func (p *placeholder) LoginURL(s connector.Scopes, callbackURL, state string) (string, error) {
	u, err := url.Parse(callbackURL)
	if err != nil {
		return "", fmt.Errorf("failed to parse callbackURL %q: %v", callbackURL, err)
	}
	if p.LocalhostIssuer {
		u.Host = localhostIdentityServer
	}
	// u.Path = "/static/not-configured.html"
	return "/static/not-configured.html", nil
}

// HandleCallback parses the request and returns the user's identity
func (*placeholder) HandleCallback(s connector.Scopes, r *http.Request) (connector.Identity, error) {
	return connector.Identity{}, fmt.Errorf("no connectors configured")
}

// Refresh updates the identity during a refresh token request.
func (*placeholder) Refresh(ctx context.Context, s connector.Scopes, identity connector.Identity) (connector.Identity, error) {
	return connector.Identity{}, fmt.Errorf("no connectors configured")
}
