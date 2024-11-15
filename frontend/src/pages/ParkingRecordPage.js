import React, { useState, useEffect } from 'react';
import axios from 'axios'; // 引入axios
import { Table } from 'antd';
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
      const response = await fetch('/api/user', {
        method: 'GET',
        credentials: 'include', // 确保请求带上 cookie
      });

      if (response.ok) {
        const userData = await response.json(); // 获取用户数据
        if (!userData || !userData.id) {
          throw new Error('User data is invalid or missing ID');
        }

        const recordResponse = await axios.get(`/api/parkingrecord/user/${userData.id}`); // 使用用户 ID

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

  const columns = [
    {
      title: '车位 ID',
      dataIndex: 'SpaceID',
      key: 'SpaceID',
    },
    {
      title: '车辆 ID',
      dataIndex: 'VehicleID',
      key: 'VehicleID',
    },
    {
      title: '停车场 ID',
      dataIndex: 'LotID',
      key: 'LotID',
    },
    {
      title: '进场时间',
      dataIndex: 'StartTime',
      key: 'StartTime',
    },
    {
      title: '离场时间',
      dataIndex: 'EndTime',
      key: 'EndTime',
    },
    {
      title: '停车费',
      dataIndex: 'Fee',
      key: 'Fee',
    },
  ];

  return (
    <div className="parking-record-page">
      <Header />
      <h1>停车记录页面</h1>
      <p>展示用户的停车历史和支付记录</p>

      {loading && <p>Loading...</p>}
      {!loading && error && <p style={{ color: 'red' }}>{error}</p>}
      {!loading && !error && records.length === 0 && <p>暂无停车记录</p>}
      {!loading && !error && records.length > 0 && (
        <Table
          columns={columns}
          dataSource={records}
          rowKey="RecordID"
          pagination={{ pageSize: 10 }}
        />
      )}

      <Footer />
    </div>
  );
};

export default ParkingRecordPage;
