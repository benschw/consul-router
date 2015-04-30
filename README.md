# consul-router

based on [jmcarbo/consul-router](https://github.com/jmcarbo/consul-router)

	go build -o consul-router
	docker build -t benschw/consul-router .
	
	docker run -d -e TARGET=service.consul -e DOMAIN=edge -p 80:8080 benschw/consul-router

