import React, { useState } from 'react';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';
import './LoginPage.scss';

const LoginPage = () => {
  const [tel, setTel] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');
  const navigate = useNavigate();

  const handleSubmit = async (event) => {
    event.preventDefault(); // 阻止表单默认提交行为

    setError(''); // 清除之前的错误信息

    try {
      const response = await axios.post(
        'http://localhost:8000/api/user/login',
        {
          tel,
          password,
        },
        {
          withCredentials: true, // 发送带有cookie的请求
        }
      );

      if (response.status === 200) {
        navigate('/dashboard'); // 登录成功后跳转到 dashboard 页面
      } else {
        setError('登录失败，请检查用户名和密码');
      }
    } catch (apiError) {
      setError(apiError.response?.data?.message || '登录失败，请检查用户名和密码');
    }
  };

  const handleCloseClick = () => {
    navigate(-1); // 返回上一页
  };

  const handleKeyDown = (event) => {
    if (event.key === 'Enter') {
      handleSubmit(event);
    }
  };

  return (
    <div className="login-page">
      <img className="shape1" src="https://s3.us-east-2.amazonaws.com/ui.glass/shape.svg" alt="" />
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
          <h1>登录帐户</h1>
          {error && <p className="error-message">{error}</p>}
          <form onSubmit={handleSubmit}>
            <label htmlFor="tel">
              电话号码:
              <input
                type="text"
                id="tel"
                name="tel"
                value={tel}
                onChange={(e) => setTel(e.target.value)}
              />
            </label>
            <br />
            <br />
            <label htmlFor="password">
              密码:
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
            <button type="submit">登录</button>
          </form>
        </div>
      </div>
    </div>
  );
};

export default LoginPage;
