package ldappool

import (
	"context"
	"crypto/tls"
	"time"

	"github.com/go-ldap/ldap/v3"
)

type Pool struct {
	uri          string
	bindDN       string
	bindPassword string
	timeout      time.Duration
	tc           *tls.Config

	maxConnections int
	connections    chan *ldap.Conn
}

func New(uri string, options ...Option) (*Pool, error) {
	pool := &Pool{
		uri: uri,
	}

	for _, opt := range options {
		opt(pool)
	}

	// Set default max connections
	if pool.maxConnections == 0 {
		pool.maxConnections = 5
	}

	pool.connections = make(chan *ldap.Conn, pool.maxConnections)
	for i := 0; i < pool.maxConnections; i++ {
		conn, err := pool.conn()
		if err != nil {
			return nil, err
		}

		pool.connections <- conn
	}

	return pool, nil
}

func (p *Pool) conn() (conn *ldap.Conn, err error) {
	conn, err = ldap.DialURL(p.uri)
	if err != nil {
		return
	}

	if p.tc != nil {
		if err = conn.StartTLS(p.tc); err != nil {
			return
		}
	}

	if p.timeout > 0 {
		conn.SetTimeout(p.timeout)
	}

	if len(p.bindDN) > 0 {
		err = conn.Bind(p.bindDN, p.bindPassword)
	}

	return conn, err
}

func (p *Pool) get(ctx context.Context) (conn *ldap.Conn, err error) {
	select {
	case <-ctx.Done():
		err = ctx.Err()
	case conn = <-p.connections:
		if conn.IsClosing() {
			conn, err = p.conn()
			if err != nil {
				return
			}
		}
	}

	return
}

func (p *Pool) release(conn *ldap.Conn) {
	p.connections <- conn
}

func (p *Pool) Close() {
	for i := 0; i < p.maxConnections; i++ {
		conn := <-p.connections
		_ = conn.Close()
	}
}
