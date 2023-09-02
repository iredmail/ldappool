## How to use it ?

```go
package main

import(
	"fmt"
	"log"
	
	"github.com/go-ldap/ldap/v3"
	"github.com/iredmail/ldappool"
)

func main() {
	opts := []ldappool.Option{
		ldappool.WithMaxConnections(10),
		ldappool.WithBindCredentials("dn", "password"),
    }
	
	pool, err := ldappool.New("ldap://ldap.example.com:389", opts...)
	if err != nil {
		log.Fatal(err)
    }
	
	defer pool.Close()

	searchRequest := ldap.NewSearchRequest(
		"dc=example,dc=com", // The base dn to search
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		"(&(objectClass=organizationalPerson))", // The filter to apply
		[]string{"dn", "cn"},                    // A list attributes to retrieve
		nil,
	)

	sr, err := pool.Search(searchRequest)
	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range sr.Entries {
		fmt.Printf("%s: %v\n", entry.DN, entry.GetAttributeValue("cn"))
	}
}
```