// pages/RegisterPage.js
import React from 'react';
import Header from '../components/Header';
import Footer from '../components/Footer';
import './RegisterPage.scss';

const RegisterPage = () => {
  return (
    <label htmlFor="register-page">
      <Header />
      <h1>用户注册</h1>
      <form>
        <label htmlFor="username">
          用户名:
          <input id="username" type="username" />
        </label>
        <br />
        <br />
        <label htmlFor="password">
          密码:
          <input id="password" type="password" />
        </label>
        <br />
        <br />
        <button type="submit">注册</button>
      </form>
      <Footer />
    </label>
  );
};

export default RegisterPage;
