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
    paymentMethod: '', // Payment method
  });

  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);
  const [success, setSuccess] = useState(false);
  const [fee, setFee] = useState(null); // Declare fee state
  const [showModal, setShowModal] = useState(false); // Declare showModal state
  const [paymentComplete, setPaymentComplete] = useState(false); // Declare paymentComplete state

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData({ ...formData, [name]: value });
  };

  const getFee = async () => {
    try {
      const startTime = new Date(formData.startTime).toISOString();
      const endTime = new Date(formData.endTime).toISOString();

      const response = await fetch(
        `/api/reservation/lot/${formData.lotID}?start=${startTime}&end=${endTime}`,
        {
          method: 'GET',
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
      setFee(data.fee); // Set the fee value received from the API
    } catch (err) {
      setError(err.message);
    }
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    setLoading(true);
    setError(null);
    setSuccess(false);

    // 验证 startTime 是否早于 endTime
    const startTime = new Date(formData.startTime);
    const endTime = new Date(formData.endTime);
    if (startTime >= endTime) {
      setError('开始时间必须早于结束时间');
      setLoading(false);
      return;
    }

    try {
      await getFee(); // Get the fee before creating the reservation

      const startTimeISO = startTime.toISOString();
      const endTimeISO = endTime.toISOString();

      const response = await fetch('/api/reservation/', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          ReservationID: parseInt(formData.reservationID, 10),
          StartTime: startTimeISO,
          EndTime: endTimeISO,
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
        setShowModal(true); // Show the modal to choose payment method
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

    // 验证 startTime 是否早于 endTime
    const startTime = new Date(formData.startTime);
    const endTime = new Date(formData.endTime);
    if (startTime >= endTime) {
      setError('开始时间必须早于结束时间');
      setLoading(false);
      return;
    }

    try {
      const startTimeISO = startTime.toISOString();
      const endTimeISO = endTime.toISOString();

      const response = await fetch(`/api/reservation/id/${formData.reservationID}`, {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          LotID: parseInt(formData.lotID, 10),
          SpaceID: parseInt(formData.spaceID, 10),
          StartTime: startTimeISO,
          EndTime: endTimeISO,
        }),
      });

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
      const response = await fetch(`/api/reservation/id/${formData.reservationID}`, {
        method: 'DELETE',
        headers: {
          'Content-Type': 'application/json',
        },
      });

      if (!response.ok) {
        const errorData = await response.json();
        throw new Error(
          `HTTP error! status: ${response.status}, message: ${errorData.message || '未知错误'}`
        );
      }

      const data = await response.json();
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

  const handlePaymentConfirm = () => {
    setPaymentComplete(true); // Mark payment as complete
  };

  const handleCloseModal = () => {
    setShowModal(false); // Close the modal
    setPaymentComplete(false); // Reset payment completion state
  };

  const getPaymentImageSrc = (method) => {
    switch (method) {
      case 'FREE':
        return 'WeChat-Pay.jpg';
      case 'OCCUPIED':
        return 'Alipay.png';
      case 'RESERVED':
        return 'Credit-Card.jpg';
      default:
        return '';
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

      {success && <p>预约成功！</p>}
      {error && <p style={{ color: 'red' }}>{error}</p>}

      {showModal && (
        <div className="modal">
          <div className="modal-content">
            <h1>费用: {fee}元</h1>
            {!paymentComplete ? (
              <>
                <h2>请选择支付方式:</h2>
                <select name="paymentMethod" value={formData.paymentMethod} onChange={handleChange}>
                  <option value="">请选择支付方式</option>
                  <option value="FREE">微信支付</option>
                  <option value="OCCUPIED">支付宝</option>
                  <option value="RESERVED">信用卡</option>
                </select>
                <br />
                <br />
                <div className="payment-buttons">
                  <button
                    type="button"
                    onClick={handlePaymentConfirm}
                    disabled={loading}
                    className="form-button"
                  >
                    确认支付
                  </button>
                  <button
                    type="button"
                    onClick={handleCloseModal}
                    disabled={loading}
                    className="form-button"
                  >
                    稍后支付
                  </button>
                </div>
              </>
            ) : (
              <>
                <img
                  src={getPaymentImageSrc(formData.paymentMethod)}
                  alt="支付方式"
                  className="payment-image"
                />
                <button type="button" onClick={handleCloseModal}>
                  支付完成
                </button>
              </>
            )}
          </div>
        </div>
      )}

      <Footer />
    </div>
  );
};

export default BookingPage;
