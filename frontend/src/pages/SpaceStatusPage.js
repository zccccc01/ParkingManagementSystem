// src/pages/ParkingSpacePage.js
// TODO: 查看实时空闲车位信息
import React, { useState } from 'react';
import axios from 'axios';
import Header from '../components/Header';
import Footer from '../components/Footer';
import './ParkingSpacePage.scss';

const ParkingSpacePage = () => {
  const [lotid, setLotid] = useState('');
  const [parkingSpaces, setParkingSpaces] = useState(null);
  const [error, setError] = useState(null);

  const fetchParkingSpaces = async () => {
    try {
      if (Number.isNaN(Number(lotid))) {
        throw new Error('Tips: 无效的停车场 ID');
      }

      console.log(`Fetching parking spaces for lot ID: ${lotid}`);
      const response = await axios.get(`http://localhost:8000/api/parkingspace/lot/${lotid}`);

      console.log('API Response:', response.data);
      setParkingSpaces(response.data.spaces);
      setError(null);
    } catch (err) {
      console.error('API Request Failed:', err);
      const errorMessage =
        err.response && err.response.data && err.response.data.error
          ? err.response.data.error
          : err.message;
      setError(errorMessage);
    }
  };

  const handleInputChange = (e) => {
    const { name, value } = e.target;
    if (name === 'lotid') {
      setLotid(value);
    }
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    fetchParkingSpaces();
  };

  return (
    <div className="ParkingSpacePage">
      <Header />
      <h1>查询车位状态信息页面</h1>
      <form onSubmit={handleSubmit}>
        <label>
          停车场 ID:
          <input type="text" name="lotid" value={lotid} onChange={handleInputChange} />
        </label>
        <button type="submit">查询</button>
      </form>

      {error && <p>Error: {error}</p>}
      {!parkingSpaces && !error && <p>Tips: 请输入停车场 ID 进行查询</p>}

      {parkingSpaces && (
        <div>
          <h1>车位状态</h1>
          <table>
            <thead>
              <tr>
                <th>车位 ID</th>
                <th>状态</th>
                <th>停车场 ID</th>
              </tr>
            </thead>
            <tbody>
              {parkingSpaces.map((space) => (
                <tr key={space.SpaceID}>
                  <td>{space.SpaceID}</td>
                  <td>{space.Status}</td>
                  <td>{space.ParkingLotID}</td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      )}
      <Footer />
    </div>
  );
};

export default ParkingSpacePage;
