# 使用add_data_*.py这个脚本添加数据

## 1. 安装相关库

在`add_data_*.py`中,需要安装`mysql-connector-python`库来连接MySQL数据库.

```bash
pip install mysql-connector-python
```

## 2. 修改数据库连接信息
在`add_data_*.py`中,修改数据库连接信息,使其指向你本地的数据库.

```python
config = {
    'host': 'your_host',
    'user': 'your_user',
    'password': 'your_password',
    'database': 'your_database'
}
```

## 3. 运行脚本

### 运行顺序
1. 建users表
2. 建vehicle表
3. 建parkinglot表
4. 建parkingspace表
5. 建parkingrecord表
6. 建reservation表
7. 建violationrecord表
8. 建paymentrecord表

在命令行中运行`add_all_data.py`脚本添加数据

```bash
python add_all_data.py
```

在命令行中运行`remove_data_script.py`脚本删库跑路

```bash
python remove_data_script.py
```
