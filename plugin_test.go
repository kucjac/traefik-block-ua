package traefik_block_ua_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	tbua "github.com/kucjac/traefik-block-ua"
)

const pluginName = "userAgentBlocker"

type noopHandler struct{}

func (n noopHandler) ServeHTTP(rw http.ResponseWriter, _ *http.Request) {
	rw.WriteHeader(http.StatusTeapot)
}

func TestPlugin(t *testing.T) {
	t.Run("NoUserAgents", func(t *testing.T) {
		cfg := tbua.CreateConfig()
		p, err := tbua.New(context.Background(), noopHandler{}, cfg, pluginName)
		if err != nil {
			t.Fatalf("no error expected, but is: %v", err)
		}

		req := httptest.NewRequest(http.MethodGet, "/foobar", nil)
		rr := httptest.NewRecorder()
		p.ServeHTTP(rr, req)

		if rr.Code != http.StatusTeapot {
			t.Fatalf("expected: %v, is: %v", http.StatusTeapot, rr.Code)
		}
	})

	t.Run("NoUserAgents", func(t *testing.T) {
		cfg := tbua.CreateConfig()
		p, err := tbua.New(context.Background(), noopHandler{}, cfg, pluginName)
		if err != nil {
			t.Fatalf("no error expected, but is: %v", err)
		}

		req := httptest.NewRequest(http.MethodGet, "/foobar", nil)
		rr := httptest.NewRecorder()
		p.ServeHTTP(rr, req)

		if rr.Code != http.StatusTeapot {
			t.Fatalf("expected: %v, is: %v", http.StatusTeapot, rr.Code)
		}
	})

	t.Run("ValidUserAgent", func(t *testing.T) {
		cfg := tbua.CreateConfig()
		cfg.UserAgents = []string{"SpamBot"}
		p, err := tbua.New(context.Background(), noopHandler{}, cfg, pluginName)
		if err != nil {
			t.Fatalf("no error expected, but is: %v", err)
		}

		req := httptest.NewRequest(http.MethodGet, "/foobar", nil)
		req.Header.Set("User-Agent", "Mozilla/5.0 AppleWebKit/537.36 (KHTML, like Gecko; compatible; Googlebot/2.1; +http://www.google.com/bot.html) Chrome/W.X.Y.Z Safari/537.36")
		rr := httptest.NewRecorder()
		p.ServeHTTP(rr, req)

		if rr.Code != http.StatusTeapot {
			t.Fatalf("expected: %v, is: %v", http.StatusTeapot, rr.Code)
		}
	})

	t.Run("ForbiddenUserAgent", func(t *testing.T) {
		cfg := tbua.CreateConfig()
		cfg.UserAgents = []string{"Googlebot"}
		p, err := tbua.New(context.Background(), noopHandler{}, cfg, pluginName)
		if err != nil {
			t.Fatalf("no error expected, but is: %v", err)
		}

		req := httptest.NewRequest(http.MethodGet, "/foobar", nil)
		req.Header.Set("User-Agent", "Mozilla/5.0 AppleWebKit/537.36 (KHTML, like Gecko; compatible; Googlebot/2.1; +http://www.google.com/bot.html) Chrome/W.X.Y.Z Safari/537.36")
		rr := httptest.NewRecorder()
		p.ServeHTTP(rr, req)

		if rr.Code != http.StatusForbidden {
			t.Fatalf("expected: %v, is: %v", http.StatusForbidden, rr.Code)
		}
	})
}
