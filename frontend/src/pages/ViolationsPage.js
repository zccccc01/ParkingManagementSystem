// pages/ViolationsPage.js
import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { Table } from 'antd';
import Header from '../components/Header';
import Footer from '../components/Footer';
import './ViolationsPage.scss';

const ViolationsPage = () => {
  const [violationRecords, setViolationRecords] = useState([]);
  const [loading, setLoading] = useState(false);
  const [errorMessage, setErrorMessage] = useState(null);

  const fetchUserData = async () => {
    try {
      const response = await fetch('http://localhost:8000/api/user', {
        method: 'GET',
        credentials: 'include', // 确保请求带上 cookie
      });

      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }

      const userData = await response.json();
      return userData;
    } catch (fetchError) {
      console.error('Failed to fetch user data:', fetchError);
      throw fetchError;
    }
  };

  const fetchViolationRecords = async (userId) => {
    try {
      const response = await axios.get(`http://localhost:8000/api/violationrecord/user/${userId}`);
      console.log('Response data:', response.data); // 添加调试信息

      // 检查返回的数据是否包含 violationRecords 数组
      if (!response.data || !Array.isArray(response.data.violationRecords)) {
        throw new Error('No violation record data returned or data is not an array');
      }

      return response.data.violationRecords;
    } catch (fetchError) {
      console.error('Failed to fetch violation records:', fetchError);
      throw fetchError;
    }
  };

  const loadViolationRecords = async () => {
    try {
      setLoading(true);
      const userData = await fetchUserData();
      const records = await fetchViolationRecords(userData.id);
      setViolationRecords(records);
    } catch (loadError) {
      setErrorMessage(loadError.message);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    loadViolationRecords();
  }, []);

  const columns = [
    {
      title: '记录 ID',
      dataIndex: 'RecordID',
      key: 'RecordID',
    },
    {
      title: '罚款金额',
      dataIndex: 'FineAmount',
      key: 'FineAmount',
    },
    {
      title: '违规类型',
      dataIndex: 'ViolationType',
      key: 'ViolationType',
    },
    {
      title: '状态',
      dataIndex: 'Status',
      key: 'Status',
    },
  ];

  return (
    <div className="violations-page">
      <Header />
      <h1>违规停车页面</h1>
      <p>展示违规记录和处罚信息</p>

      {loading && <p>Loading...</p>}
      {!loading && errorMessage && <p style={{ color: 'red' }}>{errorMessage}</p>}
      {!loading && !errorMessage && (
        <div>
          {violationRecords.length === 0 ? (
            <p>没有违规记录</p>
          ) : (
            <Table
              columns={columns}
              dataSource={violationRecords}
              rowKey="RecordID"
              pagination={{ pageSize: 10 }}
            />
          )}
        </div>
      )}

      <Footer />
    </div>
  );
};

export default ViolationsPage;
