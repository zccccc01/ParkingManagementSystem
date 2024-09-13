userRepo := repository.NewUserRepository(db)
//已存在的userID
user1 := models.User{
	UserID:   1,
	UserName: "LZZ",
	Password: "123456",
	Tel:      "123456789",
}
//不存在的userID
user2 := models.User{
	UserID:   111,
	UserName: "ZCC",
	Password: "123456",
	Tel:      "133488789",
}
err1 := userRepo.Create(&user1)
err2 := userRepo.Create(&user2)
str1, err3 := userRepo.GetTelByID(1)
str2, err4 := userRepo.GetTelByID(110)
err5 := userRepo.UpdatePasswordByID(1, "987456")
err6 := userRepo.UpdateTelByID(110, "123456789")
err7 := userRepo.Delete(1)
err8 := userRepo.Delete(110)
if err1 != nil {
	log.Fatalf("failed to create user: %v", err1) // 使用日志记录错误，而不是panic
}
if err2 != nil {
	log.Fatalf("failed to create user: %v", err2) // 使用日志记录错误，而不是panic
}
if err3 != nil {
	log.Fatalf("failed to get user: %v", err3) // 使用日志记录错误，而不是panic
}
if err4 != nil {
	log.Fatalf("failed to get user: %v", err4) // 使用日志记录错误，而不是panic
}
if err5 != nil {
	log.Fatalf("failed to update user: %v", err5) // 使用日志记录错误，而不是panic
}
if err6 != nil {
	log.Fatalf("failed to update user: %v", err6) // 使用日志记录错误，而不是panic
}
if err7 != nil {
	log.Fatalf("failed to delete user: %v", err7) // 使用日志记录错误，而不是panic
}
if err8 != nil {
	log.Fatalf("failed to delete user: %v", err8) // 使用日志记录错误，而不是panic
}
log.Printf("user1 tel: %s", str1)
log.Printf("user2 tel: %s", str2)