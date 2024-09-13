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
def insert_reservation_data():
    # 获取parkinglot, vehicle, 和 parkingspace表中存在的ID列表
    cursor.execute("SELECT ParkingLotID FROM parkinglot")
    parking_lot_ids = [row[0] for row in cursor.fetchall()]
    cursor.execute("SELECT VehicleID FROM vehicle")
    vehicle_ids = [row[0] for row in cursor.fetchall()]
    cursor.execute("SELECT SpaceID FROM parkingspace")
    space_ids = [row[0] for row in cursor.fetchall()]

    for i in range(12, 101):
        lot_id = random.choice(parking_lot_ids)
        vehicle_id = random.choice(vehicle_ids)
        space_id = random.choice(space_ids)
        start_time = datetime.now() + timedelta(days=random.randint(1, 30), hours=random.randint(0, 23),
                                                minutes=random.randint(0, 59))
        end_time = start_time + timedelta(hours=random.randint(1, 12), minutes=random.randint(0, 59))
        status = random.choice(['Completed', 'Cancelled', 'Doing'])

        query = (
            "INSERT INTO reservation (ReservationID, StartTime, EndTime, SpaceID, VehicleID, LotID, Status) "
            "VALUES (%s, %s, %s, %s, %s, %s, %s)"
        )
        values = (i, start_time, end_time, space_id, vehicle_id, lot_id, status)
        cursor.execute(query, values)

    conn.commit()


# 执行插入数据
insert_reservation_data()

# 关闭数据库连接
cursor.close()
conn.close()

print("100条数据已成功插入到reservation表中。")