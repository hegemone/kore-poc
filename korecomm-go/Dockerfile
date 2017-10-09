FROM golang:1.8.3-alpine3.6

COPY ./ /go/src/github.com/hegemone/kore-poc/korecomm-go/
WORKDIR /go/src/github.com/hegemone/kore-poc/korecomm-go/

RUN apk update && apk add make gcc musl-dev
RUN make build
RUN install -mode=+x \
  /go/src/github.com/hegemone/kore-poc/korecomm-go/build/korecomm \
  /usr/bin/korecomm
RUN install -mode=+x \
  /go/src/github.com/hegemone/kore-poc/korecomm-go/extra/entrypoint.sh \
  /usr/bin/entrypoint.sh
RUN mkdir -p /usr/lib/kore && \
  install \
  /go/src/github.com/hegemone/kore-poc/korecomm-go/build/*.so \
  /usr/lib/kore

ENTRYPOINT ["entrypoint.sh"]
