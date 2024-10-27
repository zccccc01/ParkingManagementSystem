import mysql.connector
import random
import string
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

# 统一的哈希密码(123456对应的哈希密码)
hashed_password = '$2a$14$CiAORhxAjL/xJMczSFhP.utXyEFIQ0gIIC6dM0TQP2bvNGoLgCKB.'


# 插入users表数据的函数
def insert_users():
    for i in range(1, 51):
        username = ''.join(random.choices(string.ascii_uppercase, k=3))
        tel = "1" + ''.join(random.choices(string.digits, k=10))
        query = "INSERT INTO users (UserID, Username, Password, Tel) VALUES (%s, %s, %s, %s)"
        values = (i, username, hashed_password, tel)
        cursor.execute(query, values)


# 插入vehicle表数据的函数
def insert_vehicles():
    for i in range(1, 51):
        plate_number = ''.join(random.choices(string.ascii_uppercase + string.digits, k=7))
        color = random.choice(['Red', 'Blue', 'Green', 'Black', 'White'])
        query = "INSERT INTO vehicle (VehicleID, UserID, PlateNumber, Color) VALUES (%s, %s, %s, %s)"
        values = (i, i, plate_number, color)  # 这里确保每个用户对应一个车辆
        cursor.execute(query, values)


# 插入parkinglot表数据的函数
def insert_parkinglots():
    # 中国城市及其对应的经纬度
    cities = {
        'BJ': (39.9042, 116.4074),  # 北京
        'SH': (31.2304, 121.4737),  # 上海
        'GZ': (23.1291, 113.2644),  # 广州
        'SZ': (22.5431, 114.0579),  # 深圳
        'CD': (30.5728, 104.0668),  # 成都
        'HZ': (30.2741, 120.1551),  # 杭州
        'WH': (30.5928, 114.3055),  # 武汉
        'NJ': (32.0617, 118.7778),  # 南京
        'TJ': (39.3434, 117.3616),  # 天津
        'XA': (34.3416, 108.9398)  # 西安
    }

    for i, (abbreviation, (latitude, longitude)) in enumerate(cities.items(), start=1):
        parking_name = abbreviation + '-CityLot'
        capacity = 50  # 每个停车场的容量
        rates = round(random.uniform(8.0, 18.0), 2)
        query = "INSERT INTO parkinglot (ParkingLotID, ParkingName, Longitude, Latitude, Capacity, Rates) VALUES (%s, %s, %s, %s, %s, %s)"
        values = (i, parking_name, longitude, latitude, capacity, rates)
        cursor.execute(query, values)


# 插入parkingspace表数据的函数
def insert_parkingspaces():
    for parking_lot_id in range(1, 11):  # 对于每个停车场
        for i in range(1, 51):  # 每个停车场50个车位
            space_id = (parking_lot_id - 1) * 50 + i  # 计算车位ID
            status = random.choice(['FREE', 'OCCUPIED', 'RESERVED'])
            query = "INSERT INTO parkingspace (SpaceID, Status, ParkingLotID) VALUES (%s, %s, %s)"
            values = (space_id, status, parking_lot_id)
            cursor.execute(query, values)


# 插入parkingrecord表数据的函数
def insert_parkingrecords():
    cursor.execute("SELECT SpaceID FROM parkingspace")
    space_ids = [row[0] for row in cursor.fetchall()]
    cursor.execute("SELECT VehicleID FROM vehicle")
    vehicle_ids = [row[0] for row in cursor.fetchall()]
    cursor.execute("SELECT ParkingLotID FROM parkinglot")
    lot_ids = [row[0] for row in cursor.fetchall()]
    for i in range(1, 101):
        space_id = random.choice(space_ids)
        vehicle_id = random.choice(vehicle_ids)
        lot_id = (space_id - 1) // 50 + 1  # 根据SpaceID计算ParkingLotID
        start_time = datetime.now() - timedelta(days=random.randint(1, 30), hours=random.randint(0, 23),
                                                minutes=random.randint(0, 59))
        end_time = start_time + timedelta(hours=random.randint(1, 12), minutes=random.randint(0, 59))
        fee = round(random.uniform(5.0, 50.0), 2)
        query = "INSERT INTO parkingrecord (RecordID, SpaceID, VehicleID, LotID, StartTime, EndTime, Fee) VALUES (%s, %s, %s, %s, %s, %s, %s)"
        values = (i, space_id, vehicle_id, lot_id, start_time, end_time, fee)
        cursor.execute(query, values)


# 插入reservation表数据的函数
def insert_reservations():
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
        start_time = datetime.now() + timedelta(days=random.randint(1, 30), hours=random.randint(0, 23),
                                                minutes=random.randint(0, 59))
        end_time = start_time + timedelta(hours=random.randint(1, 12), minutes=random.randint(0, 59))
        status = random.choice(['Completed', 'Cancelled', 'Doing'])
        query = "INSERT INTO reservation (ReservationID, StartTime, EndTime, SpaceID, VehicleID, LotID, Status) VALUES (%s, %s, %s, %s, %s, %s, %s)"
        values = (i, start_time, end_time, space_id, vehicle_id, lot_id, status)
        cursor.execute(query, values)


# 插入violationrecord表数据的函数
def insert_violations():
    cursor.execute("SELECT RecordID FROM parkingrecord")
    record_ids = [row[0] for row in cursor.fetchall()]
    for i in range(1, 101):
        record_id = random.choice(record_ids)
        fine_amount = round(random.uniform(50.0, 500.0), 2)
        violation_type = random.choice(['OVERSTAY', 'NOPAY'])
        status = random.choice(['PAID', 'UNPAID', 'DISPUTED'])
        query = "INSERT INTO violationrecord (ViolationID, RecordID, FineAmount, ViolationType, Status) VALUES (%s, %s, %s, %s, %s)"
        values = (i, record_id, fine_amount, violation_type, status)
        cursor.execute(query, values)


# 插入paymentrecord表数据的函数
def insert_payments():
    cursor.execute("SELECT RecordID FROM parkingrecord")
    record_ids = [row[0] for row in cursor.fetchall()]
    cursor.execute("SELECT ReservationID FROM reservation")
    reservation_ids = [row[0] for row in cursor.fetchall()]
    for i in range(1, 101):
        record_id = random.choice(record_ids) if record_ids else None
        reservation_id = random.choice(reservation_ids) if reservation_ids else None
        amount = round(random.uniform(5.0, 500.0), 2)
        payment_timestamp = datetime.now() - timedelta(days=random.randint(0, 365), hours=random.randint(0, 23),
                                                       minutes=random.randint(0, 59))
        payment_method = random.choice(['Credit Card', 'VX Pay', 'AliPay', 'Cash'])
        query = "INSERT INTO paymentrecord (PaymentID, RecordID, ReservationID, Amount, PaymentTimestamp, PaymentMethod) VALUES (%s, %s, %s, %s, %s, %s)"
        values = (i, record_id, reservation_id, amount, payment_timestamp, payment_method)
        cursor.execute(query, values)


# 执行插入数据
insert_users()
insert_vehicles()
insert_parkinglots()
insert_parkingspaces()
insert_parkingrecords()
insert_reservations()
insert_violations()
insert_payments()

# 提交事务
conn.commit()

# 关闭数据库连接
cursor.close()
conn.close()

print("Data has been inserted into all tables.")
