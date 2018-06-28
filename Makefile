THIS_DIR   = $(dir $(realpath $(firstword $(MAKEFILE_LIST))))
GO         = env GOPATH="$(THIS_DIR)" go
GO_INSTALL = $(GO) install
BIN_DIR    = $(THIS_DIR)/bin

BINARIES   = bin/server
COMMANDS   =  

.PHONY: $(COMMANDS) clean

all: $(BINARIES)

bin/server: src/server/server.go
	$(GO_INSTALL) server

clean:
	$(GO) clean -i $(patsubst bin/%,%,$(BINARIES))
