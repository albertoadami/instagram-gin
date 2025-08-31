FROM alpine:latest

RUN apk add --no-cache ca-certificates postgresql-client wget

RUN wget -O /usr/local/bin/goose \
    https://github.com/pressly/goose/releases/download/v3.15.0/goose_linux_x86_64 && \
    chmod +x /usr/local/bin/goose

COPY . /migrations
WORKDIR /migrations
ENTRYPOINT ["goose"]
CMD ["--help"]