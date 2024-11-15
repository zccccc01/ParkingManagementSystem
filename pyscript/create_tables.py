import mysql.connector
import logging

# 配置日志
logging.basicConfig(level=logging.INFO, format='%(asctime)s - %(levelname)s - %(message)s')

# 数据库连接配置
config = {
    'host': '10.1.0.20',
    'user': 'root',
    'password': '123456',
    'database': 'mydb'  # 请替换为您的数据库名称
}

# 连接到数据库并记录日志
try:
    logging.info("正在连接到数据库...")
    conn = mysql.connector.connect(**config)
    cursor = conn.cursor()
    logging.info("成功连接到数据库！")

    # 确保使用正确的数据库
    cursor.execute("USE mydb;")
    logging.info("数据库已切换到 mydb。")

    # 创建表的SQL语句
    create_tables_sql = """
    -- 创建 users 表
    CREATE TABLE IF NOT EXISTS `users` (
        `UserID` int NOT NULL,
        `Username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
        `Password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
        `Tel` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
        PRIMARY KEY (`UserID`),
        UNIQUE KEY `Tel` (`Tel`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

    -- 创建 vehicle 表
    CREATE TABLE IF NOT EXISTS `vehicle` (
        `VehicleID` int NOT NULL,
        `UserID` int DEFAULT NULL,
        `PlateNumber` varchar(255) DEFAULT NULL,
        `Color` varchar(255) DEFAULT NULL,
        PRIMARY KEY (`VehicleID`),
        UNIQUE KEY `PN` (`PlateNumber`) USING BTREE,
        KEY `UID` (`UserID`),
        CONSTRAINT `UID` FOREIGN KEY (`UserID`) REFERENCES `users` (`UserID`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

    -- 创建 parkinglot 表
    CREATE TABLE IF NOT EXISTS `parkinglot` (
        `ParkingLotID` int NOT NULL,
        `ParkingName` varchar(255) NOT NULL,
        `Longitude` decimal(9,6) NOT NULL,
        `Latitude` decimal(9,6) NOT NULL,
        `Capacity` int DEFAULT NULL,
        `Rates` decimal(10,2) NOT NULL,
        PRIMARY KEY (`ParkingLotID`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

    -- 创建 parkingspace 表
    -- 创建 parkingspace 表
    CREATE TABLE IF NOT EXISTS `parkingspace` (
    `SpaceID` int NOT NULL,
    `Status` enum('FREE','OCCUPIED','RESERVED') NOT NULL,
    `ParkingLotID` int NOT NULL,
    PRIMARY KEY (`SpaceID`),
    KEY `PLID` (`ParkingLotID`),
    CONSTRAINT `PLID` FOREIGN KEY (`ParkingLotID`) REFERENCES `parkinglot` (`ParkingLotID`) ON DELETE RESTRICT ON UPDATE RESTRICT
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


    -- 创建 parkingrecord 表
    CREATE TABLE IF NOT EXISTS `parkingrecord` (
        `RecordID` int NOT NULL,
        `SpaceID` int DEFAULT NULL,
        `VehicleID` int DEFAULT NULL,
        `LotID` int DEFAULT NULL,
        `StartTime` datetime DEFAULT NULL,
        `EndTime` datetime DEFAULT NULL,
        `Fee` double DEFAULT NULL,
        PRIMARY KEY (`RecordID`),
        KEY `PSID` (`SpaceID`),
        KEY `VEID` (`VehicleID`),
        KEY `PALID` (`LotID`),
        CONSTRAINT `PALID` FOREIGN KEY (`LotID`) REFERENCES `parkinglot` (`ParkingLotID`),
        CONSTRAINT `PSID` FOREIGN KEY (`SpaceID`) REFERENCES `parkingspace` (`SpaceID`),
        CONSTRAINT `VEID` FOREIGN KEY (`VehicleID`) REFERENCES `vehicle` (`VehicleID`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

    -- 创建 reservation 表
    CREATE TABLE IF NOT EXISTS `reservation` (
        `ReservationID` int NOT NULL,
        `StartTime` datetime DEFAULT NULL,
        `EndTime` datetime DEFAULT NULL,
        `SpaceID` int DEFAULT NULL,
        `VehicleID` int DEFAULT NULL,
        `LotID` int DEFAULT NULL,
        `Status` enum('Completed','Cancelled','Doing') DEFAULT NULL,
        PRIMARY KEY (`ReservationID`),
        KEY `SID` (`SpaceID`),
        KEY `VID` (`VehicleID`),
        KEY `LID` (`LotID`),
        CONSTRAINT `LID` FOREIGN KEY (`LotID`) REFERENCES `parkinglot` (`ParkingLotID`),
        CONSTRAINT `SID` FOREIGN KEY (`SpaceID`) REFERENCES `parkingspace` (`SpaceID`),
        CONSTRAINT `VID` FOREIGN KEY (`VehicleID`) REFERENCES `vehicle` (`VehicleID`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

    -- 创建 violationrecord 表
    CREATE TABLE IF NOT EXISTS `violationrecord` (
        `ViolationID` int NOT NULL,
        `RecordID` int DEFAULT NULL,
        `FineAmount` double DEFAULT NULL,
        `ViolationType` enum('OVERSTAY','NOPAY') CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
        `Status` enum('PAID','UNPAID','DISPUTED') DEFAULT NULL,
        PRIMARY KEY (`ViolationID`),
        KEY `REID` (`RecordID`),
        CONSTRAINT `REID` FOREIGN KEY (`RecordID`) REFERENCES `parkingrecord` (`RecordID`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

    -- 创建 paymentrecord 表
    CREATE TABLE IF NOT EXISTS `paymentrecord` (
        `PaymentID` int NOT NULL,
        `RecordID` int DEFAULT NULL,
        `ReservationID` int DEFAULT NULL,
        `Amount` double DEFAULT NULL,
        `PaymentTimestamp` datetime DEFAULT NULL,
        `PaymentMethod` varchar(255) DEFAULT NULL,
        PRIMARY KEY (`PaymentID`),
        KEY `RID` (`ReservationID`),
        KEY `RECID` (`RecordID`),
        CONSTRAINT `RECID` FOREIGN KEY (`RecordID`) REFERENCES `parkingrecord` (`RecordID`),
        CONSTRAINT `RID` FOREIGN KEY (`ReservationID`) REFERENCES `reservation` (`ReservationID`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
    """

    # 执行创建表操作并记录日志
    logging.info("正在创建表...")
    conn.start_transaction()  # 开始事务
    try:
        for statement in create_tables_sql.split(';'):
            statement = statement.strip()
            if statement:
                try:
                    cursor.execute(statement)
                    logging.info(f"成功执行: {statement[:30]}...")
                except mysql.connector.Error as err:
                    logging.error(f"创建表时出错: {err} -- SQL: {statement[:30]}...")
                    conn.rollback()  # 如果出错，回滚事务
                    break
        conn.commit()  # 提交事务
        logging.info("所有表创建成功！")
    except mysql.connector.Error as err:
        logging.error(f"创建表时出错: {err}")
        conn.rollback()  # 如果发生错误，回滚事务
    finally:
        cursor.close()
        conn.close()
        logging.info("数据库连接已关闭。")

except mysql.connector.Error as err:
    logging.error(f"连接数据库时出错: {err}")
