// pages/ParkingRecordPage.js
import React from 'react';
import Header from '../components/Header';
import Footer from '../components/Footer';
import './ParkingRecordPage.scss';

const ParkingRecordPage = () => {
  return (
    <div className="parking-record-page">
      <Header />
      <h1>停车记录页面</h1>
      <p>展示用户的停车历史和支付记录</p>
      <Footer />
    </div>
  );
};

export default ParkingRecordPage;
