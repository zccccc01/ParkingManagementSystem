import axios from 'axios';

const apiBaseUrl = 'http://localhost:8000/api';

// 注册函数
const register = async (userID, tel, password, confirmPassword) => {
  try {
    const response = await axios.post(`${apiBaseUrl}/register`, {
      userID,
      tel,
      password,
      confirmPassword,
    });
    return response.data;
  } catch (error) {
    throw new Error(error.response.data.message || '注册失败，请检查输入的信息');
  }
};

// 登录函数
const login = async (username, password) => {
  try {
    const response = await axios.post(`${apiBaseUrl}/auth/login`, {
      username,
      password,
    });
    return response.data;
  } catch (error) {
    throw new Error('登录失败，请检查用户名和密码');
  }
};

export { register, login }; // 使用命名导出
