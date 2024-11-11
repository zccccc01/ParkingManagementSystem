# 智能停车管理系统

## 前端运行

### 方法 1

```bash
cd ./frontend
npm install
npm start
```

### 方法 2(windows)

```bash
cd ./frontend
npm install
npm run build
mv build ../
```

修改/nginx/conf/nginx.conf

```
root C:/gocode/ParkingManagementSystem/build; #修改为你的完整路径(一定要用/)

location / {
    try_files $uri $uri/ /index.html;
}

location /static/ {
    alias C:/gocode/ParkingManagementSystem/build/static/; #修改为你的完整路径
}
```

```bash
cd ./nginx
mkdir temp; cd temp; mkdir client_body_temp
cd ..
./nginx.exe
```

### 优雅关闭 nginx(windows)

```bash
cd ./nginx
./nginx.exe -s quit
```

## 后端运行

```bash
cd ./backend/cmd
air
```

## 使用docker运行

```bash
cd ./frontend
npm install
npm run build
mv build ../web
docker-compose up -d --build
```

## 前端

- react

## 后端

- fiter 框架
- grom
- air

## 数据库

- mysql
- redis

## 有疑问向我提issue
## 有些部分写的挺乱,有待改进
## 欢迎star
## 欢迎fork
## 欢迎pr