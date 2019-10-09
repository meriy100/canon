FROM golang:1.13.1 AS builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
WORKDIR /go/src/github.com/meriy100/canon
COPY . .
RUN go build server.go

FROM alpine
COPY --from=builder /go/src/github.com/meriy100/canon /app

CMD /app/server $PORT