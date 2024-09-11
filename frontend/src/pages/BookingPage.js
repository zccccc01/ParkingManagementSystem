// pages/BookingPage.js
import React from 'react';
import Header from '../components/Header';
import Footer from '../components/Footer';

const BookingPage = () => {
  return (
    <div className="booking-page">
      <Header />
      <h1>预定页面</h1>
      <p>选择停车时段和支付方式</p>
      <Footer />
    </div>
  );
};

export default BookingPage;
