FROM gliderlabs/alpine:latest
MAINTAINER Random J Farmer <random.j.farmer@gmail.com>

VOLUME /srv

EXPOSE 8080

ENV ZKM_DB_NAME   /srv/zkill-mirror.bolt
ENV ZKM_BOBS_NAME /srv/zkill-mirror.bobs
ENV ZKM_PORT      8080

RUN apk add --no-cache ca-certificates

ADD zkill-mirror zkill-mirror.toml /

CMD ["/zkill-mirror", "serve"]
