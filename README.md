# Go Push Server API

## Overview
Go언어로 구현된 푸시 서버입니다.

## Go Version
Go v1.19으로 개발되었습니다.

## Getting Started
### 0. Configurate application and docker settings
1. Change name `application_config_example.yml` to `application_config.yml` and `.env_example` to `.env`
2. Write value in files by format

### 1. Redis
```bash
sudo mkdir -p /usr/local/docker/redis
sudo wget http://download.redis.io/redis-stable/redis.conf -O /usr/local/docker/redis/redis.conf
```

### 2. Docker
```bash
docker compose -f "docker-compose.yml" up -d --build
```

### 3. MySQL
```bash
docker exec -it mysql-8.0.30 /bin/bash
```
```SQL
CREATE SCHEMA `push_server` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;
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
