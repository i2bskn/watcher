INSTALL_PATH := /usr/local/bin
WORKING_PATH := /tmp
ARCHIVE_NAME := watchers
ARCHIVE_PATH := $(WORKING_PATH)/$(ARCHIVE_NAME)
ARCHIVE_FILE := $(ARCHIVE_PATH).tar.gz


.PHONY: clean

cmd/service-operator/service-operator:
	cd cmd/service-operator && go build

build: cmd/service-operator/service-operator
	rm -rf $(ARCHIVE_PATH) $(ARCHIVE_FILE)
	mkdir -p $(ARCHIVE_PATH)
	mv cmd/service-operator/service-operator $(ARCHIVE_PATH)
	cd $(WORKING_PATH) && tar zcvf $(ARCHIVE_FILE) $(ARCHIVE_NAME)

install: cmd/service-operator/service-operator
	install cmd/service-operator/service-operator $(INSTALL_PATH)

clean:
	rm -f cmd/service-operator/service-operator
