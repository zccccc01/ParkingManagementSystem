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
        <table>
          <thead>
            <tr>
              <th>车位 ID</th>
              <th>状态</th>
              <th>停车场 ID</th>
              <th className="vertical-line" />
              <th>车位 ID</th>
              <th>状态</th>
              <th>停车场 ID</th>
            </tr>
          </thead>
          <tbody>
            {parkingSpaces.map((space, index) =>
              index % 2 === 0 ? (
                <tr key={space.SpaceID}>
                  <td>{space.SpaceID}</td>
                  <td>{space.Status}</td>
                  <td>{space.ParkingLotID}</td>
                  <td className="vertical-line" />
                  {parkingSpaces[index + 1] ? (
                    <>
                      <td>{parkingSpaces[index + 1].SpaceID}</td>
                      <td>{parkingSpaces[index + 1].Status}</td>
                      <td>{parkingSpaces[index + 1].ParkingLotID}</td>
                    </>
                  ) : null}
                </tr>
              ) : null
            )}
          </tbody>
        </table>
      )}
      <Footer />
    </div>
  );
};

export default ParkingSpacePage;
