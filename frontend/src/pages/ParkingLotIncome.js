// src/pages/ParkingLotMap.js

import React, { useEffect } from 'react';
import L from 'leaflet';
import Header from '../components/Header';
import Footer from '../components/Footer';
import './ParkingLotMap.scss';

const ParkingLotMap = () => {
  useEffect(() => {
    const initMap = (data) => {
      const map = L.map('map').setView([35.0, 105.0], 4); // 中心在中国

      // 使用 OpenStreetMap 作为底图
      L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
        maxZoom: 19,
        attribution: '© OpenStreetMap',
      }).addTo(map);

      // 手动设置中国的边界
      const bounds = [
        [53.55, 73.66], // 北西
        [3.86, 135.05], // 南东
      ];
      L.rectangle(bounds, { color: '#ff7800', weight: 1 }).addTo(map);

      // 添加标记
      data.parkingLots.forEach((item) => {
        L.circle([item.Latitude, item.Longitude], {
          radius: 500, // 增加半径以提高可见性
          color: 'red',
          fillColor: '#ff5733', // 使用更亮的填充颜色
          fillOpacity: 0.8,
        }).addTo(map).bindPopup(`
          <div class="popup-content">
            <strong>停车场ID:</strong> ${item.ParkingLotID}<br/>
            <strong>名称:</strong> ${item.ParkingName}<br/>
            <strong>经纬度:</strong> (${item.Longitude}, ${item.Latitude})<br/>
            <strong>总收入:</strong> ${item.income}
          </div>
        `);
      });

      // 将地图视野设置为中国范围
      map.fitBounds(bounds);
    };

    const fetchData = async () => {
      try {
        const response = await fetch('http://localhost:8000/api/parkinglot/allincome/all');
        if (!response.ok) {
          throw new Error('Network response was not ok');
        }
        const data = await response.json();
        initMap(data);
      } catch (error) {
        console.error('There has been a problem with your fetch operation:', error);
        alert('数据获取失败，请稍后再试。');
      }
    };

    fetchData();
  }, []);

  return (
    <div className="container">
      <Header />
      <div id="map" />
      <Footer />
    </div>
  );
};

export default ParkingLotMap;
