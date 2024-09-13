// src/components/Footer.js
import React from 'react';
import styles from './Footer.module.scss';

const Footer = () => {
  return (
    <footer className={styles.footer}>
      <p>版权所有 © 2024 of the PMS Team</p>
      <p>
        <a href="/privacy-policy">隐私政策 </a>
        <a href="/terms-of-service">服务条款 </a>
        <a href="/about-us">关于我们 </a>
      </p>
    </footer>
  );
};

export default Footer;
