// TODO: 添加一个例如https://im.qq.com/index/一样的宣传页面，添加“开始使用”的按钮，点击跳转登陆页面
// TODO: 宣传页面就相当于游客功能，登陆后才能使用完整功能
// TODO: 登录成功跳转首页
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
        navigate('/'); // 登录成功后跳转到首页
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
      <div className="square">
        <ul>
          <li />
          <li />
          <li />
          <li />
          <li />
        </ul>
      </div>
      <div className="circle">
        <ul>
          <li />
          <li />
          <li />
          <li />
          <li />
        </ul>
      </div>
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
            <footer>
              <br />
              Not a member? <a href="/register">Sign up now</a>
            </footer>
          </form>
        </div>
      </div>
    </div>
  );
};

export default LoginPage;
