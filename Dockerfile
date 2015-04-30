FROM benschw/litefs

ADD consul-router /usr/bin/consul-router

EXPOSE 8080

ENTRYPOINT /usr/bin/consul-router
