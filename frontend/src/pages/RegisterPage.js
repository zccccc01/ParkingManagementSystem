import React, { useState } from 'react';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';
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

    setError(''); // 清除之前的错误信息

    // 验证 UserID
    const parsedUserID = parseInt(userID, 10);
    if (!parsedUserID || parsedUserID <= 0) {
      setError('用户ID必须是一个有效的正整数');
      return;
    }

    // 验证密码
    if (password !== confirmPassword) {
      setError('密码和确认密码不一致');
      return;
    }

    try {
      const response = await axios.post(
        '/api/user/register',
        {
          id: parsedUserID,
          tel,
          password,
        },
        {
          withCredentials: true, // 发送带有cookie的请求
        }
      );

      if (response.status === 201) {
        navigate('/login'); // 注册成功后跳转到登录页面
      } else {
        setError('注册失败，请检查输入的信息');
      }
    } catch (registrationError) {
      setError(registrationError.response?.data?.message || '注册失败，请检查输入的信息');
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
    <div className="register-page">
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
          <h1>用户注册</h1>
          {error && <p className="error-message">{error}</p>}
          <form onSubmit={handleSubmit}>
            <label htmlFor="userID">
              用户ID:
              <input
                type="number"
                id="userID"
                name="userID"
                value={userID}
                onChange={(e) => setUserID(e.target.value)}
                required
                min="1"
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
                required
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
                required
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
                required
              />
            </label>
            <br />
            <br />
            <button type="submit">注册</button>
            <footer>
              <br />
              Already have an account? <a href="/login">Sign in now</a>
            </footer>
          </form>
        </div>
      </div>
    </div>
  );
};

export default RegisterPage;
