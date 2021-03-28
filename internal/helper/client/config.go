package client

type Config struct {
	LDAPHost     string
	LDAPPort     int
	BindUser     string
	BindPassword string

	StartTLS    bool
	TLS         bool
	TLSInsecure bool
}
