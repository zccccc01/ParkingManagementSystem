import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import axios from 'axios';
import './RegisterPage.scss';

const RegisterPage = () => {
  const [userID, setUserID] = useState('');
  const [tel, setTel] = useState('');
  const [password, setPassword] = useState('');
  const [confirmPassword, setConfirmPassword] = useState('');
  const [error, setError] = useState('');
  const navigate = useNavigate();

  const handleSubmit = async (event) => {
    event.preventDefault(); // 阻止表单默认提交行为

    try {
      const response = await axios.post('http://localhost:8000/api/register', {
        userID,
        tel,
        password,
        confirmPassword,
      });

      if (response.status === 201) {
        localStorage.setItem('token', 'dummy-token'); // 存储 dummy token 到本地
        navigate('/dashboard'); // 注册成功后跳转到 dashboard 页面
      } else {
        setError('注册失败，请检查输入的信息');
      }
    } catch (registrationError) {
      setError(registrationError.response.data.message || '注册失败，请检查输入的信息');
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
      <img className="shape2" src="https://s3.us-east-2.amazonaws.com/ui.glass/shape.svg" alt="" />
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
          {error && <p className="error-message">{error}</p>}
          <form onSubmit={handleSubmit}>
            <label htmlFor="userID">
              用户ID:
              <input
                type="text"
                id="userID"
                name="userID"
                value={userID}
                onChange={(e) => setUserID(e.target.value)}
              />
            </label>
            <br />
            <br />
            <label htmlFor="tel">
              电话号码:
              <input
                type="tel"
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
            <label htmlFor="confirm-password">
              确认密码:
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
