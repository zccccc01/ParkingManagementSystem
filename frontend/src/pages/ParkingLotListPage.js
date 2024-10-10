import React, { useState, useEffect } from 'react';
import Header from '../components/Header';
import Footer from '../components/Footer';
import './ParkingLotListPage.scss';

const ParkingLotListPage = () => {
  const [showTable, setShowTable] = useState(false);
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
    // 可以在这里设置自动加载数据或者只在需要时加载
  }, []);

  return (
    <div className="parking-lot-list-page">
      <Header />
      <h1>停车场列表</h1>
      <button
        type="button"
        onClick={() => {
          setShowTable(true);
          fetchParkingLots();
        }}
      >
        查看停车场信息
      </button>
      {showTable && (
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
      )}
      <Footer />
    </div>
  );
};

export default ParkingLotListPage;
