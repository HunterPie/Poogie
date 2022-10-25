FROM golang:1.19.0-alpine3.15 as builder

LABEL org.opencontainers.image.version = 1.8.0

RUN apk add -U --no-cache ca-certificates curl git make gcc g++ libc-dev musl-dev

WORKDIR /app

ADD go.mod go.sum ./
RUN go mod download -x

ADD . .
RUN make

RUN cp ./build/* .

FROM alpine:latest

RUN apk add -U --no-cache ca-certificates curl tzdata

ENV TZ=America/Sao_Paulo
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && \
    echo $TZ > /etc/timezone

WORKDIR /app
COPY --from=builder /app/ .

ENTRYPOINT [ "/app/poogie" ]
