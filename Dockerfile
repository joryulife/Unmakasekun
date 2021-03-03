FROM golang:1.15.7-alpine3.12 as builder

RUN apk update \
  && apk add --no-cache git curl \
  && go get -u github.com/cosmtrek/air \
  && chmod +x ${GOPATH}/bin/air

WORKDIR /Unmakasekun

COPY go.mod go.sum ./

RUN go mod download
ENV CHANNEL_SECRET=c63cda167f5859ceda5cde822a5a7b5f
ENV CHANNEL_TOKEN=dvLM1fbPEISIwrO5AMIga2wktVeR1PHVG/BhETbrKYl6uNp3swME7x8oPnbHJGnQcsGNHev6mKF4SOI52Blj8spZjBkJUN9Q2qTVKiKXfnA67jWJb5LEwONSgHCZ/UQXzljh+CrkrMVyd7zMzjLJXgdB04t89/1O/w1cDnyilFU=

COPY . .

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s" -o /main ./cmd

FROM alpine:3.12

COPY --from=builder /main .

ENV PORT=${PORT}
ENTRYPOINT ["/main web"]