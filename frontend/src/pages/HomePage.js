// pages/HomePage.js
import React from 'react';
import SearchBar from '../components/SearchBar';
import Header from '../components/Header';
import Footer from '../components/Footer';

const HomePage = () => {
  return (
    <div className="home-page">
      <Header />
      <SearchBar />
      <p>欢迎来到停车管理系统！</p>
      <Footer />
    </div>
  );
};

export default HomePage;
