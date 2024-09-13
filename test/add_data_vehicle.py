import mysql.connector
import random
import string

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

# 定义颜色列表
colors = ['Red', 'Blue', 'Green', 'Black', 'White', 'Silver', 'Gray', 'Yellow', 'Orange', 'Pink']


# 插入数据的函数
def insert_vehicle_data():
    # 获取users表中存在的UserID列表
    cursor.execute("SELECT UserID FROM users")
    user_ids = [row[0] for row in cursor.fetchall()]

    for i in range(1, 101):
        user_id = random.choice(user_ids) if user_ids else None  # 随机选择一个存在的UserID
        plate_number = ''.join(random.choices(string.ascii_uppercase + string.digits, k=7))  # 随机生成车牌号
        color = random.choice(colors)  # 随机选择颜色

        query = (
            "INSERT INTO vehicle (VehicleID, UserID, PlateNumber, Color) "
            "VALUES (%s, %s, %s, %s)"
        )
        values = (i, user_id, plate_number, color)
        cursor.execute(query, values)

    conn.commit()


# 执行插入数据
insert_vehicle_data()

# 关闭数据库连接
cursor.close()
conn.close()

print("100条数据已成功插入到vehicle表中。")