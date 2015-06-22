SHELL=/bin/bash
VERSION := $(shell cat VERSION)
ITTERATION := $(shell date +%s)

# # drone build
# sudo apt-get update
# sudo apt-get install ruby-dev build-essential rubygems wget curl
# sudo gem install fpm
# make deps test build deb gzip

all: build

deps:
	go get -t -v ./...

test: 
	go test ./... -v

build:
	mkdir -p build/output
	mkdir -p build/root/usr/bin
	mkdir -p build/root/etc/init.d
	go build -o build/root/usr/bin/srv-proxy
	cp srv-proxy.init build/root/etc/init.d/srv-proxy
	chmod 755 build/root/etc/init.d/srv-proxy

install:
	install -t /usr/bin build/root/usr/bin/srv-proxy

clean:
	rm -rf build

# sudo apt-get install ruby-dev build-essential
# sudo gem install fpm
deb: build
	fpm -s dir -t deb -n srv-proxy -v $(VERSION) -p build/srv-proxy-amd64.deb \
		--deb-priority optional \
		--force \
		--iteration $(ITTERATION) \
		--deb-compression bzip2 \
		--url https://github.com/benschw/srv-proxy \
		--description "Reverse proxy for loadbalancing with SRV records" \
		-m "Ben Schwartz <benschw@gmail.com>" \
		--license "Apache License 2.0" \
		--vendor "fliglio.com" -a amd64 \
		build/root/=/

