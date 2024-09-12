// pages/ParkingLotListPage.js
import React from 'react';
import Header from '../components/Header';
import Footer from '../components/Footer';
import './ParkingLotListPage.scss';

const ParkingLotListPage = () => {
  return (
    <div className="parking-lot-list-page">
      <Header />
      <h1>停车场列表</h1>
      <p>附近停车场的实时信息</p>
      <Footer />
    </div>
  );
};

export default ParkingLotListPage;
