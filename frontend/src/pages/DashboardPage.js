// pages/DashboardPage.js
import React from 'react';
import Header from '../components/Header';
import Footer from '../components/Footer';

const DashboardPage = () => {
  return (
    <div className="dashboard-page">
      <Header />
      <h1>用户仪表盘</h1>
      <p>车辆信息、预定的停车位和停车历史</p>
      <Footer />
    </div>
  );
};

export default DashboardPage;
