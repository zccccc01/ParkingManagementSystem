// pages/AdminDashboard.js
import React from 'react';
import Header from '../components/AdminHeader';
import Footer from '../components/Footer';
import styles from './AdminDashboard.scss';

const AdminDashboard = () => {
  return (
    <div className="admin-dashboard">
      <Header />
      <h1>管理员仪表盘</h1>
      <p>提供停车场管理功能和数据分析</p>
      <div>
        <img
          src="https://github.com/egonelbre/gophers/blob/master/.thumb/animation/gopher-dance-long-3x.gif?raw=true"
          className={styles.logo}
          alt="gopher dance"
        />
      </div>
      <Footer />
    </div>
  );
};

export default AdminDashboard;
