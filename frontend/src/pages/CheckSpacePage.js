// src/pages/CheckSpacePage.js
import React, { useState } from 'react';
import axios from 'axios';
import { Table } from 'antd';
import Header from '../components/AdminHeader';
import Footer from '../components/Footer';
import './CheckSpacePage.scss';

const CheckSpacePage = () => {
  const [parkingLotID, setParkingLotID] = useState('');
  const [startTime, setStartTime] = useState('');
  const [endTime, setEndTime] = useState('');
  const [occupancyData, setOccupancyData] = useState([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);

  const formatTime = (time) => {
    const date = new Date(time);
    return date.toISOString().slice(0, 19);
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    setLoading(true);
    setError(null);

    const formattedStartTime = formatTime(startTime);
    const formattedEndTime = formatTime(endTime);

    try {
      const url = `/api/parkinglot/id/${parkingLotID}/start/${formattedStartTime}/end/${formattedEndTime}`;
      console.log('Request URL:', url);

      const response = await axios.get(url);

      console.log('Response Data:', response.data);
      setOccupancyData(response.data); // 直接设置响应数据
      setError(null);
    } catch (err) {
      console.error('Failed to fetch parking spaces:', err);
      if (err.response) {
        // 请求已发出，但服务器响应的状态码不在 2xx 范围内
        setError(
          `Server Error: ${err.response.status} - ${err.response.data.message || err.response.statusText}`
        );
      } else if (err.request) {
        // 请求已发出，但没有收到响应
        setError('No response from server');
      } else {
        // 发生了一些问题，导致请求无法发出
        setError('Error creating request');
      }
    } finally {
      setLoading(false);
    }
  };

  const columns = [
    {
      title: '车位ID',
      dataIndex: 'SpaceID',
      key: 'SpaceID',
    },
    {
      title: '状态',
      dataIndex: 'Status',
      key: 'Status',
    },
    {
      title: '停车场ID',
      dataIndex: 'ParkingLotID',
      key: 'ParkingLotID',
    },
  ];

  return (
    <div className="CheckSpacePage">
      <Header />
      <h1>查询停车位页面</h1>
      <p>查询停车场特定时段的占用状态</p>
      <form onSubmit={handleSubmit}>
        <div>
          <label htmlFor="parkingLotID">停车场ID:</label>
          <input
            type="number"
            id="parkingLotID"
            value={parkingLotID}
            onChange={(e) => setParkingLotID(e.target.value)}
            required
          />
        </div>
        <br />
        <div>
          <label htmlFor="startTime">开始时间:</label>
          <input
            type="datetime-local"
            id="startTime"
            value={startTime}
            onChange={(e) => setStartTime(e.target.value)}
            required
          />
        </div>
        <br />
        <div>
          <label htmlFor="endTime">结束时间:</label>
          <input
            type="datetime-local"
            id="endTime"
            value={endTime}
            onChange={(e) => setEndTime(e.target.value)}
            required
          />
        </div>
        <br />
        <button type="submit" disabled={loading}>
          查询
        </button>
      </form>
      {loading && <p>加载中...</p>}
      {error && <p className="error">加载失败: {error}</p>}
      {occupancyData.length > 0 && (
        <Table
          columns={columns}
          dataSource={occupancyData}
          pagination={{
            pageSize: 10, // 每页显示的条目数
            showSizeChanger: true, // 允许用户改变每页显示的条目数
            showQuickJumper: true, // 允许用户快速跳转到指定页
          }}
          loading={loading}
          rowKey="SpaceID"
        />
      )}
      <Footer />
    </div>
  );
};

export default CheckSpacePage;
