BINARY=terraform-provider-ldap
PLUGIN_PATH=$(HOME)/.terraform.d/plugins/registry.terraform.io/elastic-infra/ldap/$(shell git describe --tags | tr -d v)/darwin_amd64
TEST_ENV := LDAP_HOST=localhost LDAP_PORT=389 LDAP_BIND_USER="cn=admin,dc=example,dc=com" LDAP_BIND_PASSWORD=admin

.DEFAULT_GOAL: $(BINARY)

$(BINARY):
	go build -o bin/$(BINARY)

install: $(BINARY)
	install -d $(PLUGIN_PATH)
	install -m 775 bin/$(BINARY) $(PLUGIN_PATH)/

test:
	go test -v ./...

docker_test:
	$(TEST_ENV) go test -v ./...
