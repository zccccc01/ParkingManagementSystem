import React, { useState, useEffect } from 'react';
import SearchBar from '../components/SearchBar';
import Header from '../components/Header';
import Footer from '../components/Footer';
import './HomePage.module.scss'; // 导入SCSS样式

const HomePage = () => {
  const [currentImage, setCurrentImage] = useState(0);
  const images = [
    '//p3.dcarimg.com/img/tos-cn-i-dcdx/5cd480f6fcc24ad3a7353de566a7229e~360x0.webp',
    '//p3.dcarimg.com/img/tos-cn-i-dcdx/7a507b6ec14c4ead9c95c336893aa9d2~360x0.webp',
    '//p3.dcarimg.com/img/motor-mis-img/f8250b5672059b6380b04eeeb0d7fe33~360x0.webp',
    '//p3.dcarimg.com/img/tos-cn-i-dcdx/8c6e1943c2094ec2a14fd460d2c0ba4d~360x0.webp',
  ];

  useEffect(() => {
    const interval = setInterval(() => {
      setCurrentImage(currentImage === images.length - 1 ? 0 : currentImage + 1);
    }, 3000);

    return () => clearInterval(interval); // 清除定时器，防止内存泄漏
  }, [currentImage, images.length]); // 添加 images.length 到依赖数组

  return (
    <div className="HomePage">
      <Header />
      <SearchBar />
      <br />
      <br />
      <br />
      <p>欢迎来到停车管理系统！</p>
      <div>
        <img id="image-slider" width="540" height="340" src={images[currentImage]} alt="slider" />
      </div>
      <hr />
      <Footer />
    </div>
  );
};

export default HomePage;
