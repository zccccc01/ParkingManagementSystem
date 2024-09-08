// src/pages/LoginPage.js
import React from 'react';

const LoginPage = () => {
  return (
    <div className="login-page">
      <h1>Login</h1>
      <form>
        <label htmlFor="username">Username:</label>
        <input type="text" id="username" name="username" />
        
        <label htmlFor="password">Password:</label>
        <input type="password" id="password" name="password" />
        
        <button type="submit">Login</button>
      </form>
    </div>
  );
};

export default LoginPage;