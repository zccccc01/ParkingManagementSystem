// pages/AdminDashboard.js
import React, { useState, useEffect } from 'react';
import axios from 'axios';
import Header from '../components/Header';
import Footer from '../components/Footer';

const AdminDashboard = () => {
  const [income, setIncome] = useState(null);
  const [rate, setRate] = useState(null);
  const [loading, setLoading] = useState(false);
  const [selectedId, setSelectedId] = useState('');

  const fetchData = async (id) => {
    try {
      setLoading(true);
      const incomeResponse = await axios.get(`http://localhost:8000/api/parkinglot/income/${id}`);
      if (incomeResponse.data !== undefined) {
        setIncome(incomeResponse.data);
      } else {
        throw new Error('No income data returned');
      }

      const rateResponse = await axios.get(
        `http://localhost:8000/api/parkinglot/occupancy-rate/${id}`
      );
      if (rateResponse.data !== undefined) {
        setRate(rateResponse.data);
      } else {
        throw new Error('No occupancy-rate data returned');
      }
    } catch (error) {
      console.error('Failed to fetch data:', error);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    if (selectedId) {
      fetchData(selectedId);
    }
  }, [selectedId]);

  return (
    <div className="admin-dashboard">
      <Header />
      <h1>管理员仪表盘</h1>
      <p>提供停车场管理功能和数据分析</p>

      <div>
        <label htmlFor="userId">请输入停车场ID:</label>
        <input
          type="text"
          id="userId"
          value={selectedId}
          onChange={(e) => setSelectedId(e.target.value)}
        />
      </div>

      {loading ? (
        <p>Loading...</p>
      ) : (
        <div>
          <h2>收入信息</h2>
          {income !== null ? <p>总收入: {income}</p> : <p>无收入数据</p>}
          <h2>占用率信息</h2>
          {rate !== null ? <p>占用率: {rate}</p> : <p>无占用率数据</p>}
        </div>
      )}

      <Footer />
    </div>
  );
};

export default AdminDashboard;
