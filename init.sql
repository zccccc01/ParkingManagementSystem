-- 创建 users 表
CREATE TABLE `users` (
  `UserID` int NOT NULL,
  `Username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `Password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `Tel` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  PRIMARY KEY (`UserID`),
  UNIQUE KEY `Tel` (`Tel`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- 创建 vehicle 表
CREATE TABLE `vehicle` (
  `VehicleID` int NOT NULL,
  `UserID` int DEFAULT NULL,
  `PlateNumber` varchar(255) DEFAULT NULL,
  `Color` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`VehicleID`),
  UNIQUE KEY `PN` (`PlateNumber`) USING BTREE,
  KEY `UID` (`UserID`),
  CONSTRAINT `UID` FOREIGN KEY (`UserID`) REFERENCES `users` (`UserID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- 创建 parkinglot 表
CREATE TABLE `parkinglot` (
  `ParkingLotID` int NOT NULL,
  `ParkingName` varchar(255) NOT NULL,
  `Longitude` decimal(9,6) NOT NULL,
  `Latitude` decimal(9,6) NOT NULL,
  `Capacity` int DEFAULT NULL,
  `Rates` decimal(10,2) NOT NULL,
  PRIMARY KEY (`ParkingLotID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- 创建 parkingspace 表
CREATE TABLE `parkingspace` (
  `SpaceID` int NOT NULL,
  `Status` enum('FREE','OCCUPIED','RESERVED') CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `ParkingLotID` int NOT NULL,
  PRIMARY KEY (`SpaceID`),
  KEY `PLID` (`ParkingLotID`),
  CONSTRAINT `PLID` FOREIGN KEY (`ParkingLotID`) REFERENCES `parkinglot` (`ParkingLotID`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- 创建 parkingrecord 表
CREATE TABLE `parkingrecord` (
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
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- 创建 reservation 表
CREATE TABLE `reservation` (
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
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- 创建 violationrecord 表
CREATE TABLE `violationrecord` (
  `ViolationID` int NOT NULL,
  `RecordID` int DEFAULT NULL,
  `FineAmount` double DEFAULT NULL,
  `ViolationType` enum('OVERSTAY','NOPAY') CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `Status` enum('PAID','UNPAID','DISPUTED') DEFAULT NULL,
  PRIMARY KEY (`ViolationID`),
  KEY `REID` (`RecordID`),
  CONSTRAINT `REID` FOREIGN KEY (`RecordID`) REFERENCES `parkingrecord` (`RecordID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- 创建 paymentrecord 表
CREATE TABLE `paymentrecord` (
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
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
