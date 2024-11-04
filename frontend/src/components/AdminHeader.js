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
      <h1>Parking Management System</h1>
      <br />
      <br />
    </header>
  );
};

export default Header;
