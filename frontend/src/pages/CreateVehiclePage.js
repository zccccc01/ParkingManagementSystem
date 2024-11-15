// pages/CreateVehiclePage.js
import React, { useState, useEffect } from 'react';
import axios from 'axios'; // 引入axios
import Header from '../components/Header';
import Footer from '../components/Footer';
import './CreateVehiclePage.scss';

const CreateVehiclePage = () => {
  const [vehicle, setVehicle] = useState({
    VehicleID: '',
    PlateNumber: '',
    Color: '',
  });

  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);
  const [success, setSuccess] = useState(false);
  const [userID, setUserID] = useState('');

  const fetchUserID = async () => {
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

        setUserID(userData.id);
      } else {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
    } catch (fetchError) {
      console.error('Failed to fetch user ID:', fetchError);
      setError(fetchError.message);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchUserID();
  }, []);

  const handleChange = (e) => {
    const { name, value } = e.target;
    setVehicle({ ...vehicle, [name]: value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    setLoading(true);
    setError(null);
    setSuccess(false);

    try {
      console.log('Submitting with UserID:', userID); // 添加日志输出
      const response = await axios.post('/api/vehicle', {
        VehicleID: parseInt(vehicle.VehicleID, 10),
        UserID: parseInt(userID, 10),
        PlateNumber: vehicle.PlateNumber,
        Color: vehicle.Color,
      });

      console.log('Response:', response); // 绑定日志输出

      if (response.status === 200 || response.status === 201) {
        // 处理 200 和 201 状态码
        const { data } = response; // 使用对象解构
        if (data.VehicleID) {
          setSuccess(true);
        } else {
          setError('绑定失败，请稍后再试');
        }
      } else {
        setError(`HTTP error! status: ${response.status}`);
      }
    } catch (err) {
      console.error('Failed to submit form:', err);
      setError(`提交失败: ${err.response ? err.response.data.message : err.message}`);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="CreateVehiclePage">
      <Header />
      <h1>绑定车辆页面</h1>
      <p>用户绑定车辆</p>
      <form onSubmit={handleSubmit} className="create-vehicle-form">
        <div className="form-group">
          <label htmlFor="VehicleID">车辆ID:</label>
          <input
            type="text"
            name="VehicleID"
            value={vehicle.VehicleID}
            onChange={handleChange}
            required
            className="form-input"
          />
        </div>
        <div className="form-group">
          <label htmlFor="PlateNumber">车牌号:</label>
          <input
            type="text"
            name="PlateNumber"
            value={vehicle.PlateNumber}
            onChange={handleChange}
            required
            className="form-input"
          />
        </div>
        <div className="form-group">
          <label htmlFor="Color">颜色:</label>
          <input
            type="text"
            name="Color"
            value={vehicle.Color}
            onChange={handleChange}
            required
            className="form-input"
          />
        </div>
        <button type="submit" disabled={loading} className="form-button">
          {loading ? '提交中...' : '提交'}
        </button>
      </form>
      {success && <p>车辆绑定成功！</p>}
      {error && <p style={{ color: 'red' }}>{error}</p>}
      <Footer />
    </div>
  );
};

export default CreateVehiclePage;
