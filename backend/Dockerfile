FROM golang:1.15 as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o /go/bin/theipoguide .


FROM scratch as production

COPY --from=builder /go/bin/theipoguide /go/bin/theipoguide

ENTRYPOINT ["/go/bin/theipoguide"]


FROM golang:1.15 as development

COPY --from=builder /go/bin/theipoguide /go/bin/theipoguide
COPY --from=builder /app /app

COPY ./air.conf /etc/air/air.conf

RUN go get github.com/cosmtrek/air

WORKDIR /app

ENTRYPOINT ["air", "-c", "/etc/air/air.conf"]


FROM golang:1.15 as test

WORKDIR /app

RUN apt-get update && apt-get install -y --no-install-recommends mariadb-client
