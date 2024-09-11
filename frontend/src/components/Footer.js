import React from 'react';
import styles from './Footer.module.scss';

const Footer = () => {
  return (
    <footer className={styles.footer}>
      <p>版权所有 © 2024 of the PMS Team</p>
      <p>
        <a href="/privacy-policy">隐私政策</a>
      </p>
      <p>
        <a href="/terms-of-service">服务条款</a>
      </p>
    </footer>
  );
};

export default Footer;
