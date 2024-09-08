	// 创建ParkingLot对象
	newLot := models.ParkingLot{
		ParkingLotID: 3,
		ParkingName:  "ZCST Parking",
		Longitude:    decimal.RequireFromString("11.00"),
		Latitude:     decimal.RequireFromString("75.000"),
		Capacity:     100,
		Rates:        decimal.RequireFromString("50"),
	}
	result := parkingLotRepo.Create(&newLot)
	// 用result判断是否创建成功
	if result != nil {
		log.Fatalf("failed to create record: %v", result)
	}
	// 打印创建的记录
	log.Printf("Created parking lot: %+v", newLot)
	/*
		-----------------------------------------------------
	*/
	// 查找所有记录
	newParkingLot, err := parkingLotRepo.FindAll()
	if err == nil {
		log.Printf("Found parking lot: %+v", newParkingLot)
	}
	log.Fatalf("failed to find parking lot: %v", err)
	/*
		-----------------------------------------------------
	*/
	// 按ID查找记录
	id := 2
	parkingLot, err := parkingLotRepo.FindByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Println("Parking lot not found")
		} else {
			log.Fatalf("failed to find parking lot: %v", err)
		}
	} else {
		log.Printf("Found parking lot: %+v", parkingLot)
	}
	/*
		-----------------------------------------------------
	*/
	// 按Name查找记录
	name := "ZCST Parking"
	parkingLot, err := parkingLotRepo.FindByName(name)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Println("Parking lot not found")
		} else {
			log.Fatalf("failed to find parking lot: %v", err)
		}
	} else {
		log.Printf("Found parking lot: %+v", parkingLot)
	}
	/*
		-----------------------------------------------------
	*/
	// 删除记录
	result := parkingLotRepo.Delete(2)
	if result != nil {
		log.Fatalf("failed to delete record: %v", result)
	}
	log.Println("Record deleted successfully")
	/*
		-----------------------------------------------------
	*/
	// 更新记录
	// 创建一个 ParkingLot 实例并设置字段
	lot := &models.ParkingLot{
		ParkingLotID: 1,
		ParkingName:  "Central Parking Updated",
		Longitude:    decimal.RequireFromString("13"),
		Latitude:     decimal.RequireFromString("65"),
		Capacity:     150,
		Rates:        decimal.RequireFromString("90"),
	}
	// 调用 Update 方法更新记录
	id := 1
	parkingLotRepo.Update(lot, id)
	res, err1 := parkingLotRepo.FindByID(id)
	if err1 != nil {
		log.Fatalf("failed to find parking lot: %v", err1)
	}
	log.Printf("Updated parking lot: %+v", res)