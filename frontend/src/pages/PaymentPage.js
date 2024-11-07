// pages/PaymentPage.js
import React, { useState, useEffect } from 'react';
import axios from 'axios';
import Header from '../components/Header';
import Footer from '../components/Footer';
import './PaymentPage.scss';

const PaymentPage = () => {
  const [record, setRecord] = useState('');
  const [reservation, setReservation] = useState('');
  const [plate, setPlate] = useState('');
  const [recordStatus, setRecordStatus] = useState('');
  const [recordFee, setRecordFee] = useState('');
  const [reservationStatus, setReservationStatus] = useState('');
  const [reservationFee, setReservationFee] = useState('');
  const [plateFee, setPlateFee] = useState('');
  const [showModal, setShowModal] = useState(false);
  const [loading, setLoading] = useState(false); // 添加 loading 状态
  const [error, setError] = useState(''); // 添加错误状态
  const [selectedPaymentMethod, setSelectedPaymentMethod] = useState(''); // 添加选中的支付方式

  const fetchData = async (type, value) => {
    try {
      setLoading(true);
      setError('');

      let statusResponse;
      let feeResponse;

      if (type === 'reservation') {
        statusResponse = await axios.get(
          `http://localhost:8000/api/paymentrecord/status/reservation/${value}`
        );
        feeResponse = await axios.get(
          `http://localhost:8000/api/paymentrecord/reservation/${value}`
        );
      } else if (type === 'record') {
        statusResponse = await axios.get(
          `http://localhost:8000/api/paymentrecord/status/record/${value}`
        );
        feeResponse = await axios.get(`http://localhost:8000/api/paymentrecord/record/${value}`);
      } else if (type === 'plate') {
        feeResponse = await axios.get(`http://localhost:8000/api/paymentrecord/plate/${value}`);
      }

      if (statusResponse && statusResponse.data !== undefined) {
        if (type === 'reservation') {
          setReservationStatus(statusResponse.data);
        } else if (type === 'record') {
          setRecordStatus(statusResponse.data);
        }
      } else {
        console.log('status not found');
      }

      if (feeResponse && feeResponse.data !== undefined) {
        if (type === 'reservation') {
          setReservationFee(feeResponse.data);
        } else if (type === 'record') {
          setRecordFee(feeResponse.data);
        } else if (type === 'plate') {
          setPlateFee(feeResponse.data);
        }
      } else {
        console.log('fee not found');
      }
    } catch (err) {
      console.error('Failed to fetch data:', err.response ? err.response.data : err.message);
      setError('获取数据失败，请检查输入或稍后再试。');
    } finally {
      setLoading(false); // 确保加载状态关闭
    }
  };

  const handleInputChange = (e) => {
    const { name, value } = e.target;
    if (name === 'record') {
      setRecord(value);
      setReservation('');
      setPlate('');
      setRecordStatus('');
      setRecordFee('');
    } else if (name === 'reservation') {
      setReservation(value);
      setRecord('');
      setPlate('');
      setReservationStatus('');
      setReservationFee('');
    } else if (name === 'plate') {
      setPlate(value);
      setRecord('');
      setReservation('');
      setPlateFee('');
    } else if (name === 'paymentMethod') {
      setSelectedPaymentMethod(value);
    }
  };

  const handleQuery = () => {
    if (record) {
      fetchData('record', record);
    } else if (reservation) {
      fetchData('reservation', reservation);
    } else if (plate) {
      fetchData('plate', plate);
    }
  };

  const handleSubmit = () => {
    // 模拟支付成功
    setShowModal(true);
  };

  const handleCloseModal = () => {
    setShowModal(false);
  };

  // 使用 useEffect 监听状态变化
  useEffect(() => {
    // 重置状态
    setRecordStatus('');
    setRecordFee('');
    setReservationStatus('');
    setReservationFee('');
    setPlateFee('');
  }, [record, reservation, plate]);

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
    <div className="payment-page">
      <Header />
      <h1>支付页面</h1>
      <p>处理停车费用的支付</p>

      {error && <p className="error">{error}</p>}

      <div className="payment-form">
        <div className="payment-form-item">
          <label htmlFor="record">停车记录ID:</label>
          <input
            type="text"
            id="record"
            name="record"
            value={record}
            onChange={handleInputChange}
          />
          <label htmlFor="status">状态: {recordStatus}</label>
          <label htmlFor="fee">费用: {recordFee}</label>
        </div>
        <br />
        <div className="payment-form-item">
          <label htmlFor="reservation">预约记录ID:</label>
          <input
            type="text"
            id="reservation"
            name="reservation"
            value={reservation}
            onChange={handleInputChange}
          />
          <label htmlFor="status">状态: {reservationStatus}</label>
          <label htmlFor="fee">费用: {reservationFee}</label>
        </div>
        <br />
        <div className="payment-form-item">
          <label htmlFor="plate">车牌号码 :</label>
          <input type="text" id="plate" name="plate" value={plate} onChange={handleInputChange} />
          <label htmlFor="fee">费用: {plateFee}</label>
        </div>
        <br />
        <label>
          支付方式:
          <select name="paymentMethod" onChange={handleInputChange}>
            <option value="">请选择支付方式</option>
            <option value="FREE">微信支付</option>
            <option value="OCCUPIED">支付宝</option>
            <option value="RESERVED">信用卡</option>
          </select>
        </label>
        <br />
        <br />
        <div className="query-button-container">
          <button type="button" onClick={handleQuery} disabled={loading}>
            {loading ? '查询中...' : '查询'}
          </button>
          <button type="button" onClick={handleSubmit} disabled={loading}>
            {loading ? '加载中...' : '支付'}
          </button>
        </div>
      </div>
      {showModal && (
        <div className="modal">
          <div className="modal-content">
            <img src={getPaymentImageSrc(selectedPaymentMethod)} alt="支付方式" />
            <br />
            <button type="button" onClick={handleCloseModal}>
              支付完成
            </button>
          </div>
        </div>
      )}
      <Footer />
    </div>
  );
};

export default PaymentPage;
