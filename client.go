package ldappool

import (
	"context"
	"crypto/tls"
	"time"

	"github.com/go-ldap/ldap/v3"
)

func (p *Pool) Start() {
	conn, err := p.get()
	defer p.put(conn)
	if err != nil {
		return
	}

	conn.Start()
}

func (p *Pool) StartTLS(config *tls.Config) error {
	conn, err := p.get()
	defer p.put(conn)
	if err != nil {
		return err
	}

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
	defer p.put(conn)
	if err != nil {
		return err
	}

	return conn.GetLastError()
}

func (p *Pool) IsClosing() bool {
	conn, err := p.get()
	defer p.put(conn)
	if err != nil {
		return true
	}

	return conn.IsClosing()
}

func (p *Pool) SetTimeout(duration time.Duration) {
	conn, err := p.get()
	defer p.put(conn)
	if err != nil {
		return
	}

	conn.SetTimeout(duration)
}

func (p *Pool) TLSConnectionState() (cs tls.ConnectionState, found bool) {
	conn, err := p.get()
	defer p.put(conn)
	if err != nil {
		return
	}

	return conn.TLSConnectionState()
}

func (p *Pool) Bind(username, password string) error {
	conn, err := p.get()
	defer p.put(conn)
	if err != nil {
		return err
	}

	return conn.Bind(username, password)
}

func (p *Pool) UnauthenticatedBind(username string) error {
	conn, err := p.get()
	defer p.put(conn)
	if err != nil {
		return err
	}

	return conn.UnauthenticatedBind(username)
}

func (p *Pool) SimpleBind(request *ldap.SimpleBindRequest) (*ldap.SimpleBindResult, error) {
	conn, err := p.get()
	defer p.put(conn)
	if err != nil {
		return nil, err
	}

	return conn.SimpleBind(request)
}

func (p *Pool) ExternalBind() error {
	conn, err := p.get()
	defer p.put(conn)
	if err != nil {
		return err
	}

	return conn.ExternalBind()
}

func (p *Pool) NTLMUnauthenticatedBind(domain, username string) error {
	conn, err := p.get()
	defer p.put(conn)
	if err != nil {
		return err
	}

	return conn.NTLMUnauthenticatedBind(domain, username)
}

func (p *Pool) Unbind() error {
	conn, err := p.get()
	defer p.put(conn)
	if err != nil {
		return err
	}

	return conn.Unbind()
}

func (p *Pool) Add(request *ldap.AddRequest) error {
	conn, err := p.get()
	defer p.put(conn)
	if err != nil {
		return err
	}

	return conn.Add(request)
}

func (p *Pool) Del(request *ldap.DelRequest) error {
	conn, err := p.get()
	defer p.put(conn)
	if err != nil {
		return err
	}

	return conn.Del(request)
}

func (p *Pool) Modify(request *ldap.ModifyRequest) error {
	conn, err := p.get()
	defer p.put(conn)
	if err != nil {
		return err
	}

	return conn.Modify(request)
}

func (p *Pool) ModifyDN(request *ldap.ModifyDNRequest) error {
	conn, err := p.get()
	defer p.put(conn)
	if err != nil {
		return err
	}

	return conn.ModifyDN(request)
}

func (p *Pool) ModifyWithResult(request *ldap.ModifyRequest) (*ldap.ModifyResult, error) {
	conn, err := p.get()
	defer p.put(conn)
	if err != nil {
		return nil, err
	}

	return conn.ModifyWithResult(request)
}

func (p *Pool) Compare(dn, attribute, value string) (bool, error) {
	conn, err := p.get()
	defer p.put(conn)
	if err != nil {
		return false, err
	}

	return conn.Compare(dn, attribute, value)
}

func (p *Pool) PasswordModify(request *ldap.PasswordModifyRequest) (*ldap.PasswordModifyResult, error) {
	conn, err := p.get()
	defer p.put(conn)
	if err != nil {
		return nil, err
	}

	return conn.PasswordModify(request)
}

func (p *Pool) Search(request *ldap.SearchRequest) (*ldap.SearchResult, error) {
	conn, err := p.get()
	defer p.put(conn)
	if err != nil {
		return nil, err
	}

	return conn.Search(request)
}

func (p *Pool) SearchAsync(ctx context.Context, searchRequest *ldap.SearchRequest, bufferSize int) ldap.Response {
	conn, err := p.get()
	defer p.put(conn)
	if err != nil {
		return ResponseError(err)
	}

	return conn.SearchAsync(ctx, searchRequest, bufferSize)
}

func (p *Pool) SearchWithPaging(searchRequest *ldap.SearchRequest, pagingSize uint32) (*ldap.SearchResult, error) {
	conn, err := p.get()
	defer p.put(conn)
	if err != nil {
		return nil, err
	}

	return conn.SearchWithPaging(searchRequest, pagingSize)
}

func (p *Pool) DirSync(searchRequest *ldap.SearchRequest, flags, maxAttrCount int64, cookie []byte) (*ldap.SearchResult, error) {
	conn, err := p.get()
	defer p.put(conn)
	if err != nil {
		return nil, err
	}

	return conn.DirSync(searchRequest, flags, maxAttrCount, cookie)
}

func (p *Pool) DirSyncAsync(ctx context.Context, searchRequest *ldap.SearchRequest, bufferSize int, flags, maxAttrCount int64, cookie []byte) ldap.Response {
	conn, err := p.get()
	defer p.put(conn)
	if err != nil {
		return ResponseError(err)
	}

	return conn.DirSyncAsync(ctx, searchRequest, bufferSize, flags, maxAttrCount, cookie)
}

func (p *Pool) Syncrepl(ctx context.Context, searchRequest *ldap.SearchRequest, bufferSize int, mode ldap.ControlSyncRequestMode, cookie []byte, reloadHint bool) ldap.Response {
	conn, err := p.get()
	defer p.put(conn)
	if err != nil {
		return ResponseError(err)
	}

	return conn.Syncrepl(ctx, searchRequest, bufferSize, mode, cookie, reloadHint)
}

func (p *Pool) Extended(request *ldap.ExtendedRequest) (er *ldap.ExtendedResponse, err error) {
	conn, err := p.get()
	defer p.put(conn)
	if err != nil {
		return
	}

	return conn.Extended(request)
}
