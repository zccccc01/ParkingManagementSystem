import React, { useState, useEffect } from 'react';
import SearchBar from '../components/SearchBar';
import Header from '../components/Header';
import Footer from '../components/Footer';
import './HomePage.module.scss'; // 导入SCSS样式

const HomePage = () => {
  const [currentImage, setCurrentImage] = useState(0);
  const images = [
    '//p3.dcarimg.com/img/tos-cn-i-dcdx/5cd480f6fcc24ad3a7353de566a7229e~360x0.webp',
    '//p3.dcarimg.com/img/tos-cn-i-dcdx/7a507b6ec14c4ead9c95c336893aa9d2~360x0.webp',
    '//p3.dcarimg.com/img/motor-mis-img/f8250b5672059b6380b04eeeb0d7fe33~360x0.webp',
    '//p3.dcarimg.com/img/tos-cn-i-dcdx/8c6e1943c2094ec2a14fd460d2c0ba4d~360x0.webp',
  ];

  const [userInfo, setUserInfo] = useState(null);
  const [logoutMessage, setLogoutMessage] = useState('');

  useEffect(() => {
    const interval = setInterval(() => {
      setCurrentImage(currentImage === images.length - 1 ? 0 : currentImage + 1);
    }, 3000);

    return () => clearInterval(interval); // 清除定时器，防止内存泄漏
  }, [currentImage, images.length]); // 添加 images.length 到依赖数组

  const fetchUserInfo = async () => {
    try {
      const response = await fetch('http://localhost:8000/api/user', {
        method: 'GET',
        credentials: 'include', // 确保请求带上 cookie
      });

      if (response.ok) {
        const data = await response.json();
        setUserInfo(`User ID: ${data.id}, Phone: ${data.tel}, Name: ${data.name}`);
      } else {
        setUserInfo('Failed to retrieve user info');
      }
    } catch (error) {
      setUserInfo('An error occurred');
    }
  };

  const handleLogout = async () => {
    try {
      const response = await fetch('http://localhost:8000/api/logout', {
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

  return (
    <div className="HomePage">
      <Header />
      <SearchBar />
      <p>欢迎来到停车管理系统！</p>
      <div>
        <img id="image-slider" width="540" height="340" src={images[currentImage]} alt="slider" />
      </div>
      <hr />
      <button type="button" onClick={fetchUserInfo}>
        获取用户信息
      </button>
      <p id="userInfo">{userInfo}</p>
      <button type="button" onClick={handleLogout}>
        登出
      </button>
      <p id="logoutMessage">{logoutMessage}</p>
      <Footer />
    </div>
  );
};

export default HomePage;
