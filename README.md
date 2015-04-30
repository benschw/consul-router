# srv-proxy

based on [jmcarbo/consul-router](https://github.com/jmcarbo/consul-router)

	go build -o srv-proxy
	docker build -t benschw/srv-proxy .
	
	docker run -d -e TARGET=service.consul -e DOMAIN=edge -p 80:8080 benschw/srv-proxy

