package traefik_block_ua

import (
	"context"
	"fmt"
	"log"
	"net/http"

	ua "github.com/mileusna/useragent"
)

// Config defines the plugin dynamic configuration.
type Config struct {
	UserAgents []string
}

// CreateConfig creates a new config.
func CreateConfig() *Config {
	return &Config{}
}

// Plugin is the traefik plugin implementation.
type Plugin struct {
	next        http.Handler
	name        string
	knownAgents map[string]struct{}
}

// New creates a new plugin handler.
func New(_ context.Context, next http.Handler, cfg *Config, name string) (http.Handler, error) {
	if next == nil {
		return nil, fmt.Errorf("no next handler provided")
	}
	if cfg == nil {
		return nil, fmt.Errorf("no config provided")
	}

	knownAgents := map[string]struct{}{}
	for _, ka := range cfg.UserAgents {
		knownAgents[ka] = struct{}{}
	}

	return &Plugin{
		next:        next,
		name:        name,
		knownAgents: knownAgents,
	}, nil
}

// ServeHTTP implements http.Handler interface.
func (p *Plugin) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	agent := ua.Parse(req.Header.Get("User-Agent"))

	if _, ok := p.knownAgents[agent.Name]; ok {
		log.Printf("%s: %s - access denied - blocked user agent: %s", p.name, req.URL.String(), agent.Name)
		rw.WriteHeader(http.StatusForbidden)
		return
	}
	p.next.ServeHTTP(rw, req)
}
