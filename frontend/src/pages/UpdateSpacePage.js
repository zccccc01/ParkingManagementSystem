import React, { useState } from 'react';
import axios from 'axios';
import Header from '../components/Header';
import Footer from '../components/Footer';
import './ParkingSpacePage.scss';

const UpdateSpacePage = () => {
  const [parkingSpaces, setParkingSpaces] = useState([]); // 初始化为一个空数组
  const [lotid, setLotid] = useState('');
  const [spaceid, setSpaceid] = useState('');
  const [status, setStatus] = useState('');
  const [error, setError] = useState(null);
  const [successMessage, setSuccessMessage] = useState(null); // 添加成功消息状态

  const handleInputChange = (e) => {
    const { name, value } = e.target;
    if (name === 'lotid') {
      setLotid(value);
    } else if (name === 'spaceid') {
      setSpaceid(value);
    } else if (name === 'status') {
      setStatus(value);
    }
  };

  const handleUpdateStatus = async (e) => {
    e.preventDefault();
    try {
      if (Number.isNaN(Number(lotid)) || Number.isNaN(Number(spaceid))) {
        throw new Error('Tips: 无效的停车场 ID 或车位 ID');
      }

      console.log(
        `Updating parking space status for lot ID: ${lotid}, space ID: ${spaceid}, status: ${status}`
      );
      const response = await axios.put(
        `http://localhost:8000/api/parkingspace/status/lot/${lotid}/space/${spaceid}`,
        { status }
      );

      console.log('API Response:', response.data);
      setParkingSpaces(
        parkingSpaces.map((space) =>
          space.SpaceID === spaceid ? { ...space, Status: status } : space
        )
      );
      setError(null);
      setSuccessMessage('车位状态更新成功'); // 设置成功消息
    } catch (err) {
      console.error('API Request Failed:', err);
      const errorMessage =
        err.response && err.response.data && err.response.data.error
          ? err.response.data.error
          : err.message;
      setError(errorMessage);
      setSuccessMessage(null); // 清除成功消息
    }
  };

  return (
    <div className="ParkingSpacePage">
      <Header />
      <h1>更新车位状态</h1>
      <form onSubmit={handleUpdateStatus}>
        <label>
          停车场 ID:
          <input type="text" name="lotid" value={lotid} onChange={handleInputChange} />
        </label>
        <label>
          停车位 ID:
          <input type="text" name="spaceid" value={spaceid} onChange={handleInputChange} />
        </label>
        <label>
          状态:
          <select name="status" value={status} onChange={handleInputChange}>
            <option value="">请选择状态</option>
            <option value="FREE">FREE</option>
            <option value="OCCUPIED">OCCUPIED</option>
            <option value="RESERVED">RESERVED</option>
          </select>
        </label>
        <button type="submit">更新状态</button>
      </form>
      {error && <p style={{ color: 'red' }}>{error}</p>}
      {successMessage && <p style={{ color: 'green' }}>{successMessage}</p>} {/* 渲染成功消息 */}
      <Footer />
    </div>
  );
};

export default UpdateSpacePage;
