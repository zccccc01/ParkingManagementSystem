import mysql.connector
import random

# 数据库连接配置
config = {
    'host': '127.0.0.1',
    'user': 'root',
    'password': '123456',
    'database': 'chao_db'  # 请替换为您的数据库名称
}

# 连接到数据库
conn = mysql.connector.connect(**config)
cursor = conn.cursor()


# 插入数据的函数
def insert_parkingspace_data():
    # 获取parkinglot表中存在的ParkingLotID列表
    cursor.execute("SELECT ParkingLotID FROM parkinglot")
    parking_lot_ids = [row[0] for row in cursor.fetchall()]

    for i in range(1, 101):
        parking_lot_id = random.choice(parking_lot_ids)  # 随机选择一个存在的ParkingLotID
        status = random.choice(['FREE', 'OCCUPIED', 'RESERVED'])  # 随机选择状态

        query = (
            "INSERT INTO parkingspace (SpaceID, Status, ParkingLotID) "
            "VALUES (%s, %s, %s)"
        )
        values = (i, status, parking_lot_id)
        cursor.execute(query, values)

    conn.commit()


# 执行插入数据
insert_parkingspace_data()

# 关闭数据库连接
cursor.close()
conn.close()

print("100条数据已成功插入到parkingspace表中。")