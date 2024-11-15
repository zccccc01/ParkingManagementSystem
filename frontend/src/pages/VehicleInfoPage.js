// pages/VehicleInfoPage.js
import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { Table } from 'antd';
import Header from '../components/Header';
import Footer from '../components/Footer';
import './VehicleInfoPage.scss';

const VehicleInfoPage = () => {
  const [vehicles, setVehicles] = useState([]);
  const [loading, setLoading] = useState(false);

  const fetchVehicles = async () => {
    try {
      setLoading(true);
      const response = await fetch('/api/user', {
        method: 'GET',
        credentials: 'include', // 确保请求带上 cookie
      });

      if (response.ok) {
        const userData = await response.json(); // 获取用户数据
        const vehicleResponse = await axios.get(`/api/vehicle/user/${userData.id}`); // 使用用户 ID
        if (vehicleResponse.data) {
          setVehicles(vehicleResponse.data);
        } else {
          throw new Error('No vehicle data returned');
        }
      } else {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
    } catch (error) {
      console.error('Failed to fetch vehicles:', error);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchVehicles();
  }, []);

  const columns = [
    {
      title: '用户 ID',
      dataIndex: 'UserID',
      key: 'UserID',
    },
    {
      title: '车牌号',
      dataIndex: 'PlateNumber',
      key: 'PlateNumber',
    },
    {
      title: '颜色',
      dataIndex: 'Color',
      key: 'Color',
    },
  ];

  return (
    <div className="vehicle-info-page">
      <Header />
      <h1>车辆信息页面</h1>
      <p>展示用户所有车辆信息</p>

      {loading ? (
        <p>Loading...</p>
      ) : (
        <Table columns={columns} dataSource={vehicles} rowKey="id" pagination={{ pageSize: 10 }} />
      )}

      <Footer />
    </div>
  );
};

export default VehicleInfoPage;
