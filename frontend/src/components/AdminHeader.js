import React, { useEffect } from 'react';
import styles from './Header.module.scss';
import AdminSidebar from './AdminSidebar.tsx';

const Header = () => {
  useEffect(() => {
    const checkLoginStatus = async () => {
      try {
        const response = await fetch('http://localhost:8000/api/user', {
          method: 'GET',
          credentials: 'include', // 确保请求带上 cookie
        });
        if (!response.ok) {
          console.error('User is not logged in');
        }
      } catch (error) {
        console.error('Error checking login status:', error);
      }
    };

    checkLoginStatus();
  }, []);

  return (
    <header className={styles.header}>
      <AdminSidebar />
      <div>
        <h1>Parking Management System</h1>
        <img
          src="https://cdn.buymeacoffee.com/uploads/slider_images/2018/05/98302662377036c7092e36f690ab4069.gif"
          className={styles.logo}
          alt="logo"
        />
      </div>
      <br />
      <br />
      <br />
      <br />
    </header>
  );
};

export default Header;
