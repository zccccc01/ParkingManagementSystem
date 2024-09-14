import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import './RegisterPage.scss';

const RegisterPage = () => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [confirmPassword, setConfirmPassword] = useState('');
  const navigate = useNavigate();

  const handleSubmit = async (event) => {
    event.preventDefault(); // 阻止表单默认提交行为

    // 示例逻辑：直接模拟注册成功
    localStorage.setItem('token', 'dummy-token'); // 存储 dummy token 到本地
    navigate('/dashboard'); // 注册成功后跳转到 dashboard 页面
  };

  const handleCloseClick = () => {
    navigate(-1); // 返回上一页
  };

  const handleKeyDown = (event) => {
    if (event.key === 'Enter') {
      handleCloseClick();
    }
  };

  return (
    <div className="login-page">
      <div className="container">
        <div className="modal">
          <button
            type="button"
            className="close-btn"
            onClick={handleCloseClick}
            onKeyDown={handleKeyDown}
            aria-label="close"
            title="close"
          >
            x
          </button>
          <h1>用户注册</h1>
          <form onSubmit={handleSubmit}>
            <label htmlFor="username">
              用户名：
              <input
                type="text"
                id="username"
                name="username"
                value={username}
                onChange={(e) => setUsername(e.target.value)}
              />
            </label>
            <br />
            <br />
            <label htmlFor="password">
              密码：
              <input
                type="password"
                id="password"
                name="password"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
              />
            </label>
            <br />
            <br />
            <label htmlFor="confirm-password">
              确认密码：
              <input
                type="password"
                id="confirm-password"
                name="confirm-password"
                value={confirmPassword}
                onChange={(e) => setConfirmPassword(e.target.value)}
              />
            </label>
            <br />
            <br />
            <button type="submit">注册</button>
          </form>
        </div>
      </div>
    </div>
  );
};

export default RegisterPage;
