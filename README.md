[![Build Status](https://drone.io/github.com/benschw/srv-proxy/status.png)](https://drone.io/github.com/benschw/srv-proxy/latest)


[Download](https://drone.io/github.com/benschw/srv-proxy/files/srv-proxy)

# srv-proxy

based on [jmcarbo/consul-router](https://github.com/jmcarbo/consul-router)

## Usage

	go build
	TARGET=service.consul DOMAIN=edge NS_IP=127.0.0.1 NS_PORT=53 ./srv-proxy

## Docker

	go build -o srv-proxy
	docker build -t benschw/srv-proxy .
	
	docker run -d -e TARGET=service.consul -e DOMAIN=edge -p 80:8080 benschw/srv-proxy



