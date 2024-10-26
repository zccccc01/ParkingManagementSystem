import React, { useState, useEffect } from 'react';
import axios from 'axios'; // 引入axios
import Header from '../components/Header';
import Footer from '../components/Footer';
import './ParkingRecordPage.scss';

const ParkingRecordPage = () => {
  const [records, setRecords] = useState([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);

  const fetchRecords = async () => {
    try {
      setLoading(true);
      const response = await fetch('http://localhost:8000/api/user', {
        method: 'GET',
        credentials: 'include', // 确保请求带上 cookie
      });

      if (response.ok) {
        const userData = await response.json(); // 获取用户数据
        if (!userData || !userData.id) {
          throw new Error('User data is invalid or missing ID');
        }

        const recordResponse = await axios.get(
          `http://localhost:8000/api/parkingrecord/user/${userData.id}`
        ); // 使用用户 ID

        if (Array.isArray(recordResponse.data.records)) {
          setRecords(recordResponse.data.records);
        } else {
          throw new Error('Invalid record data format');
        }
      } else {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
    } catch (fetchError) {
      console.error('Failed to fetch record:', fetchError);
      setError(fetchError.message);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchRecords();
  }, []);

  return (
    <div className="parking-record-page">
      <Header />
      <h1>停车记录页面</h1>
      <p>展示用户的停车历史和支付记录</p>

      {loading && <p>Loading...</p>}
      {!loading && error && <p style={{ color: 'red' }}>{error}</p>}
      {!loading && !error && records.length === 0 && <p>暂无停车记录</p>}
      {!loading && !error && records.length > 0 && (
        <table>
          <thead>
            <tr>
              <th>空间ID</th>
              <th>车辆 ID</th>
              <th>批号</th>
              <th>开始时间</th>
              <th>结束时间</th>
              <th>费用</th>
            </tr>
          </thead>
          <tbody>
            {records.map((record) => (
              <tr key={record.RecordID}>
                <td>{record.SpaceID}</td>
                <td>{record.VehicleID}</td>
                <td>{record.LotID}</td>
                <td>{record.StartTime}</td>
                <td>{record.EndTime}</td>
                <td>{record.Fee}</td>
              </tr>
            ))}
          </tbody>
        </table>
      )}

      <Footer />
    </div>
  );
};

export default ParkingRecordPage;
