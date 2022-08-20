# Go Push Server API

## Overview
Go언어로 구현된 푸시 서버입니다.

## Go Version
Go v1.19으로 개발되었습니다.

## Getting Started
```
mkdir -p /usr/docker/redis
wget http://download.redis.io/redis-stable/redis.conf -O /usr/docker/redis/redis.conf
```
```
docker-compose -f "docker-compose.yml" up -d --build
```

## Server Specification
* HTTP
* MySQL
* Redis
* Docker, Docker Compose

## Used Library
* [gin](https://github.com/gin-gonic/gin)
* [entgo](https://github.com/ent/ent)
* [go-cache](https://github.com/patrickmn/go-cache)
* [viper](https://github.com/spf13/viper)
* [swaggo](https://github.com/swaggo/swag)
* [firebase-admin-go](https://github.com/firebase/firebase-admin-go)
