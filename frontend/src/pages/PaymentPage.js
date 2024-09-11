// pages/PaymentPage.js
import React from 'react';
import Header from '../components/Header';
import Footer from '../components/Footer';

const PaymentPage = () => {
  return (
    <div className="payment-page">
      <Header />
      <h1>支付页面</h1>
      <p>处理停车费用的支付</p>
      <Footer />
    </div>
  );
};

export default PaymentPage;
