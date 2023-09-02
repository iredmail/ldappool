package ldappool

import (
	"context"
	"crypto/tls"
	"time"

	"github.com/go-ldap/ldap/v3"
)

func (p *Pool) Start(ctx context.Context) error {
	conn, err := p.get(ctx)
	if err == nil {
		defer p.release(conn)
		conn.Start()
	}

	return err
}

func (p *Pool) StartTLS(ctx context.Context, config *tls.Config) error {
	conn, err := p.get(ctx)
	if err != nil {
		return err
	}

	defer p.release(conn)

	return conn.StartTLS(config)
}

func (p *Pool) SetTimeout(ctx context.Context, duration time.Duration) error {
	conn, err := p.get(ctx)
	if err == nil {
		defer p.release(conn)
		conn.SetTimeout(duration)
	}

	return err
}

func (p *Pool) TLSConnectionState(ctx context.Context) (tc tls.ConnectionState, v bool) {
	conn, err := p.get(ctx)
	if err != nil {
		return
	}

	defer p.release(conn)

	return conn.TLSConnectionState()
}

func (p *Pool) Bind(ctx context.Context, username, password string) error {
	conn, err := p.get(ctx)
	if err != nil {
		return err
	}

	defer p.release(conn)

	return conn.Bind(username, password)
}

func (p *Pool) UnauthenticatedBind(ctx context.Context, username string) error {
	conn, err := p.get(ctx)
	if err != nil {
		return err
	}

	defer p.release(conn)

	return conn.UnauthenticatedBind(username)
}

func (p *Pool) SimpleBind(ctx context.Context, request *ldap.SimpleBindRequest) (*ldap.SimpleBindResult, error) {
	conn, err := p.get(ctx)
	if err != nil {
		return nil, err
	}

	defer p.release(conn)

	return conn.SimpleBind(request)
}

func (p *Pool) Add(ctx context.Context, request *ldap.AddRequest) error {
	conn, err := p.get(ctx)
	if err != nil {
		return err
	}

	defer p.release(conn)

	return conn.Add(request)
}

func (p *Pool) Del(ctx context.Context, request *ldap.DelRequest) error {
	conn, err := p.get(ctx)
	if err != nil {
		return err
	}

	defer p.release(conn)

	return conn.Del(request)
}

func (p *Pool) Modify(ctx context.Context, request *ldap.ModifyRequest) error {
	conn, err := p.get(ctx)
	if err != nil {
		return err
	}

	defer p.release(conn)

	return conn.Modify(request)
}

func (p *Pool) ModifyDN(ctx context.Context, request *ldap.ModifyDNRequest) error {
	conn, err := p.get(ctx)
	if err != nil {
		return err
	}

	defer p.release(conn)

	return conn.ModifyDN(request)
}

func (p *Pool) Compare(ctx context.Context, dn, attribute, value string) (bool, error) {
	conn, err := p.get(ctx)
	if err != nil {
		return false, err
	}

	defer p.release(conn)

	return conn.Compare(dn, attribute, value)
}

func (p *Pool) PasswordModify(ctx context.Context, request *ldap.PasswordModifyRequest) (*ldap.PasswordModifyResult, error) {
	conn, err := p.get(ctx)
	if err != nil {
		return nil, err
	}

	defer p.release(conn)

	return conn.PasswordModify(request)
}

func (p *Pool) Search(ctx context.Context, request *ldap.SearchRequest) (*ldap.SearchResult, error) {
	conn, err := p.get(ctx)
	if err != nil {
		return nil, err
	}

	defer p.release(conn)

	return conn.Search(request)
}

func (p *Pool) SearchWithPaging(ctx context.Context, searchRequest *ldap.SearchRequest, pagingSize uint32) (*ldap.SearchResult, error) {
	conn, err := p.get(ctx)
	if err != nil {
		return nil, err
	}

	defer p.release(conn)

	return conn.SearchWithPaging(searchRequest, pagingSize)
}
