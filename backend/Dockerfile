FROM golang:1.15 as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o /go/bin/theipoguide .

RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.13.0/migrate.linux-amd64.tar.gz | tar xvz
RUN mv migrate.linux-amd64 /go/bin/migrate


FROM scratch as production

COPY --from=builder /go/bin/theipoguide /go/bin/theipoguide

ENTRYPOINT ["/go/bin/theipoguide"]


FROM golang:1.15 as development

RUN apt-get update && apt-get install -y --no-install-recommends mariadb-client

COPY --from=builder /go/bin/theipoguide /go/bin/theipoguide
COPY --from=builder /go/bin/migrate /go/bin/migrate
COPY --from=builder /app /app

COPY ./air.conf /etc/air/air.conf

RUN go get github.com/cosmtrek/air

WORKDIR /app

ENTRYPOINT ["air", "-c", "/etc/air/air.conf"]


FROM golang:1.15 as test

RUN apt-get update && apt-get install -y --no-install-recommends mariadb-client

COPY --from=builder /go/bin/migrate /go/bin/migrate

WORKDIR /app

