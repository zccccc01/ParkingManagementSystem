// pages/UserPage.js
// TODO: 用户上传头像，数据库保留头像路径
// TODO: 设置一个默认头像，防止用户没有上传头像时显示空白
import React, { useState, useEffect } from 'react';
import Header from '../components/Header';
import Footer from '../components/Footer';
import './UserPage.scss';

const User = () => {
  const [userInfo, setUserInfo] = useState(null);
  const [logoutMessage, setLogoutMessage] = useState('');

  const fetchUserInfo = async () => {
    try {
      const response = await fetch('http://localhost:8000/api/user', {
        method: 'GET',
        credentials: 'include', // 确保请求带上 cookie
      });

      if (response.ok) {
        const data = await response.json();
        setUserInfo(`User ID: ${data.id}\nPhone: ${data.tel}\nName: ${data.name}`);
      } else {
        setUserInfo('Failed to retrieve user info');
      }
    } catch (error) {
      setUserInfo('An error occurred');
    }
  };

  const handleLogout = async () => {
    try {
      const response = await fetch('http://localhost:8000/api/user/logout', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        credentials: 'include', // 发送带有cookie的请求
      });

      if (response.ok) {
        setLogoutMessage('Logout successful');
        window.location.href = ''; // 登出后返回主页
      } else {
        setLogoutMessage('Logout failed');
      }
    } catch (error) {
      setLogoutMessage('An error occurred');
    }
  };

  useEffect(() => {
    // 组件加载时自动获取用户信息
    fetchUserInfo();
  }, []);

  return (
    <div className="user-page">
      <Header />
      <h1>用户页面</h1>
      <p id="userInfo">{userInfo}</p>
      <button type="button" onClick={handleLogout}>
        登出
      </button>
      <p id="logoutMessage">{logoutMessage}</p>
      <Footer />
    </div>
  );
};

export default User;
