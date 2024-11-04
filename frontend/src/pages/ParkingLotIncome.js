// src/pages/ParkingLotMap.js

import React, { useEffect } from 'react';
import Header from '../components/AdminHeader';
import Footer from '../components/Footer';
import './ParkingLotMap.scss';

const ParkingLotMap = () => {
  useEffect(() => {
    // 声明异步加载回调函数
    window.AMapSecurityConfig = {
      securityJsCode: '0f5f46bf484a43dd2f079b6f69163e92', // 你的安全密钥
    };

    const fetchData = async () => {
      try {
        const response = await fetch('http://localhost:8000/api/parkinglot/allincome/all');
        if (!response.ok) {
          throw new Error('Network response was not ok');
        }
        return response.json();
      } catch (error) {
        console.error('There has been a problem with your fetch operation:', error);
      }
      return null; // 确保有返回值
    };

    const addMarkers = (map, data) => {
      data.parkingLots.forEach((item) => {
        const marker = new window.AMap.Marker({
          position: [item.Longitude, item.Latitude],
          content: '<div class="marker-circle"></div>', // 使用类名
          offset: new window.AMap.Pixel(-10, -10), // 使标记中心对准经纬度
        });

        // 点击事件，显示详细信息
        marker.on('click', function () {
          const infoWindow = new window.AMap.InfoWindow({
            content: `
              <div class="info-window-content">
                <strong>停车场ID:</strong> ${item.ParkingLotID}<br/>
                <strong>名称:</strong> ${item.ParkingName}<br/>
                <strong>经纬度:</strong> (${item.Longitude}, ${item.Latitude})<br/>
                <strong>总收入:</strong> ${item.income}
              </div>`,
            offset: new window.AMap.Pixel(0, -30), // 信息窗口偏移量
          });
          infoWindow.open(map, marker.getPosition());
        });

        marker.setMap(map); // 添加标记到地图
      });
    };

    window.onLoad = function () {
      const map = new window.AMap.Map('container', {
        center: [105.0, 35.0], // 中心在中国
        zoom: 4,
      });

      // 添加标记示例
      fetchData().then((data) => {
        if (data) {
          addMarkers(map, data);
        }
      });
    };

    // 加载高德地图脚本
    const script = document.createElement('script');
    script.src = `https://webapi.amap.com/maps?v=2.0&key=c7c1a9a0575d923704b38b454ecd3c11&callback=onLoad`;
    script.async = true;
    document.body.appendChild(script);

    // 清理函数
    return () => {
      document.body.removeChild(script);
      delete window.onLoad; // 清理回调函数
    };
  }, []);

  return (
    <div className="ParkingLotMap">
      <Header />
      <h1>停车场收入地图</h1>
      <div className="container">
        <div id="container" className="map-container" />
      </div>
      <Footer />
    </div>
  );
};

export default ParkingLotMap;
