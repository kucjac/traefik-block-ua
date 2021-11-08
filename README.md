# traefik-block-ua

[![Build Status](https://github.com/kucjac/traefik-block-ua/actions/workflows/ci.yml/badge.svg)](https://github.com/kucjac/traefik-block-ua/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/kucjac/traefik-block-ua)](https://goreportcard.com/report/github.com/kucjac/traefik-block-ua)
[![Latest GitHub release](https://img.shields.io/github/v/release/kucjac/traefik-block-ua?sort=semver)](https://github.com/kucjac/traefik-block-ua/releases/latest)
[![License](https://img.shields.io/badge/license-Apache%202.0-brightgreen.svg)](LICENSE)

*traefik-block-ua is a traefik plugin to whitelist requests based on the user agents*

## Configuration

### Static

#### Local

```yaml
experimental:
  localPlugins:
    geoblock:
      moduleName: github.com/kucjac/traefik-block-ua
```

#### Pilot

```yaml
pilot:
  token: "xxxxxxxxx"

experimental:
  plugins:
    userAgentBlocker:
      moduleName: github.com/kucjac/traefik-block-ua
      version: v0.1.0
```

### Dynamic

```yaml
http:
  middlewares:
    userAgentBlocker:
      plugin:
        userAgentBlocker:
          userAgents:
            - "CustomCrawlerName"
```