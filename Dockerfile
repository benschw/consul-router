FROM benschw/litefs

ADD srv-proxy /usr/bin/srv-proxy

EXPOSE 8080

ENTRYPOINT /usr/bin/srv-proxy
