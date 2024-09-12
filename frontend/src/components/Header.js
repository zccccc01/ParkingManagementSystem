import React from 'react';
import styled from 'styled-components';
import styles from './Header.module.scss';

const StyledHeader = styled.header`
  background-color: #f8f9fa;
  padding: 1rem;
  display: flex;
  flex-direction: column;
  align-items: center;
`;

const HeaderTitle = styled.h1`
  margin-bottom: 1rem; /* 添加一些底部间距 */
`;

const HeaderNav = styled.nav`
  width: 100%;
  display: flex;
  justify-content: center;
`;

const Header = () => {
  return (
    <StyledHeader className={styles.header}>
      <HeaderTitle>Parking Management System</HeaderTitle>
      <br />
      <HeaderNav>
        <ul>
          <li>
            <a href="/">首页</a>
          </li>
          <li>
            <a href="/register">注册</a>
          </li>
          <li>
            <a href="/login">登录</a>
          </li>
          <li>
            <a href="/dashboard">仪表盘</a>
          </li>
          <li>
            <a href="/parking-lots">停车场列表</a>
          </li>
          <li>
            <a href="/admin-dashboard">管理员仪表盘</a>
          </li>
        </ul>
      </HeaderNav>
    </StyledHeader>
  );
};

export default Header;
