import mysql.connector

# 数据库连接配置
config = {
    'host': '127.0.0.1',
    'user': 'root',
    'password': '123456',
    'database': 'chao_db'
}

# 连接到数据库
conn = mysql.connector.connect(**config)
cursor = conn.cursor()

# 表名列表，按照依赖的反方向排序
tables = [
    'paymentrecord', 'violationrecord', 'reservation', 'parkingrecord',
    'parkingspace', 'vehicle', 'users', 'parkinglot'
]


# 删除数据的函数
def delete_data():
    for table in tables:
        query = f"DELETE FROM {table}"
        try:
            cursor.execute(query)
            conn.commit()
            print(f"All data has been deleted from {table}.")
        except mysql.connector.Error as err:
            print(f"Error deleting data from {table}: {err}")


# 执行删除数据
delete_data()

# 关闭数据库连接
cursor.close()
conn.close()
