import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import { Helmet } from 'react-helmet'; // 引入 react-helmet
import HomePage from './pages/HomePage'; // 直接导入 HomePage
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
import NotFoundPage from './pages/NotFoundPage'; // 新增的 404 页面

const App = () => (
  <Router>
    <Helmet>
      <title>首页</title>
      <meta
        name="description"
        content="A comprehensive parking management system for easy and efficient parking."
      />
    </Helmet>
    <Routes>
      <Route exact path="/" element={<HomePage />} /> {/* 直接使用 HomePage */}
      <Route path="/register" element={<RegisterPage />} />
      <Route path="/login" element={<LoginPage />} />
      {/* 直接暴露需要认证的页面 */}
      <Route path="/dashboard" element={<DashboardPage />} />
      <Route path="/parking-lots" element={<ParkingLotListPage />} />
      <Route path="/parking-spots/:id" element={<ParkingSpotDetailPage />} />
      <Route path="/bookings" element={<BookingPage />} />
      <Route path="/payments" element={<PaymentPage />} />
      <Route path="/parking-records" element={<ParkingRecordPage />} />
      <Route path="/admin-dashboard" element={<AdminDashboard />} />
      <Route path="/violations" element={<ViolationsPage />} />
      {/* 404 页面 */}
      <Route path="*" element={<NotFoundPage />} />
    </Routes>
  </Router>
);

export default App;
