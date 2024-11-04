import React, { useState } from 'react';
import Header from '../components/Header';
import Footer from '../components/Footer';
import './BookingPage.scss';

const BookingPage = () => {
  const [formData, setFormData] = useState({
    reservationID: '',
    startTime: '',
    endTime: '',
    spaceID: '',
    vehicleID: '',
    lotID: '',
  });

  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);
  const [success, setSuccess] = useState(false);

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData({ ...formData, [name]: value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    setLoading(true);
    setError(null);
    setSuccess(false);

    try {
      // 转换时间格式为 ISO 8601 并添加时区信息
      const startTime = new Date(formData.startTime).toISOString();
      const endTime = new Date(formData.endTime).toISOString();

      const response = await fetch('http://localhost:8000/api/reservation/', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          ReservationID: parseInt(formData.reservationID, 10),
          StartTime: startTime,
          EndTime: endTime,
          SpaceID: parseInt(formData.spaceID, 10),
          VehicleID: parseInt(formData.vehicleID, 10),
          LotID: parseInt(formData.lotID, 10),
        }),
      });

      if (!response.ok) {
        const errorData = await response.json();
        throw new Error(
          `HTTP error! status: ${response.status}, message: ${errorData.message || '未知错误'}`
        );
      }

      const data = await response.json();
      if (data.message === 'Reservation created successfully') {
        setSuccess(true);
      } else {
        setError('预约失败，请稍后再试');
      }
    } catch (err) {
      setError(err.message);
    } finally {
      setLoading(false);
    }
  };

  const handleUpdate = async (e) => {
    e.preventDefault();
    setLoading(true);
    setError(null);
    setSuccess(false);

    try {
      // 转换时间格式为 ISO 8601 并添加时区信息
      const startTime = new Date(formData.startTime).toISOString();
      const endTime = new Date(formData.endTime).toISOString();

      const response = await fetch(
        `http://localhost:8000/api/reservation/id/${formData.reservationID}`,
        {
          method: 'PUT',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({
            LotID: parseInt(formData.lotID, 10),
            SpaceID: parseInt(formData.spaceID, 10),
            StartTime: startTime,
            EndTime: endTime,
          }),
        }
      );

      if (!response.ok) {
        const errorData = await response.json();
        throw new Error(
          `HTTP error! status: ${response.status}, message: ${errorData.message || '未知错误'}`
        );
      }

      const data = await response.json();
      if (data.message === 'Reservation updated successfully') {
        setSuccess(true);
      } else {
        setError('更新失败，请稍后再试');
      }
    } catch (err) {
      setError(err.message);
    } finally {
      setLoading(false);
    }
  };

  const handleDelete = async (e) => {
    e.preventDefault();
    setLoading(true);
    setError(null);
    setSuccess(false);

    try {
      const response = await fetch(
        `http://localhost:8000/api/reservation/id/${formData.reservationID}`,
        {
          method: 'DELETE',
          headers: {
            'Content-Type': 'application/json',
          },
        }
      );

      if (!response.ok) {
        const errorData = await response.json();
        throw new Error(
          `HTTP error! status: ${response.status}, message: ${errorData.message || '未知错误'}`
        );
      }

      const data = await response.json();
      console.log(data); // 添加这行日志，查看响应体内容

      if (data.message === 'Reservation cancelled successfully') {
        setSuccess(true);
      } else {
        setError('取消失败，请稍后再试');
      }
    } catch (err) {
      setError(err.message);
    } finally {
      setLoading(false);
    }
  };
  return (
    <div className="booking-page">
      <Header />
      <h1>预定页面</h1>
      <p>选择停车时段和支付方式</p>
      <form onSubmit={handleSubmit} className="booking-form">
        <div className="form-group">
          <label htmlFor="reservationID">预约ID:</label>
          <input
            type="number"
            id="reservationID"
            name="reservationID"
            value={formData.reservationID}
            onChange={handleChange}
            required
            className="form-input"
          />
        </div>
        <div className="form-group">
          <label htmlFor="startTime">开始时间:</label>
          <input
            type="datetime-local"
            id="startTime"
            name="startTime"
            value={formData.startTime}
            onChange={handleChange}
            required
            className="form-input"
          />
        </div>
        <div className="form-group">
          <label htmlFor="endTime">结束时间:</label>
          <input
            type="datetime-local"
            id="endTime"
            name="endTime"
            value={formData.endTime}
            onChange={handleChange}
            required
            className="form-input"
          />
        </div>
        <div className="form-group">
          <label htmlFor="spaceID">停车位:</label>
          <input
            type="number"
            id="spaceID"
            name="spaceID"
            value={formData.spaceID}
            onChange={handleChange}
            required
            className="form-input"
          />
        </div>
        <div className="form-group">
          <label htmlFor="vehicleID">车辆 ID:</label>
          <input
            type="number"
            id="vehicleID"
            name="vehicleID"
            value={formData.vehicleID}
            onChange={handleChange}
            required
            className="form-input"
          />
        </div>
        <div className="form-group">
          <label htmlFor="lotID">停车场:</label>
          <input
            type="number"
            id="lotID"
            name="lotID"
            value={formData.lotID}
            onChange={handleChange}
            required
            className="form-input"
          />
        </div>
        <button type="submit" disabled={loading} className="form-button">
          {loading ? '提交中...' : '提交'}
        </button>
        <button type="button" onClick={handleUpdate} disabled={loading} className="form-button">
          {loading ? '更新中...' : '更新'}
        </button>
        <button type="button" onClick={handleDelete} disabled={loading} className="form-button">
          {loading ? '删除中...' : '删除'}
        </button>
      </form>
      {success && <p>操作成功！</p>}
      {error && <p style={{ color: 'red' }}>{error}</p>}
      <Footer />
    </div>
  );
};

export default BookingPage;
