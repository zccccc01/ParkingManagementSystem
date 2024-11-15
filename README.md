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

### 配置数据库主从

1. **配置主节点**:
   在 `ParkingManagementSystem/` 目录下运行以下命令启动 MySQL 主从复制环境

- 进入 `db_master` 容器:

```bash
docker exec -it db_master mysql -u root -p123456
```

- 登录 MySQL 并设置 root 用户密码:

```bash
ALTER USER 'root'@'%' IDENTIFIED WITH mysql_native_password BY '123456';
FLUSH PRIVILEGES;
exit;
```

- 重启主节点容器:

```bash
docker-compose restart db_master
```

- 获取主节点的 `File` 和 `Position` 字段值:

```bash
docker exec -it db_master mysql -u root -p123456 -e "SHOW MASTER STATUS;"
```

记下输出的 `File` 和 `Position` 值.你将需要在从节点配置时使用这些值.

2. **配置从节点**:

   对于每个从节点(`db_slave1` 和 `db_slave2`):

   - 进入 `db_slave1` 容器(针对第一个从节点):

   ```bash
   docker exec -it db_slave1 mysql -u root -p123456
   ```

   - 登录 MySQL 并配置从节点:

   ```bash
   CHANGE MASTER TO
     MASTER_HOST='db_master',
     MASTER_USER='root',
     MASTER_PASSWORD='123456',
     MASTER_LOG_FILE='mysql-master-bin.000004',  # 替换为主节点的 File 值
     MASTER_LOG_POS=157;  # 替换为主节点的 Position 值
   START SLAVE;
   SHOW SLAVE STATUS \G;
   ```

   - 对于第二个从节点 `db_slave2`,重复上述步骤

   - 通过 `SHOW SLAVE STATUS \G;` 检查每个从节点的状态,确保 `Slave_IO_Running` 和 `Slave_SQL_Running` 都为 `Yes`.

## 前端

- react

## 后端

- fiter 框架
- gorm
- air

## 数据库

- mysql
- redis
