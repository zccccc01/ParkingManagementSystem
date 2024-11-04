// src/pages/AboutUsPage.js
import React from 'react';
import './AboutUsPage.scss'; // 引入 SCSS 文件

const AboutUsPage = () => {
  return (
    <div className="AboutUsPage">
      <h1>关于我们</h1>
      <p>我们是一支致力于提供高效智慧停车管理系统的团队。</p>
      <p>我们的目标是简化停车流程，提升用户体验。</p>
      <hr />
      <img
        alt="View zccccc01's full-sized avatar"
        src="https://avatars.githubusercontent.com/u/121802791?v=4"
        className="avatar avatar-user width-full border color-bg-default"
      />
      <p>Back-end and DataBase developer</p>
      <a href="https://github.com/zccccc01/" target="_blank" rel="noopener noreferrer">
        zccccc01
      </a>
      <br />
      <hr />
      <img
        alt="View Rethymus's full-sized avatar"
        src="https://avatars.githubusercontent.com/u/113156426?v=4"
        className="avatar avatar-user width-full border color-bg-default"
      />
      <p>Front-end developer and DataBase developer</p>
      <a href="https://github.com/Rethymus/" target="_blank" rel="noopener noreferrer">
        Rethymus
      </a>
      <br />
      <hr />
      <img
        alt="View Holocaust956's full-sized avatar"
        src="https://avatars.githubusercontent.com/u/176380517?v=4"
        className="avatar avatar-user width-full border color-bg-default"
      />
      <p>Back-end developer and DataBase developer</p>
      <a href="https://github.com/Holocaust956/" target="_blank" rel="noopener noreferrer">
        Holocaust956
      </a>
      <br />
      <hr />
    </div>
  );
};

export default AboutUsPage;
