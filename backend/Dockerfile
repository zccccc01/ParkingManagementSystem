# 第一阶段：构建阶段
FROM golang:1.23-alpine AS builder

# ENV CGO_ENABLED 0 关闭CGO(支持与 C 语言接口的互通)

# 设置工作目录
WORKDIR /app

# 使用国内 Go 模块代理来加速下载（可选）
ENV GOPROXY=https://mirrors.tencent.com/go/

# 将 go.mod 和 go.sum 复制到工作目录
COPY go.mod go.sum /app/

# 下载依赖
RUN go mod download

# 将当前目录的源码复制到工作目录
COPY . /app/

# 构建 Go 应用，将构建输出放置在 /app/main
RUN go build -o main ./cmd

# 第二阶段：生产部署阶段
FROM alpine:3.20.3

# 设置国内镜像源，加速 apk 包下载
RUN echo "http://mirrors.aliyun.com/alpine/v3.20/main" > /etc/apk/repositories && \
    echo "http://mirrors.aliyun.com/alpine/v3.20/community" >> /etc/apk/repositories

# 设置时区为上海
RUN apk add --no-cache tzdata && \
    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone && \
    apk del tzdata

# 设置工作目录
WORKDIR /root/

# 从构建阶段复制编译后的二进制文件
COPY --from=builder /app/main .

# 运行应用
CMD ["./main"]
