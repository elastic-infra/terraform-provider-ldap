package main

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var testAccProviders map[string]*schema.Provider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider()
	testAccProviders = map[string]*schema.Provider{
		"ldap": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ *schema.Provider = Provider()
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("LDAP_HOST"); v == "" {
		t.Fatal("LDAP_HOST must be set for acceptance tests")
	}
	if v := os.Getenv("LDAP_PORT"); v == "" {
		t.Fatal("LDAP_PORT must be set for acceptance tests")
	}
	if v := os.Getenv("LDAP_BIND_USER"); v == "" {
		t.Fatal("LDAP_BIND_USER must be set for acceptance tests")
	}
	if v := os.Getenv("LDAP_BIND_PASSWORD"); v == "" {
		t.Fatal("LDAP_BIND_PASSWORD must be set for acceptance tests")
	}
}
