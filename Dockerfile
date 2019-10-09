FROM golang:1.13.1 AS builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
WORKDIR /go/src/github.com/meriy100/canon
COPY . .
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure
RUN go build server.go

FROM alpine
COPY --from=builder /go/src/github.com/meriy100/canon /app
ENV PORT=${PORT}

ENV DB_USER=${DB_USER}
ENV DB_PASS=${DB_PASS}
ENV DB_HOST=${DB_HOST}
ENV DB_PORT=${DB_PORT}
ENV DB_NAME=${DB_NAME}
ENV DB_SSL_MODE=${DB_SSL_MODE}
CMD /app/server