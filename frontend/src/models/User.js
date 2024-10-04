import express from 'express';
import mysql from 'mysql2/promise';

const app = express();
app.use(express.json());
app.use(express.urlencoded({ extended: false }));

// 数据库配置
const pool = mysql.createPool({
  host: 'localhost',
  user: 'root',
  password: '123456',
  database: 'username',
  waitForConnections: true,
  connectionLimit: 10,
  queueLimit: 0
});

// 定义异步查询函数
const query = async (sql, params = []) => {
  const [rows] = await pool.execute(sql, params);
  return rows;
};

// 验证数据库连接并释放连接
const verifyDatabaseConnection = async () => {
  try {
    const connection = await pool.getConnection();
    console.log('Database connection has been established successfully.');
    connection.release();
    return true;
  } catch (error) {
    console.error('Error connecting to the database:', error);
    return false;
  }
};

// 创建管理员账号
const createAdminUserIfNotExists = async () => {
  try {
    if (!(await verifyDatabaseConnection())) {
      return;
    }

    // 查询管理员账号是否存在
    const results = await query('SELECT * FROM users WHERE username = ?', ['root']);
    if (results.length === 0) {
      // 创建管理员账号
      await query('INSERT INTO users (username, password, role) VALUES (?, ?, ?)', ['root', 'root', 'admin']);
      console.log('Admin user created successfully.');
    } else {
      console.log('Admin user already exists.');
    }
  } catch (error) {
    console.error('Error creating admin user:', error);
  }
};

// 执行创建管理员账号的操作
(async () => {
  await createAdminUserIfNotExists();
})();

// 处理登录请求
app.get('/login', async (req, res) => {
  try {
    const user_name = req.query.user;
    const password = req.query.password;

    const sql = 'SELECT * FROM usermessage WHERE username = ?';
    const sqlParams = [user_name];

    const result = await query(sql, sqlParams);

    let response;
    if (result.length === 0) {
      response = {
        code: 101,
        mag: '用户不存在'
      };
    } else if (result[0].password === password) {
      response = {
        code: 200,
        mag: '登录成功',
        data: {
          userId: result[0].userId
        }
      };
    } else {
      response = {
        code: 400,
        mag: '密码错误'
      };
    }
    res.send(response);
  } catch (error) {
    res.status(500).send(error.message);
  }
});

// 处理获取分页数据请求
app.post('/getpage', async (req, res) => {
  try {
    const { name, content, pagesize = 10, pagenum = 1 } = req.body;
    const pagestart = (pagenum - 1) * pagesize;

    const sql = `SELECT * FROM casedata WHERE case_name LIKE BINARY '%${name || ""}%' AND case_content LIKE BINARY '%${content || ""}%' LIMIT ${pagestart}, ${pagesize}`;

    const result = await query(sql);

    let response;
    if (result.length === 0) {
      response = {
        code: 303,
        mag: '暂无数据'
      };
    } else {
      response = {
        code: 200,
        mag: {
          data: result,
          total: result.length,
          pagesize: pagesize,
          pagenum: pagenum
        }
      };
    }
    res.send(response);
  } catch (error) {
    res.status(500).send(error.message);
  }
});

// 处理获取用户信息请求
app.get('/api/user', async (req, res) => {
  try {
    const userId = req.cookies.userId; // 假设通过cookie传递用户ID

    const sql = 'SELECT * FROM users WHERE id = ?';
    const sqlParams = [userId];

    const result = await query(sql, sqlParams);

    if (result.length === 0) {
      res.status(404).json({ message: 'User not found' });
    } else {
      res.json(result[0]);
    }
  } catch (error) {
    res.status(500).send(error.message);
  }
});

// 处理登出请求
app.post('/api/logout', async (req, res) => {
  try {
    // 清除cookie中的用户信息
    res.clearCookie('userId');
    res.json({ message: 'Logout successful' });
  } catch (error) {
    res.status(500).send(error.message);
  }
});

// 启动服务器
const server = app.listen(8000, () => {
  const port = server.address().port;
  console.log(`Server is running on port ${port}`);
});

export default {};
