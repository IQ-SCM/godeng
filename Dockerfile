FROM golang:1.18 AS build_base

WORKDIR /tmp/godeng

COPY . .

RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build -o bin/godeng cmd/main.go

FROM alpine:3.14

WORKDIR /app

COPY --from=build_base /tmp/godeng/bin/godeng /app/godeng

CMD ["/app/godeng", "--config", "/app/config.json"]