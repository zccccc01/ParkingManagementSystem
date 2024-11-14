import mysql.connector
import random
import string
from decimal import Decimal
from datetime import datetime, timedelta
import logging

# 配置日志
logging.basicConfig(level=logging.INFO, format='%(asctime)s - %(levelname)s - %(message)s')
logger = logging.getLogger()

# 数据库连接配置
config = {
    'host': '10.1.0.20',
    'user': 'root',
    'password': '123456',
    'database': 'mydb'  # 请替换为您的数据库名称
}

# 连接到数据库
conn = mysql.connector.connect(**config)
cursor = conn.cursor()

# 统一的密码哈希(123456)
hashed_password = '$2a$14$CiAORhxAjL/xJMczSFhP.utXyEFIQ0gIIC6dM0TQP2bvNGoLgCKB.'


# 插入users表数据的函数
def insert_users():
    logger.info("开始插入users表数据...")
    for i in range(1, 51):
        username = ''.join(random.choices(string.ascii_uppercase, k=3))
        tel = "1" + ''.join(random.choices(string.digits, k=10))
        query = "INSERT INTO users (UserID, Username, Password, Tel) VALUES (%s, %s, %s, %s)"
        values = (i, username, hashed_password, tel)
        cursor.execute(query, values)
        logger.debug(f"插入用户数据: {values}")
    logger.info("users表数据插入完成.")


# 插入vehicle表数据的函数
def insert_vehicles():
    logger.info("开始插入vehicle表数据...")
    for i in range(1, 51):
        plate_number = ''.join(random.choices(string.ascii_uppercase + string.digits, k=7))
        color = random.choice(['Red', 'Blue', 'Green', 'Black', 'White'])
        query = "INSERT INTO vehicle (VehicleID, UserID, PlateNumber, Color) VALUES (%s, %s, %s, %s)"
        values = (i, i, plate_number, color)  # 这里确保每个用户对应一个车辆
        cursor.execute(query, values)
        logger.debug(f"插入车辆数据: {values}")
    logger.info("vehicle表数据插入完成.")


# 插入parkinglot表数据的函数
def insert_parkinglots():
    logger.info("开始插入parkinglot表数据...")
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

    rates = {
        'BJ': 18.0,  # 北京
        'SH': 16.0,  # 上海
        'GZ': 12.0,  # 广州
        'SZ': 15.0,  # 深圳
        'CD': 12.5,  # 成都
        'HZ': 14.5,  # 杭州
        'WH': 13.0,  # 武汉
        'NJ': 10.0,  # 南京
        'TJ': 12.0,  # 天津
        'XA': 11.5  # 西安
    }

    for i, (abbreviation, (latitude, longitude)) in enumerate(cities.items(), start=1):
        parking_name = abbreviation + '-CityLot'
        capacity = 50  # 每个停车场的容量
        rate = rates[abbreviation]  # 使用对应城市的费率
        query = "INSERT INTO parkinglot (ParkingLotID, ParkingName, Longitude, Latitude, Capacity, Rates) VALUES (%s, %s, %s, %s, %s, %s)"
        values = (i, parking_name, longitude, latitude, capacity, rate)
        cursor.execute(query, values)
        logger.debug(f"插入停车场数据: {values}")
    logger.info("parkinglot表数据插入完成.")


# 插入parkingspace表数据的函数
def insert_parkingspaces():
    logger.info("开始插入parkingspace表数据...")
    for parking_lot_id in range(1, 11):  # 对于每个停车场
        for i in range(1, 51):  # 每个停车场50个车位
            space_id = (parking_lot_id - 1) * 50 + i  # 计算车位ID
            status = random.choice(['FREE', 'OCCUPIED', 'RESERVED'])
            query = "INSERT INTO parkingspace (SpaceID, Status, ParkingLotID) VALUES (%s, %s, %s)"
            values = (space_id, status, parking_lot_id)
            cursor.execute(query, values)
            logger.debug(f"插入车位数据: {values}")
    logger.info("parkingspace表数据插入完成.")


# 插入parkingrecord表数据的函数
def insert_parkingrecords():
    logger.info("开始插入parkingrecord表数据...")
    cursor.execute("SELECT SpaceID FROM parkingspace")
    space_ids = [row[0] for row in cursor.fetchall()]
    cursor.execute("SELECT VehicleID FROM vehicle")
    vehicle_ids = [row[0] for row in cursor.fetchall()]
    cursor.execute("SELECT ParkingLotID FROM parkinglot")
    lot_ids = [row[0] for row in cursor.fetchall()]

    cursor.execute("SELECT ParkingLotID, Rates FROM parkinglot")
    rates = {row[0]: row[1] for row in cursor.fetchall()}

    for i in range(1, 51):  # 插入50条数据
        space_id = random.choice(space_ids)
        vehicle_id = random.choice(vehicle_ids)
        lot_id = (space_id - 1) // 50 + 1  # 根据SpaceID计算ParkingLotID

        start_time = datetime.now() - timedelta(days=random.randint(1, 30), hours=random.randint(0, 23),
                                                minutes=random.randint(0, 59))
        end_time = start_time + timedelta(hours=random.randint(1, 12), minutes=random.randint(0, 59))

        total_hours = Decimal((end_time - start_time).total_seconds()) / Decimal(3600.0)

        rate = rates.get(lot_id, Decimal(0))
        fee = round(rate * total_hours, 2)

        query = "INSERT INTO parkingrecord (RecordID, SpaceID, VehicleID, LotID, StartTime, EndTime, Fee) VALUES (%s, %s, %s, %s, %s, %s, %s)"
        values = (i, space_id, vehicle_id, lot_id, start_time, end_time, fee)
        cursor.execute(query, values)
        logger.debug(f"插入停车记录数据: {values}")
    logger.info("parkingrecord表数据插入完成.")


