INSTALL_PATH := /usr/local/bin

.PHONY: clean

cmd/service-operator/service-operator:
	cd cmd/service-operator && go build

install: cmd/service-operator/service-operator
	install cmd/service-operator/service-operator $(INSTALL_PATH)

clean:
	rm -f cmd/service-operator/service-operator
