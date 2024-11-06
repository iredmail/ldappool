package ldappool

import (
	"context"

	"github.com/go-ldap/ldap/v3"
)

func (p *Pool) Add(request *ldap.AddRequest) error {
	conn, err := p.conn()
	if err != nil {
		return err
	}

	defer p.close(conn)

	return conn.Add(request)
}

func (p *Pool) Del(request *ldap.DelRequest) error {
	conn, err := p.conn()
	if err != nil {
		return err
	}

	defer p.close(conn)

	return conn.Del(request)
}

func (p *Pool) Modify(request *ldap.ModifyRequest) error {
	conn, err := p.conn()
	if err != nil {
		return err
	}

	defer p.close(conn)

	return conn.Modify(request)
}

func (p *Pool) ModifyDN(request *ldap.ModifyDNRequest) error {
	conn, err := p.conn()
	if err != nil {
		return err
	}

	defer p.close(conn)

	return conn.ModifyDN(request)
}

func (p *Pool) ModifyWithResult(request *ldap.ModifyRequest) (*ldap.ModifyResult, error) {
	conn, err := p.conn()
	if err != nil {
		return nil, err
	}

	defer p.close(conn)

	return conn.ModifyWithResult(request)
}

func (p *Pool) Compare(dn, attribute, value string) (bool, error) {
	conn, err := p.conn()
	if err != nil {
		return false, err
	}

	defer p.close(conn)

	return conn.Compare(dn, attribute, value)
}

func (p *Pool) PasswordModify(request *ldap.PasswordModifyRequest) (*ldap.PasswordModifyResult, error) {
	conn, err := p.conn()
	if err != nil {
		return nil, err
	}

	defer p.close(conn)

	return conn.PasswordModify(request)
}

func (p *Pool) Search(request *ldap.SearchRequest) (*ldap.SearchResult, error) {
	conn, err := p.conn()
	if err != nil {
		return nil, err
	}

	defer p.close(conn)

	return conn.Search(request)
}

func (p *Pool) SearchAsync(ctx context.Context, searchRequest *ldap.SearchRequest, bufferSize int) ldap.Response {
	conn, err := p.conn()
	if err != nil {
		return ResponseError(err)
	}

	defer p.close(conn)

	return conn.SearchAsync(ctx, searchRequest, bufferSize)
}

func (p *Pool) SearchWithPaging(searchRequest *ldap.SearchRequest, pagingSize uint32) (*ldap.SearchResult, error) {
	conn, err := p.conn()
	if err != nil {
		return nil, err
	}

	defer p.close(conn)

	return conn.SearchWithPaging(searchRequest, pagingSize)
}

func (p *Pool) DirSync(searchRequest *ldap.SearchRequest, flags, maxAttrCount int64, cookie []byte) (*ldap.SearchResult, error) {
	conn, err := p.conn()
	if err != nil {
		return nil, err
	}

	defer p.close(conn)

	return conn.DirSync(searchRequest, flags, maxAttrCount, cookie)
}

func (p *Pool) DirSyncAsync(ctx context.Context, searchRequest *ldap.SearchRequest, bufferSize int, flags, maxAttrCount int64, cookie []byte) ldap.Response {
	conn, err := p.conn()
	if err != nil {
		return ResponseError(err)
	}

	defer p.close(conn)

	return conn.DirSyncAsync(ctx, searchRequest, bufferSize, flags, maxAttrCount, cookie)
}

func (p *Pool) Syncrepl(ctx context.Context, searchRequest *ldap.SearchRequest, bufferSize int, mode ldap.ControlSyncRequestMode, cookie []byte, reloadHint bool) ldap.Response {
	conn, err := p.conn()
	if err != nil {
		return ResponseError(err)
	}

	defer p.close(conn)

	return conn.Syncrepl(ctx, searchRequest, bufferSize, mode, cookie, reloadHint)
}
