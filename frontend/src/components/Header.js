import React from 'react';
import styles from './Header.module.scss';

const Header = () => {
  return (
    <header className={styles.header}>
      <h1>Parking Management System</h1>
      <div>
        <ul>
          <li>
            <a href="/register">注册</a>
          </li>
          <li>
            <a href="/login">登录</a>
          </li>
        </ul>
        <br />
        <br />
        <br />
      </div>
      <br />
      <br />
      <br />
      <nav className="sidebar">
        <ul>
          <li>
            <a href="/">首页</a>
          </li>
          <li>
            <a href="/dashboard">仪表盘</a>
          </li>
          {/* <li>
            <a href="/parking-spots/:id">停车位详情</a>
          </li> */}
          <li>
            <a href="/bookings">停车位预约</a>
          </li>
          <li>
            <a href="/violations">违章停车</a>
          </li>
          <li>
            <a href="/parking-records">停车记录</a>
          </li>
          <li>
            <a href="/parking-lots">停车场列表</a>
          </li>
          <li>
            <a href="/admin-dashboard">管理员仪表盘</a>
          </li>
        </ul>
      </nav>
    </header>
  );
};

export default Header;
