INSTALL_PATH ?= /usr/local/bin
OUT_PATH     ?= out
PKG_NAME     := watchers
PKG_VERSION  ?= $(shell date +%Y%m%d)
PKG_PATH     := $(PKG_NAME)-$(PKG_VERSION)
PKG_TAR_PATH := $(PKG_PATH).tar.gz
BUILD_OS     ?= linux
BUILD_ARCH   ?= amd64
CMD_PATHS    := $(wildcard cmd/*)
CMD_FILES    := $(foreach path, $(CMD_PATHS), $(path)/$(notdir $(path)))

.PHONY: build tar install clean

$(PKG_TAR_PATH): $(CMD_FILES)
	rm -rf $(OUT_PATH)/$(PKG_PATH) $(OUT_PATH)/$(PKG_TAR_PATH)
	mkdir -p $(OUT_PATH)/$(PKG_PATH)
	$(foreach file, $(CMD_FILES), cp $(file) $(OUT_PATH)/$(PKG_PATH))
	cd $(OUT_PATH) && tar zcvf $(PKG_TAR_PATH) $(PKG_PATH)

$(CMD_FILES):
	$(foreach path, $(CMD_PATHS), cd $(path) && GOOS=$(BUILD_OS) GOARCH=$(BUILD_ARCH) go build)

build: $(CMD_FILES)

tar: $(PKG_TAR_PATH)

clean:
	$(foreach file, $(CMD_FILES), rm -f $(file))
	rm -rf $(OUT_PATH)

install: $(CMD_FILES)
	$(foreach file, $(CMD_FILES), install $(file) $(INSTALL_PATH))
