const express = require('express');
const bodyParser = require('body-parser');
const cors = require('cors');

const app = express();

// Middleware
app.use(bodyParser.json());
app.use(cors());

// Mock user data
const users = [
  { userID: 'user1', password: 'password1' },
  { userID: 'user2', password: 'password2' }
];

// Login endpoint
app.post('/login', (req, res) => {
  const { userID, password } = req.body;

  const user = users.find(u => u.userID === userID && u.password === password);

  if (user) {
    res.status(200).json({ message: '登录成功' });
  } else {
    res.status(401).json({ message: '登录失败，请检查用户名和密码' });
  }
});

// Start the server
const port = process.env.PORT || 8080;
app.listen(port, () => {
  console.log(`Server is running on port ${port}`);
});