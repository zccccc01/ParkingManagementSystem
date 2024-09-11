// pages/ViolationsPage.js
import React from 'react';
import Header from '../components/Header';
import Footer from '../components/Footer';

const ViolationsPage = () => {
  return (
    <div className="violations-page">
      <Header />
      <h1>违规停车页面</h1>
      <p>展示违规记录和处罚信息</p>
      <Footer />
    </div>
  );
};

export default ViolationsPage;
