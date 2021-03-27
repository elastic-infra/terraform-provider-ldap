package main

import (
	"crypto/tls"
	"fmt"

	"github.com/go-ldap/ldap/v3"
)

// Config is the set of parameters needed to configure the LDAP provider.
type Config struct {
	LDAPHost     string
	LDAPPort     int
	BindUser     string
	BindPassword string

	StartTLS    bool
	TLS         bool
	TLSInsecure bool
}

func (c *Config) initiateAndBind() (*ldap.Conn, error) {
	conn, err := c.dial()
	if err != nil {
		return nil, err
	}

	// bind to current connection
	err = conn.Bind(c.BindUser, c.BindPassword)
	if err != nil {
		conn.Close()
		return nil, err
	}

	// return the LDAP connection
	return conn, nil
}

func (c *Config) dial() (*ldap.Conn, error) {
	uri := fmt.Sprintf("%s:%d", c.LDAPHost, c.LDAPPort)

	if c.TLS {
		return ldap.DialTLS("tcp", uri, &tls.Config{
			ServerName:         c.LDAPHost,
			InsecureSkipVerify: c.TLSInsecure,
		})
	}

	conn, err := ldap.Dial("tcp", uri)
	if err != nil {
		return nil, err
	}

	if c.StartTLS {
		err = conn.StartTLS(&tls.Config{
			InsecureSkipVerify: true,
		})
		if err != nil {
			return nil, err
		}
	}
	return conn, err
}
