FROM golang:1.15 as builder

RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.13.0/migrate.linux-amd64.tar.gz | tar xvz
RUN mv migrate.linux-amd64 /go/bin/migrate

FROM alpine as migrations

COPY . /database

COPY --from=builder /go/bin/migrate  /usr/local/bin/migrate

ENTRYPOINT ["sh", "/database/migrate.sh"]
