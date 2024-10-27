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
            <th>总容量</th>
            <th>收费标准（元/h）</th>
            <th className="vertical-line" />
            <th>名称</th>
            <th>经度</th>
            <th>纬度</th>
            <th>总容量</th>
            <th>收费标准（元/h）</th>
          </tr>
        </thead>
        <tbody>
          {parkingLots.map((lot, index) =>
            index % 2 === 0 ? (
              <tr key={lot.id}>
                <td>{lot.ParkingName}</td>
                <td>{lot.Longitude}</td>
                <td>{lot.Latitude}</td>
                <td>{lot.Capacity}</td>
                <td>{lot.Rates}</td>
                <td className="vertical-line" />
                {parkingLots[index + 1] ? (
                  <>
                    <td>{parkingLots[index + 1].ParkingName}</td>
                    <td>{parkingLots[index + 1].Longitude}</td>
                    <td>{parkingLots[index + 1].Latitude}</td>
                    <td>{parkingLots[index + 1].Capacity}</td>
                    <td>{parkingLots[index + 1].Rates}</td>
                  </>
                ) : (
                  <>
                    <td />
                    <td />
                    <td />
                    <td />
                    <td />
                  </>
                )}
              </tr>
            ) : null
          )}
        </tbody>
      </table>
      <Footer />
    </div>
  );
};

export default ParkingLotListPage;
