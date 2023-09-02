package ldappool

import (
	"crypto/tls"
	"time"
)

type Option func(p *Pool)

func WithMaxConnections(max int) Option {
	return func(p *Pool) {
		p.maxConnections = max
	}
}

func WithTLSConfig(tc *tls.Config) Option {
	return func(p *Pool) {
		p.tc = tc
	}
}

func WithBindCredentials(dn, password string) Option {
	return func(p *Pool) {
		p.bindDN = dn
		p.bindPassword = password
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(p *Pool) {
		p.timeout = timeout
	}
}
