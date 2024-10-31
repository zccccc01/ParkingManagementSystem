// pages/AdminDashboard.js
import React from 'react';
import Header from '../components/Header';
import Footer from '../components/Footer';

const AdminDashboard = () => {
  return (
    <div className="admin-dashboard">
      <Header />
      <h1>管理员仪表盘</h1>
      <p>提供停车场管理功能和数据分析</p>
      <Footer />
    </div>
  );
};

export default AdminDashboard;
