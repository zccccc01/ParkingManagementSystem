const express = require('express');
const bodyParser = require('body-parser');
const cors = require('cors');

const app = express();

// Middleware
app.use(bodyParser.json());
app.use(cors());

// Mock user data storage
const users = [];

// Registration endpoint
app.post('/register', (req, res) => {
  const { userID, tel, password, confirmPassword } = req.body;

  // Basic validation
  if (!userID || !tel || !password || !confirmPassword) {
    return res.status(400).json({ message: '所有字段都是必填项' });
  }

  if (password !== confirmPassword) {
    return res.status(400).json({ message: '密码不一致' });
  }

  // Check if user already exists
  const existingUser = users.find(user => user.userID === userID);
  if (existingUser) {
    return res.status(409).json({ message: '用户ID已存在' });
  }

  // Add new user
  users.push({ userID, tel, password });

  // Simulate successful registration
  res.status(201).json({ message: '注册成功' });
});

// Start the server
const port = process.env.PORT || 8080;
app.listen(port, () => {
  console.log(`Server is running on port ${port}`);
});