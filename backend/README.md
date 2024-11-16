# go verison 1.23.0

## 安装相关依赖

`go mod tity`

## 热重载

`go install github.com/air-verse/air@v1.60.0`

## 生成接口文档

`go install github.com/swaggo/swag/cmd/swag@latest`

```bash
cd ./backend
swag init -g cmd/main.go
```
