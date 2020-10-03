FROM golang:1.15

RUN go get github.com/cosmtrek/air
COPY ./air.conf /etc/air/air.conf

ENTRYPOINT ["air", "-c", "/etc/air/air.conf"]