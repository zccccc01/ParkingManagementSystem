package controllers

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/zccccc01/ParkingManagementSystem/backend/internal/models"
	"github.com/zccccc01/ParkingManagementSystem/backend/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

const SecretKey = "secret"

type UserController struct {
	UserRepo repository.UserRepository
}

func NewUserController(repo repository.UserRepository) *UserController {
	return &UserController{UserRepo: repo}
}

func (uc *UserController) Register(c *fiber.Ctx) error {
	var data struct {
		ID       int    `json:"id"`
		Tel      string `json:"tel"`
		Password string `json:"password"`
	}

	// 解析请求体
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), 14)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error encrypting password",
		})
	}

	// 创建新用户，用户名设置为与ID相同
	user := models.User{
		UserID:   data.ID,
		Tel:      data.Tel,
		Password: string(hashedPassword),
		UserName: strconv.Itoa(data.ID), // 将ID转换为字符串作为用户名
	}

	// 保存用户到数据库
	if _, err := uc.UserRepo.Create(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error creating user",
		})
	}

	// 返回成功响应，不返回密码
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"id":   user.UserID,
		"tel":  user.Tel,
		"name": user.UserName,
	})
}

func (uc *UserController) Login(c *fiber.Ctx) error {
	var data models.User
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	user, err := uc.UserRepo.FindUserByTel(data.Tel)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "user not found"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password)); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "incorrect password"})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    strconv.Itoa(user.UserID),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
	})

	token, err := claims.SignedString([]byte(SecretKey))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "could not log in"})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	})

	return c.JSON(fiber.Map{"message": "success"})
}

func (uc *UserController) AuthenticatedUser(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	token, err := jwt.ParseWithClaims(cookie, &jwt.RegisteredClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "unauthorized"})
	}

	claims := token.Claims.(*jwt.RegisteredClaims)

	// 将 claims.Issuer 转换为 int 类型
	userID, err := strconv.Atoi(claims.Issuer)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid token"})
	}

	user, err := uc.UserRepo.FindUserByID(userID)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "unauthorized"})
	}

	return c.JSON(fiber.Map{
		"id":   user.UserID,
		"tel":  user.Tel,
		"name": user.UserName,
	})
}

func (uc *UserController) Logout(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	})

	return c.JSON(fiber.Map{"message": "success"})
}
