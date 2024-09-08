import React from 'react';
import { Link } from 'react-router-dom';

function Header({ logo }) {
  return (
    <header className="header">
      <img src={logo} alt="Logo" className="logo" />
      <nav>
        <ul>
          <li><Link to="/">首页</Link></li>
          <li><Link to="/register">注册</Link></li>
          <li><Link to="/login">登录</Link></li>
          {/* 其他导航链接 */}
        </ul>
      </nav>
    </header>
  );
}

export default Header;