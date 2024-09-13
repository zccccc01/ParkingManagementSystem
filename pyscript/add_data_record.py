import mysql.connector
import random
from datetime import datetime, timedelta

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
def insert_parkingrecord_data():
    # 获取parkinglot, vehicle, 和 parkingspace表中存在的ID列表
    cursor.execute("SELECT ParkingLotID FROM parkinglot")
    parking_lot_ids = [row[0] for row in cursor.fetchall()]
    cursor.execute("SELECT VehicleID FROM vehicle")
    vehicle_ids = [row[0] for row in cursor.fetchall()]
    cursor.execute("SELECT SpaceID FROM parkingspace")
    space_ids = [row[0] for row in cursor.fetchall()]

    for i in range(1, 101):
        lot_id = random.choice(parking_lot_ids)
        vehicle_id = random.choice(vehicle_ids)
        space_id = random.choice(space_ids)
        start_time = datetime.now() - timedelta(days=random.randint(1, 30), hours=random.randint(0, 23),
                                                minutes=random.randint(0, 59))
        end_time = start_time + timedelta(hours=random.randint(1, 12), minutes=random.randint(0, 59))
        fee = round(random.uniform(5.0, 50.0), 2)  # 费用在5到50之间

        query = (
            "INSERT INTO parkingrecord (RecordID, SpaceID, VehicleID, LotID, StartTime, EndTime, Fee) "
            "VALUES (%s, %s, %s, %s, %s, %s, %s)"
        )
        values = (i, space_id, vehicle_id, lot_id, start_time, end_time, fee)
        cursor.execute(query, values)

    conn.commit()


# 执行插入数据
insert_parkingrecord_data()

# 关闭数据库连接
cursor.close()
conn.close()

print("100条数据已成功插入到parkingrecord表中。")