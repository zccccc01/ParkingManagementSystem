# 使用 Nginx 1.26 的 Alpine 轻量版本
FROM nginx:1.26-alpine

# 复制全局 nginx.conf 主配置文件
COPY ./nginx.conf /etc/nginx/nginx.conf

# 复制站点配置文件
COPY ./default.conf /etc/nginx/conf.d/default.conf

# 复制构建后的前端文件
COPY ./build /usr/share/nginx/html

# 暴露端口
EXPOSE 3000

# 启动 Nginx
CMD ["nginx", "-g", "daemon off;"]
