// models/User.js
import { DataTypes } from 'sequelize';
import sequelize from './db'; // 导入数据库连接

const User = sequelize.define('User', {
  username: {
    type: DataTypes.STRING,
    allowNull: false,
    unique: true
  },
  password: {
    type: DataTypes.STRING,
    allowNull: false
  },
  role: {
    type: DataTypes.ENUM('user', 'admin'),
    defaultValue: 'user'
  }
}, {
  tableName: 'users',
  timestamps: false
});

// 创建管理员账号
(async () => {
  try {
    const adminUser = await User.findOne({ where: { username: 'root' } });
    if (!adminUser) {
      await User.create({
        username: 'root',
        password: 'root', // 假设这里使用了加密后的密码
        role: 'admin'
      });
      console.log('Admin user created successfully.');
    } else {
      console.log('Admin user already exists.');
    }
  } catch (error) {
    console.error('Error creating admin user:', error);
  }
})();

export default User;