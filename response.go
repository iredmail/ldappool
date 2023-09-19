package ldappool

import (
	"github.com/go-ldap/ldap/v3"
)

func ResponseError(err error) ldap.Response {
	return response{err: err}
}

type response struct {
	err error
}

func (r response) Entry() *ldap.Entry {
	return nil
}

func (r response) Referral() string {
	return ""
}

func (r response) Controls() []ldap.Control {
	return nil
}

func (r response) Err() error {
	return r.err
}

func (r response) Next() bool {
	return false
}
