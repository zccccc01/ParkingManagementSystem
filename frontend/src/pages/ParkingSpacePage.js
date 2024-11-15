import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { Table } from 'antd';
import Header from '../components/Header';
import Footer from '../components/Footer';
import './ParkingSpacePage.scss';

const ParkingSpacePage = () => {
  const [parkingSpaces, setParkingSpaces] = useState([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);
  const [viewCount, setViewCount] = useState('Loading...');

  const fetchParkingSpaces = async () => {
    setLoading(true);
    try {
      const { data } = await axios.get('/api/parkingspace/status/free');
      setParkingSpaces(data.spaces);
      setError(null);
    } catch (err) {
      console.error('Failed to fetch parking spaces:', err);
      setError(err);
    } finally {
      setLoading(false);
    }
  };

  const fetchViewCount = async () => {
    try {
      const { data } = await axios.get('/parking-space/count');
      setViewCount(data.count);
    } catch (err) {
      console.error('Failed to fetch view count:', err);
    }
  };

  const incrementViewCount = async () => {
    try {
      await axios.get('/parking-space');
      fetchViewCount();
    } catch (err) {
      console.error('Failed to increment view count:', err);
    }
  };

  useEffect(() => {
    incrementViewCount();
  }, []);

  const columns = [
    {
      title: '车位 ID',
      dataIndex: 'SpaceID',
      key: 'SpaceID',
    },
    {
      title: '状态',
      dataIndex: 'Status',
      key: 'Status',
    },
    {
      title: '停车场 ID',
      dataIndex: 'ParkingLotID',
      key: 'ParkingLotID',
    },
  ];

  return (
    <div className="ParkingSpacePage">
      <Header />
      <h1>空闲车位信息页面</h1>
      <button type="button" onClick={fetchParkingSpaces}>
        查询
      </button>
      <br />
      {loading && <p>加载中...</p>}
      {error && <p className="error">加载失败: {error.message}</p>}
      {!loading && !error && (
        <div>
          <h1>车位状态</h1>
          <Table
            columns={columns}
            dataSource={parkingSpaces}
            rowKey="SpaceID"
            pagination={{ pageSize: 10 }}
          />
        </div>
      )}
      <Footer />
      <div id="viewCount" className="view-count">
        浏览次数：{viewCount}
      </div>
    </div>
  );
};

export default ParkingSpacePage;
