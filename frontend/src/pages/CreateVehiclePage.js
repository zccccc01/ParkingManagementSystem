// pages/CreateVehiclePage.js
import React, { useState } from 'react';
import Header from '../components/Header';
import Footer from '../components/Footer';
import './CreateVehiclePage.scss';

const CreateVehiclePage = () => {
  const [vehicle, setVehicle] = useState({
    VehicleID: '',
    UserID: '',
    PlateNumber: '',
    Color: '',
  });

  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);
  const [success, setSuccess] = useState(false);

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
      const response = await fetch('http://localhost:8000/api/vehicle', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          VehicleID: parseInt(vehicle.VehicleID, 10),
          UserID: parseInt(vehicle.UserID, 10),
          PlateNumber: vehicle.PlateNumber,
          Color: vehicle.Color,
        }),
      });

      console.log('Response:', response); // 绑定日志输出

      if (!response.ok) {
        const errorData = await response.json();
        throw new Error(
          `HTTP error! status: ${response.status}, message: ${errorData.message || '未知错误'}`
        );
      }

      const data = await response.json();
      console.log('Data:', data); // 绑定日志输出

      if (data.message === 'Reservation created successfully' || data.VehicleID) {
        setSuccess(true);
      } else {
        setError('绑定失败，请稍后再试');
      }
    } catch (err) {
      setError(err.message);
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
          <label htmlFor="UserID">用户ID:</label>
          <input
            type="text"
            name="UserID"
            value={vehicle.UserID}
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
