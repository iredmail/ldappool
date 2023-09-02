package ldappool

import (
	"crypto/tls"
	"time"

	"github.com/go-ldap/ldap/v3"
)

func (p *Pool) Start() {
	conn, err := p.get()
	if err != nil {
		return
	}

	defer p.release(conn)
	conn.Start()
}

func (p *Pool) StartTLS(config *tls.Config) error {
	conn, err := p.get()
	if err != nil {
		return err
	}

	defer p.release(conn)

	return conn.StartTLS(config)
}

func (p *Pool) Close() error {
	for i := 0; i < p.maxConnections; i++ {
		conn := <-p.connections
		_ = conn.Close()
	}

	return nil
}

func (p *Pool) GetLastError() error {
	conn, err := p.get()
	if err != nil {
		return err
	}

	defer p.release(conn)

	return conn.GetLastError()
}

func (p *Pool) IsClosing() bool {
	conn, err := p.get()
	if err != nil {
		return true
	}

	defer p.release(conn)

	return conn.IsClosing()
}

func (p *Pool) SetTimeout(duration time.Duration) {
	conn, err := p.get()
	if err != nil {
		return
	}

	defer p.release(conn)
	conn.SetTimeout(duration)
}

func (p *Pool) TLSConnectionState() (cs tls.ConnectionState, found bool) {
	conn, err := p.get()
	if err != nil {
		return
	}

	defer p.release(conn)

	return conn.TLSConnectionState()
}

func (p *Pool) Bind(username, password string) error {
	conn, err := p.get()
	if err != nil {
		return err
	}

	defer p.release(conn)

	return conn.Bind(username, password)
}

func (p *Pool) UnauthenticatedBind(username string) error {
	conn, err := p.get()
	if err != nil {
		return err
	}

	defer p.release(conn)

	return conn.UnauthenticatedBind(username)
}

func (p *Pool) SimpleBind(request *ldap.SimpleBindRequest) (*ldap.SimpleBindResult, error) {
	conn, err := p.get()
	if err != nil {
		return nil, err
	}

	defer p.release(conn)

	return conn.SimpleBind(request)
}

func (p *Pool) ExternalBind() error {
	conn, err := p.get()
	if err != nil {
		return err
	}

	defer p.release(conn)

	return conn.ExternalBind()
}

func (p *Pool) NTLMUnauthenticatedBind(domain, username string) error {
	conn, err := p.get()
	if err != nil {
		return err
	}

	defer p.release(conn)

	return conn.NTLMUnauthenticatedBind(domain, username)
}

func (p *Pool) Unbind() error {
	conn, err := p.get()
	if err != nil {
		return err
	}

	defer p.release(conn)

	return conn.Unbind()
}

func (p *Pool) Add(request *ldap.AddRequest) error {
	conn, err := p.get()
	if err != nil {
		return err
	}

	defer p.release(conn)

	return conn.Add(request)
}

func (p *Pool) Del(request *ldap.DelRequest) error {
	conn, err := p.get()
	if err != nil {
		return err
	}

	defer p.release(conn)

	return conn.Del(request)
}

func (p *Pool) Modify(request *ldap.ModifyRequest) error {
	conn, err := p.get()
	if err != nil {
		return err
	}

	defer p.release(conn)

	return conn.Modify(request)
}

func (p *Pool) ModifyDN(request *ldap.ModifyDNRequest) error {
	conn, err := p.get()
	if err != nil {
		return err
	}

	defer p.release(conn)

	return conn.ModifyDN(request)
}

func (p *Pool) ModifyWithResult(request *ldap.ModifyRequest) (*ldap.ModifyResult, error) {
	conn, err := p.get()
	if err != nil {
		return nil, err
	}

	defer p.release(conn)

	return conn.ModifyWithResult(request)
}

func (p *Pool) Compare(dn, attribute, value string) (bool, error) {
	conn, err := p.get()
	if err != nil {
		return false, err
	}

	defer p.release(conn)

	return conn.Compare(dn, attribute, value)
}

func (p *Pool) PasswordModify(request *ldap.PasswordModifyRequest) (*ldap.PasswordModifyResult, error) {
	conn, err := p.get()
	if err != nil {
		return nil, err
	}

	defer p.release(conn)

	return conn.PasswordModify(request)
}

func (p *Pool) Search(request *ldap.SearchRequest) (*ldap.SearchResult, error) {
	conn, err := p.get()
	if err != nil {
		return nil, err
	}

	defer p.release(conn)

	return conn.Search(request)
}

func (p *Pool) SearchWithPaging(searchRequest *ldap.SearchRequest, pagingSize uint32) (*ldap.SearchResult, error) {
	conn, err := p.get()
	if err != nil {
		return nil, err
	}

	defer p.release(conn)

	return conn.SearchWithPaging(searchRequest, pagingSize)
}

func (p *Pool) DirSync(searchRequest *ldap.SearchRequest, flags, maxAttrCount int64, cookie []byte) (*ldap.SearchResult, error) {
	conn, err := p.get()
	if err != nil {
		return nil, err
	}

	defer p.release(conn)

	return conn.DirSync(searchRequest, flags, maxAttrCount, cookie)
}
