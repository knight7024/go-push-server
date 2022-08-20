FROM golang:1.19.0-alpine as base
RUN apk update && apk add git && apk add ca-certificates

ENV GO111MODULE=on

WORKDIR /usr/src/app
COPY go.mod .
COPY go.sum .
RUN go mod download
RUN go install github.com/swaggo/swag/cmd/swag@latest

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags="-s -w" -o bin/main ./main.go

### LAYER

FROM alpine
RUN apk --no-cache add tzdata && \
	cp /usr/share/zoneinfo/Asia/Seoul /etc/localtime && \
	echo "Asia/Seoul" > /etc/timezone \
	apk del tzdata

COPY --from=base /usr/src/app/bin/main ./main
COPY --from=base /usr/src/app/application_config.yaml ./application_config.yaml

ENV GIN_MODE=release \
	PORT=${WEB_PORT}

CMD ["./main"]