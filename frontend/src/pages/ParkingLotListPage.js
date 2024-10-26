import React, { useState, useEffect } from 'react';
import Header from '../components/Header';
import Footer from '../components/Footer';
import './ParkingLotListPage.scss';

const ParkingLotListPage = () => {
  const [parkingLots, setParkingLots] = useState([]);

  const fetchParkingLots = async () => {
    try {
      const response = await fetch('http://localhost:8000/api/parkinglot');
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      const data = await response.json();
      setParkingLots(data);
    } catch (error) {
      console.error('Failed to fetch parking lots:', error);
    }
  };

  useEffect(() => {
    // 在组件加载时自动获取停车场信息
    fetchParkingLots();
  }, []);

  return (
    <div className="parking-lot-list-page">
      <Header />
      <h1>停车场列表</h1>
      <table>
        <thead>
          <tr>
            <th>名称</th>
            <th>经度</th>
            <th>纬度</th>
            <th>剩余车位</th>
            <th>费率</th>
          </tr>
        </thead>
        <tbody>
          {parkingLots.map((lot) => (
            <tr key={lot.id}>
              <td>{lot.ParkingName}</td>
              <td>{lot.Longitude}</td>
              <td>{lot.Latitude}</td>
              <td>{lot.Capacity}</td>
              <td>{lot.Rates}</td>
            </tr>
          ))}
        </tbody>
      </table>
      <Footer />
    </div>
  );
};

export default ParkingLotListPage;
