// server.js
const express = require("express");
const bcrypt = require("bcryptjs");
const jwt = require("jsonwebtoken");
const User = require("./models/User");

const app = express();
app.use(express.json());

// 统一错误处理函数
const handleServerError = (error, res) => {
  console.error(error);
  res.status(500).json({ message: "Server error" });
};

// 注册API
app.post("/api/user/register", async (req, res) => {
  const { userId, password, tel } = req.body;

  try {
    // 检查用户是否已存在
    const existingUser = await User.findOne({ where: { userId } });
    if (existingUser) {
      return res.status(400).json({ message: "Username already exists" });
    }

    // 加密密码并创建新用户
    const hashedPassword = await bcrypt.hash(password, 10);
    const newUser = await User.create({
      userId,
      password: hashedPassword,
      tel,
    });

    res.status(201).json({ message: "User registered successfully" });
  } catch (error) {
    handleServerError(error, res);
  }
});

// 登录API
app.post("/api/user/login", async (req, res) => {
  const { username, password } = req.body;

  try {
    // 检查用户是否存在
    const user = await User.findOne({ where: { username } });
    if (!user) {
      return res.status(400).json({ message: "Invalid credentials" });
    }

    // 验证密码
    const passwordMatch = await bcrypt.compare(password, user.password);
    if (!passwordMatch) {
      return res.status(400).json({ message: "Invalid credentials" });
    }

    // 生成并返回 JWT 令牌
    const token = jwt.sign({ userId: user.id, role: user.role }, "secretKey", {
      expiresIn: "1h",
    });
    res.json({ token });
  } catch (error) {
    handleServerError(error, res);
  }
});

app.listen(3000, () => {
  console.log("Server started on port 3000");
});
