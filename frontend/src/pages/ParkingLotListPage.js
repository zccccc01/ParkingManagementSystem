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
      <div className="two-column-table">
        <div className="column left">
          <table>
            <thead>
              <tr>
                <th>名称</th>
                <th>经度</th>
                <th>纬度</th>
                <th>总容量</th>
                <th>收费标准（元/h）</th>
              </tr>
            </thead>
            <tbody>
              {parkingLots
                .filter((_, index) => index % 2 === 0)
                .map((lot) => (
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
        </div>
        <div className="column right">
          <table>
            <thead>
              <tr>
                <th>名称</th>
                <th>经度</th>
                <th>纬度</th>
                <th>总容量</th>
                <th>收费标准（元/h）</th>
              </tr>
            </thead>
            <tbody>
              {parkingLots
                .filter((_, index) => index % 2 !== 0)
                .map((lot) => (
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
        </div>
        <div className="vertical-divider" />
      </div>
      <Footer />
    </div>
  );
};

export default ParkingLotListPage;
