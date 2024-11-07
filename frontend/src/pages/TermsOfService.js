// src/pages/TermsOfService.js
import React from 'react';
import styles from './TermsOfService.module.scss';

const TermsOfService = () => {
  return (
    <div className={styles.termsOfService}>
      <h1>服务条款</h1>
      <p>欢迎使用我们的停车管理系统(PMS)。请仔细阅读以下服务条款。</p>
      <h2>接受条款</h2>
      <p>使用我们的服务即表示您同意遵守这些服务条款。</p>
      <h2>服务内容</h2>
      <p>我们提供以下服务：</p>
      <ul>
        <li>查找空闲车位</li>
        <li>预订车位</li>
        <li>支付停车费用</li>
      </ul>
      <h2>用户行为</h2>
      <p>用户必须遵守以下规定：</p>
      <ul>
        <li>不得滥用服务</li>
        <li>不得进行非法活动</li>
        <li>不得侵犯他人隐私</li>
      </ul>
      <h2>责任限制</h2>
      <p>我们不承担因使用或无法使用本服务而引起的任何直接或间接损失。</p>
      <h2>联系我们</h2>
      <p>
        如果您对服务条款有任何疑问，请联系我们：
        <a href="/about-us">关于我们</a>
      </p>
    </div>
  );
};

export default TermsOfService;
