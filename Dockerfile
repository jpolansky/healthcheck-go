FROM golang:1.20rc2-bullseye

RUN go install github.com/jpolansky/healthcheck-go@latest

ENV GO1111MODULE=on
ENV GOFLAGS=-mod=vendor

ENV APP_HOME /go/src/healthcheck-go
RUN mkdir -p "$APP_HOME"

WORKDIR "$APP_HOME"
EXPOSE 8080
CMD ["healthcheck-go", "run"]