# 插入reservation表数据的函数
def insert_reservations():
    logger.info("开始插入reservation表数据...")
    try:
        # 获取 parkinglot, vehicle, parkingspace 的 ID 数据
        cursor.execute("SELECT ParkingLotID FROM parkinglot")
        parking_lot_ids = [row[0] for row in cursor.fetchall()]
        cursor.execute("SELECT VehicleID FROM vehicle")
        vehicle_ids = [row[0] for row in cursor.fetchall()]
        cursor.execute("SELECT SpaceID FROM parkingspace")
        space_ids = [row[0] for row in cursor.fetchall()]

        # 插入 50 条数据
        for i in range(1, 51):
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
            logger.debug(f"插入预定数据: {values}")

        logger.info("reservation表数据插入完成.")
    except mysql.connector.Error as err:
        logger.error(f"插入reservation表数据时出现数据库错误: {err}")
        conn.rollback()  # 出现错误时回滚事务

# 插入violationrecord表数据的函数
def insert_violations():
    logger.info("开始插入violationrecord表数据...")
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
        logger.debug(f"插入违章记录数据: {values}")
    logger.info("violationrecord表数据插入完成.")


# 插入paymentrecord表数据的函数
def insert_payments():
    logger.info("开始插入paymentrecord表数据...")
    cursor.execute("SELECT RecordID, Fee FROM parkingrecord")
    records = cursor.fetchall()
    cursor.execute("SELECT ReservationID, StartTime, EndTime, LotID FROM reservation")
    reservations = cursor.fetchall()
    cursor.execute("SELECT ParkingLotID, Rates FROM parkinglot")
    rates = {row[0]: row[1] for row in cursor.fetchall()}

    used_record_ids = set()
    used_reservation_ids = set()
    total_payments = 0  # 记录已插入的支付条数

    while total_payments < 100:
        # 随机选择使用 RecordID 或 ReservationID
        if records and (not reservations or random.choice([True, False])):
            # 从 parkingrecord 中选择
            record = random.choice(records)
            record_id = record[0]
            if record_id in used_record_ids:
                continue  # 跳过已使用的 RecordID
            used_record_ids.add(record_id)
            amount = record[1]  # 从 parkingrecord 中获取费用
            reservation_id = None  # ReservationID 置空
        elif reservations:
            # 从 reservation 中选择
            reservation = random.choice(reservations)
            reservation_id = reservation[0]
            if reservation_id in used_reservation_ids:
                continue  # 跳过已使用的 ReservationID
            used_reservation_ids.add(reservation_id)
            start_time = reservation[1]
            end_time = reservation[2]
            lot_id = reservation[3]

            # 计算费用
            total_hours = (end_time - start_time).total_seconds() / 3600
            rate = rates.get(lot_id, Decimal('0'))
            amount = round(rate * Decimal(total_hours), 2) if total_hours > 0 else 0

            record_id = None  # RecordID 置空
        else:
            continue  # 如果没有可用的 ID，跳过此次插入

        payment_timestamp = datetime.now() - timedelta(days=random.randint(0, 365), hours=random.randint(0, 23),
                                                       minutes=random.randint(0, 59))
        payment_method = random.choice(['Credit Card', 'VX Pay', 'AliPay', 'Cash'])

        query = "INSERT INTO paymentrecord (PaymentID, RecordID, ReservationID, Amount, PaymentTimestamp, PaymentMethod) VALUES (%s, %s, %s, %s, %s, %s)"
        values = (total_payments + 1, record_id, reservation_id, amount, payment_timestamp, payment_method)
        cursor.execute(query, values)
        logger.debug(f"插入支付记录数据: {values}")

        total_payments += 1  # 增加已插入支付条数

    logger.info("paymentrecord表数据插入完成.")


# 执行插入数据
try:
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

    logger.info("所有数据已成功插入数据库.")

except mysql.connector.Error as err:
    logger.error(f"数据库错误: {err}")
    conn.rollback()  # 如果发生错误，回滚事务

finally:
    # 关闭数据库连接
    cursor.close()
    conn.close()

    logger.info("数据库连接已关闭.")
