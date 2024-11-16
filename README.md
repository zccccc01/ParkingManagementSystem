# 智能停车管理系统

## 前端运行

```bash
cd ./frontend
npm install
npm start
```

## 后端运行

```bash
cd ./backend/cmd
air
```

## 利用 docker 运行

```bash
git checkout docker
docker-compose up -d db_master db_slave1 db_slave2 redis #这里要先设置数据库主从复制
docker-compose up --build python
docker-compose up -d --build backend1 backend2 backend3
docker-compose up -d --build frontend
```

## 前端

- react

## 后端

- fiter 框架
- gorm
- air
- swagger 生成接口文档

## 数据库

- mysql
- redis
