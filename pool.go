package ldappool

import (
	"context"
	"crypto/tls"
	"time"

	"github.com/go-ldap/ldap/v3"
)

type LDAPPool interface {
	Add(request *ldap.AddRequest) error
	Del(request *ldap.DelRequest) error
	Modify(request *ldap.ModifyRequest) error
	ModifyDN(request *ldap.ModifyDNRequest) error
	ModifyWithResult(request *ldap.ModifyRequest) (*ldap.ModifyResult, error)
	Compare(dn, attribute, value string) (bool, error)
	PasswordModify(request *ldap.PasswordModifyRequest) (*ldap.PasswordModifyResult, error)
	Search(request *ldap.SearchRequest) (*ldap.SearchResult, error)
	SearchAsync(ctx context.Context, searchRequest *ldap.SearchRequest, bufferSize int) ldap.Response
	SearchWithPaging(searchRequest *ldap.SearchRequest, pagingSize uint32) (*ldap.SearchResult, error)
	DirSync(searchRequest *ldap.SearchRequest, flags, maxAttrCount int64, cookie []byte) (*ldap.SearchResult, error)
	DirSyncAsync(ctx context.Context, searchRequest *ldap.SearchRequest, bufferSize int, flags, maxAttrCount int64, cookie []byte) ldap.Response
	Syncrepl(ctx context.Context, searchRequest *ldap.SearchRequest, bufferSize int, mode ldap.ControlSyncRequestMode, cookie []byte, reloadHint bool) ldap.Response
}

type Pool struct {
	uri          string
	bindDN       string
	bindPassword string
	timeout      time.Duration
	tlsConfig    *tls.Config
}

func New(uri string, options ...Option) (LDAPPool, error) {
	pool := &Pool{
		uri: uri,
	}

	for _, opt := range options {
		opt(pool)
	}

	// Set default timeout
	if pool.timeout == 0 {
		pool.timeout = time.Second * 10
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

func (p *Pool) close(conn *ldap.Conn) {
	_ = conn.Close()
}
