# 使用指定版本的 python:3.9-alpine3.17 镜像
FROM python:3.9-alpine3.17

# 设置工作目录
WORKDIR /app

# 设置国内镜像源，加速 apk 包下载
RUN echo "https://mirrors.aliyun.com/alpine/v3.17/main/" > /etc/apk/repositories && \
    echo "https://mirrors.aliyun.com/alpine/v3.17/community/" >> /etc/apk/repositories

# 安装必要的系统依赖和 Python MySQL 连接器
RUN apk add --no-cache --virtual .build-deps gcc musl-dev libffi-dev mariadb-connector-c-dev \
    && pip install -i https://mirrors.aliyun.com/pypi/simple/ mysql-connector-python \
    && apk del .build-deps  # 安装完成后删除编译依赖

# 复制脚本到容器
COPY add_all_data.py /app/add_data.py

# 执行 Python 脚本
CMD ["python", "add_data.py"]
