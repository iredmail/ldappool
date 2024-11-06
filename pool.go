package ldappool

import (
	"context"
	"crypto/tls"
	"fmt"
	"time"

	"github.com/go-ldap/ldap/v3"
)

const defaultMaxConnections = 10

type Pool struct {
	uri          string
	bindDN       string
	bindPassword string
	timeout      time.Duration
	tlsConfig    *tls.Config

	maxConnections int
	connections    chan *ldap.Conn
}

func New(uri string, options ...Option) (ldap.Client, error) {
	pool := &Pool{
		uri: uri,
	}

	for _, opt := range options {
		opt(pool)
	}

	// Set default max connections
	if pool.maxConnections == 0 {
		pool.maxConnections = defaultMaxConnections
	}

	// Set default timeout
	if pool.timeout == 0 {
		pool.timeout = time.Second * 10
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

	if p.tlsConfig != nil {
		if err = conn.StartTLS(p.tlsConfig); err != nil {
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

func (p *Pool) get() (conn *ldap.Conn, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), p.timeout)
	defer cancel()

	select {
	case <-ctx.Done():
		err = fmt.Errorf("failed in acquiring ldap connection from pool: %v", ctx.Err())
	case conn = <-p.connections:
		if conn != nil && conn.IsClosing() {
			var _conn *ldap.Conn
			_conn, err = p.conn() // Recreate ldap connection.
			if err == nil {
				conn = _conn
			}
		}
	}

	return
}

func (p *Pool) put(conn *ldap.Conn) {
	p.connections <- conn
}
