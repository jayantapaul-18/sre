GOCMD=go
MY_CLI_CONFIG=config.json
VERSION := $(shell git describe --always --long --dirty)

all: build

build:
	$(GOCMD) build -o sre
	$(info ******************** build ********************)
run:
	./sre
	$(info ******************** run ********************)
clean:
	rm -f sre
    $(info ******************** clean ********************)
build-with-config:
	$(GOCMD) build -o sre
	$(info ******************** build-with-config ********************)
