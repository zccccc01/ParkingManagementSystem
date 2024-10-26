// TODO: 将Sidebar组件与Header组件分割开来，使得Header组件只包含导航栏的内容
// TODO: 将用户接口移动到Header组件中，使得用户信息可以在任何页面中获取，展示出用户头像（和用户名）
// TODO: Sidebar组件如应用antd样式，可能需要将js文件改为tsx文件
// TODO: Header组件应包含项目Logo名称（，以及用户登录/注册按钮）
import React from 'react';
import styles from './Header.module.scss';
import Sidebar from './Sidebar.tsx';

const Header = () => {
  return (
    <header className={styles.header}>
      <Sidebar />
      <h1>Parking Management System</h1>
      <br />
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
    </header>
  );
};

export default Header;
