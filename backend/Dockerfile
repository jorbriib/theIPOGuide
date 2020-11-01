FROM golang:1.15 as development

COPY ./air.conf /etc/air/air.conf

RUN go get github.com/cosmtrek/air

ENTRYPOINT ["air", "-c", "/etc/air/air.conf"]

FROM golang:1.15 as test

RUN apt-get update && apt-get install -y --no-install-recommends mariadb-client
