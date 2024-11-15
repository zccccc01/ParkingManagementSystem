import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { Table } from 'antd';
import Header from '../components/Header';
import Footer from '../components/Footer';
import './ParkingLotListPage.scss';

const ParkingLotListPage = () => {
  const [parkingLots, setParkingLots] = useState([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);
  const [viewCount, setViewCount] = useState('Loading...'); // 使用单引号

  const fetchParkingLots = async () => {
    setLoading(true);
    try {
      const response = await axios.get('/api/parkinglot');
      setParkingLots(response.data);
      setError(null);
    } catch (err) {
      console.error('Failed to fetch parking lots:', err);
      setError(err);
    } finally {
      setLoading(false);
    }
  };

  // Fetch the view count when the component mounts
  const fetchViewCount = async () => {
    try {
      const { data } = await axios.get('/parking-lots/count'); // 更新为新的端点
      setViewCount(data.count); // 从返回的 JSON 中提取 count
    } catch (err) {
      console.error('Failed to fetch view count:', err);
    }
  };

  // 增加浏览计数并获取停车场信息
  const incrementViewCount = async () => {
    try {
      await axios.get('/parking-lots'); // 发送请求以增加计数
      fetchViewCount(); // 然后获取当前浏览计数
    } catch (err) {
      console.error('Failed to increment view count:', err);
    }
  };

  useEffect(() => {
    incrementViewCount(); // 在组件挂载时增加浏览计数
  }, []);

  const columns = [
    {
      title: '名称',
      dataIndex: 'ParkingName',
      key: 'ParkingName',
    },
    {
      title: '经度',
      dataIndex: 'Longitude',
      key: 'Longitude',
    },
    {
      title: '纬度',
      dataIndex: 'Latitude',
      key: 'Latitude',
    },
    {
      title: '总容量',
      dataIndex: 'Capacity',
      key: 'Capacity',
    },
    {
      title: '收费标准（元/h）',
      dataIndex: 'Rates',
      key: 'Rates',
    },
  ];

  return (
    <div className="parking-lot-list-page">
      <Header />
      <h1>停车场列表</h1>
      <button type="button" onClick={fetchParkingLots}>
        查询
      </button>
      <br />
      {loading && <p>加载中...</p>}
      {error && <p className="error">加载失败: {error.message}</p>}
      {!loading && !error && (
        <div>
          <h1>停车场列表</h1>
          <Table
            columns={columns}
            dataSource={parkingLots}
            rowKey="id"
            pagination={{ pageSize: 10 }}
          />
        </div>
      )}
      <Footer />
      {/* 浏览计数显示 */}
      <div id="viewCount" className="view-count">
        浏览次数：{viewCount}
      </div>
    </div>
  );
};

export default ParkingLotListPage;
