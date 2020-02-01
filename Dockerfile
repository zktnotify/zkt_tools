FROM golang:1.13.5 AS build

ENV GOPROXY=https://goproxy.io
ENV GO111MODULE on

WORKDIR /go/cache

COPY go.mod .
COPY go.sum .
RUN go mod download

WORKDIR /go/release
COPY . .

# build
RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -installsuffix cgo -o app ./cmd/server/main.go

FROM bashell/alpine-bash:latest as prod
MAINTAINER "zktnotify Bluebuff <1010060164@qq.com>"

WORKDIR zkt-tools

COPY --from=build /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=build /go/release/static ./static
COPY --from=build /go/release/app ./app
COPY --from=build /go/release/configs/config.yaml ./config.yaml
COPY --from=build /go/release/run.sh ./run.sh

RUN chmod 775 run.sh && mkdir -p logs /data/logs

EXPOSE 8080 8080

CMD ["./run.sh"]