import mysql.connector
import random
from datetime import datetime

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
def insert_violationrecord_data():
    # 获取parkingrecord表中存在的RecordID列表
    cursor.execute("SELECT RecordID FROM parkingrecord")
    parking_record_ids = [row[0] for row in cursor.fetchall()]

    for i in range(11, 101):
        record_id = random.choice(parking_record_ids) if parking_record_ids else None
        fine_amount = round(random.uniform(50.0, 500.0), 2)  # 罚款金额在50到500之间
        violation_type = random.choice(['OVERSTAY', 'NOPAY'])  # 违规类型
        status = random.choice(['PAID', 'UNPAID', 'DISPUTED'])  # 状态

        query = (
            "INSERT INTO violationrecord (ViolationID, RecordID, FineAmount, ViolationType, Status) "
            "VALUES (%s, %s, %s, %s, %s)"
        )
        values = (i, record_id, fine_amount, violation_type, status)
        cursor.execute(query, values)

    conn.commit()


# 执行插入数据
insert_violationrecord_data()

# 关闭数据库连接
cursor.close()
conn.close()

print("100条数据已成功插入到violationrecord表中。")