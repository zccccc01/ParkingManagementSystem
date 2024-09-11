/* eslint-disable jsx-a11y/label-has-associated-control */
/* eslint-disable no-console */

import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom'; // 导入 useNavigate 钩子
import Header from '../components/Header';
import Footer from '../components/Footer';

const LoginPage = () => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const navigate = useNavigate(); // 获取导航函数

  const handleSubmit = async (event) => {
    event.preventDefault(); // 阻止表单默认提交行为

    // 示例逻辑：直接模拟登录成功
    console.log('Logging in with username:', username, 'and password:', password);
    localStorage.setItem('token', 'dummy-token'); // 存储 dummy token 到本地
    navigate('/dashboard'); // 登录成功后跳转到 dashboard 页面
  };

  return (
    <div className="login-page">
      <Header />
      <h1>用户登录</h1>
      <form onSubmit={handleSubmit}>
        <label>用户名：</label>
        <input
          type="text"
          id="username"
          name="username"
          value={username}
          onChange={(e) => setUsername(e.target.value)}
          required
        />
        <br />
        <label>密码：</label>
        <input
          type="password"
          id="password"
          name="password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          required
        />
        <br />
        <button type="submit">登录</button>
      </form>
      <Footer />
    </div>
  );
};

export default LoginPage;
