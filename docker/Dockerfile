FROM golang:1.15 as development

RUN go get github.com/cosmtrek/air
COPY ./air.conf /etc/air/air.conf

ENTRYPOINT ["air", "-c", "/etc/air/air.conf"]


FROM golang:1.15 as test

RUN apt-get update && apt-get install -y --no-install-recommends mariadb-client
