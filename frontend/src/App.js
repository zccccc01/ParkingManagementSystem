import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import logo from './logo.svg';
import './App.css';

// 导入公共组件
import Header from './components/Header';
import Footer from './components/Footer';

// 导入主要页面
import HomePage from './pages/HomePage';
import RegisterPage from './pages/RegisterPage';
import LoginPage from './pages/LoginPage';
import DashboardPage from './pages/DashboardPage';
import ParkingLotListPage from './pages/ParkingLotListPage';
import ParkingSpotDetailPage from './pages/ParkingSpotDetailPage';
import BookingPage from './pages/BookingPage';
import PaymentPage from './pages/PaymentPage';
import ParkingRecordPage from './pages/ParkingRecordPage';
import AdminDashboard from './pages/AdminDashboard';
import ViolationsPage from './pages/ViolationsPage';

function App() {
  return (
    <Router>
      <div className="app-container">
        <Header /> 
        <img src={logo} className="App-logo" alt="logo" />
        <Routes>
          <Route path="/" element={<HomePage />} />
          <Route path="/register" element={<RegisterPage />} />
          <Route path="/login" element={<LoginPage />} />
          <Route path="/dashboard" element={<DashboardPage />} />
          <Route path="/parking-lots" element={<ParkingLotListPage />} />
          <Route path="/parking-spot/:id" element={<ParkingSpotDetailPage />} />
          <Route path="/booking" element={<BookingPage />} />
          <Route path="/payment" element={<PaymentPage />} />
          <Route path="/parking-records" element={<ParkingRecordPage />} />
          <Route path="/admin-dashboard" element={<AdminDashboard />} />
          <Route path="/violations" element={<ViolationsPage />} />
        </Routes>
        <Footer />
      </div>
    </Router>
  );
}

export default App;