import React, { useState, useEffect } from 'react';
import styles from './Header.module.scss';
import Sidebar from './Sidebar.tsx';

const Header = () => {
  const [isLoggedIn, setIsLoggedIn] = useState(false);

  useEffect(() => {
    const checkLoginStatus = async () => {
      try {
        const response = await fetch('http://localhost:8000/api/user', {
          method: 'GET',
          credentials: 'include', // 确保请求带上 cookie
        });
        if (response.ok) {
          setIsLoggedIn(true);
        } else {
          setIsLoggedIn(false);
        }
      } catch (error) {
        console.error('Error checking login status:', error);
        setIsLoggedIn(false);
      }
    };

    checkLoginStatus();
  }, []);

  return (
    <header className={styles.header}>
      <Sidebar />
      <div className={styles.headerContent}>
        <img id="1" src="/logo1.jpg" alt="logo" />
        <h1>Parking Management System</h1>
      </div>
      <br />
      <br />
      <div>
        <ul>
          {!isLoggedIn && (
            <>
              <li>
                <a href="/register">注册</a>
              </li>
              <li>
                <a href="/login">登录</a>
              </li>
            </>
          )}
        </ul>
      </div>
    </header>
  );
};

export default Header;
