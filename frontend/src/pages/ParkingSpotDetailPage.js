// pages/ParkingSpotDetailPage.js
import React from 'react';
import Header from '../components/Header';
import Footer from '../components/Footer';

const ParkingSpotDetailPage = ({ match }) => {
  const { id } = match.params;
  return (
    <div className="parking-spot-detail-page">
      <Header />
      <h1>停车位详情：{id}</h1>
      <p>单个停车位的详细信息和预定选项</p>
      <Footer />
    </div>
  );
};

export default ParkingSpotDetailPage;
