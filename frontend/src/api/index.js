import axios from 'axios';

const apiBaseUrl = 'http://your-backend-api-url.com/api';

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

export default login; // 使用默认导出
