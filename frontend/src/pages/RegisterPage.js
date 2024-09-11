// pages/RegisterPage.js
import React from 'react';
import Header from '../components/Header';
import Footer from '../components/Footer';

const RegisterPage = () => {
  return (
    <div className="register-page">
      <Header />
      <h1>用户注册</h1>
      <form>
        {/* eslint-disable-next-line jsx-a11y/label-has-associated-control */}
        <label>用户名：</label>
        <input type="text" id="username" name="username" required />
        <br />
        {/* eslint-disable-next-line jsx-a11y/label-has-associated-control */}
        <label>密码：</label>
        <input type="password" id="password" name="password" required />
        <br />
        {/* eslint-disable-next-line jsx-a11y/label-has-associated-control */}
        <label>邮箱：</label>
        <input type="email" id="email" name="email" required />
        <br />
        <button type="submit">注册</button>
      </form>
      <Footer />
    </div>
  );
};

export default RegisterPage;
