// src/pages/PrivacyPolicy.js
import React from 'react';
import styles from './PrivacyPolicy.module.scss';

const PrivacyPolicy = () => {
  return (
    <div className={styles.privacyPolicy}>
      <h1>隐私政策</h1>
      <p>本隐私政策描述了我们如何收集、使用和保护您的个人信息。</p>
      <h2>信息收集</h2>
      <p>我们可能会收集以下类型的个人信息：</p>
      <ul>
        <li>姓名</li>
        <li>电子邮件地址</li>
        <li>联系电话</li>
      </ul>
      <h2>信息使用</h2>
      <p>我们使用收集的信息来：</p>
      <ul>
        <li>提供和改进我们的服务</li>
        <li>与您沟通</li>
        <li>发送营销信息（如果您同意）</li>
      </ul>
      <h2>信息安全</h2>
      <p>我们采取适当的安全措施来保护您的个人信息免受未经授权的访问、披露、更改或破坏。</p>
      <h2>联系我们</h2>
      <p>
        如果您对本隐私政策有任何疑问，请联系我们：
        <a href="/about-us">关于我们</a>
      </p>
    </div>
  );
};

export default PrivacyPolicy;
