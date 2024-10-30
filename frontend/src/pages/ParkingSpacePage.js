// src/pages/ParkingSpacePage.js
// TODO: 查看实时空闲车位信息
import React, { useState } from 'react';
import axios from 'axios';
import Header from '../components/Header';
import Footer from '../components/Footer';
import './ParkingSpacePage.scss';

const ParkingSpacePage = () => {
  const [parkingSpaces, setParkingSpaces] = useState([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);

  const fetchParkingSpaces = async () => {
    setLoading(true);
    try {
      const { data } = await axios.get('http://localhost:8000/api/parkingspace/status/free'); // 使用对象解构
      setParkingSpaces(data.spaces); // 提取 spaces 数组
      setError(null);
    } catch (err) {
      console.error('Failed to fetch parking spaces:', err);
      setError(err);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="ParkingSpacePage">
      <Header />
      <h1>空闲车位信息页面</h1>
      <button type="button" onClick={fetchParkingSpaces}>
        查询
      </button>
      {loading && <p>加载中...</p>}
      {error && <p>加载失败: {error.message}</p>}
      {!loading && !error && (
        <div className="two-column-table">
          <div className="column left">
            <table>
              <thead>
                <tr>
                  <th>车位 ID</th>
                  <th>状态</th>
                  <th>停车场 ID</th>
                </tr>
              </thead>
              <tbody>
                {parkingSpaces
                  .filter((_, index) => index % 2 === 0)
                  .map((space) => (
                    <tr key={space.SpaceID}>
                      <td>{space.SpaceID}</td>
                      <td>{space.Status}</td>
                      <td>{space.ParkingLotID}</td>
                    </tr>
                  ))}
              </tbody>
            </table>
          </div>
          <div className="column right">
            <table>
              <thead>
                <tr>
                  <th>车位 ID</th>
                  <th>状态</th>
                  <th>停车场 ID</th>
                </tr>
              </thead>
              <tbody>
                {parkingSpaces
                  .filter((_, index) => index % 2 !== 0)
                  .map((space) => (
                    <tr key={space.SpaceID}>
                      <td>{space.SpaceID}</td>
                      <td>{space.Status}</td>
                      <td>{space.ParkingLotID}</td>
                    </tr>
                  ))}
              </tbody>
            </table>
          </div>
          <div className="vertical-divider" />
        </div>
      )}
      <Footer />
    </div>
  );
};

export default ParkingSpacePage;
