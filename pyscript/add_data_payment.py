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
def insert_paymentrecord_data():
    # 获取parkingrecord和reservation表中存在的RecordID和ReservationID列表
    cursor.execute("SELECT RecordID FROM parkingrecord")
    parking_record_ids = [row[0] for row in cursor.fetchall()]
    cursor.execute("SELECT ReservationID FROM reservation")
    reservation_ids = [row[0] for row in cursor.fetchall()]

    for i in range(1, 101):
        record_id = random.choice(parking_record_ids) if parking_record_ids else None
        reservation_id = random.choice(reservation_ids) if reservation_ids else None
        amount = round(random.uniform(5.0, 500.0), 2)  # 支付金额在5到500之间
        payment_timestamp = datetime.now() - timedelta(days=random.randint(0, 365), hours=random.randint(0, 23),
                                                       minutes=random.randint(0, 59))
        payment_method = random.choice(['Credit Card', 'VX Pay', 'AliPay', 'Cash'])  # 支付方式

        query = (
            "INSERT INTO paymentrecord (PaymentID, RecordID, ReservationID, Amount, PaymentTimestamp, PaymentMethod) "
            "VALUES (%s, %s, %s, %s, %s, %s)"
        )
        values = (i, record_id, reservation_id, amount, payment_timestamp, payment_method)
        cursor.execute(query, values)

    conn.commit()


# 执行插入数据
insert_paymentrecord_data()

# 关闭数据库连接
cursor.close()
conn.close()

print("100条数据已成功插入到paymentrecord表中。")