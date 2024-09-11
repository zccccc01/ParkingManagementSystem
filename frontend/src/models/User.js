import mysql from 'mysql2/promise';

// 数据库配置
const pool = mysql.createPool({
  host: 'localhost',
  user: 'username',
  password: 'password',
  database: 'database'
});

// 定义异步查询函数
const query = async (sql, params = []) => {
  const [rows] = await pool.execute(sql, params);
  return rows;
};

// 创建管理员账号
(async () => {
  try {
    // 验证数据库连接
    const connection = await pool.getConnection();
    console.log('Database connection has been established successfully.');
    connection.release();

    // 查询管理员账号是否存在
    const [results] = await query('SELECT * FROM users WHERE username = ?', ['root']);
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
})();

export default {};