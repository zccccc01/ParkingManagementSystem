import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom'; // 导入 useNavigate 钩子
import Header from '../components/Header';
import Footer from '../components/Footer';
import './LoginPage.scss';

const LoginPage = () => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const navigate = useNavigate(); // 获取导航函数

  const handleSubmit = async (event) => {
    event.preventDefault(); // 阻止表单默认提交行为

    // 示例逻辑：直接模拟登录成功
    localStorage.setItem('token', 'dummy-token'); // 存储 dummy token 到本地
    navigate('/dashboard'); // 登录成功后跳转到 dashboard 页面
  };

  return (
    <div className="login-page">
      <Header />
      <h1>用户登录</h1>
      <form onSubmit={handleSubmit}>
        <label htmlFor="username">
          用户名：
          <input
            type="username"
            id="username"
            name="username"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
          />
        </label>
        <br />
        <br />
        <label htmlFor="password">
          密码：
          <input
            type="password"
            id="password"
            name="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
          />
        </label>
        <br />
        <br />
        <button type="submit">登录</button>
      </form>
      <Footer />
    </div>
  );
};

export default LoginPage;
