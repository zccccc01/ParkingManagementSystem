import React, { useState, useEffect } from 'react';
import axios from 'axios';
import Header from '../components/Header';
import Footer from '../components/Footer';
import './ParkingSpacePage.scss';

const ParkingSpacePage = () => {
  const [parkingSpaces, setParkingSpaces] = useState([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);
  const [viewCount, setViewCount] = useState('Loading...'); // 使用单引号

  const fetchParkingSpaces = async () => {
    setLoading(true);
    try {
      const { data } = await axios.get('http://localhost:8000/api/parkingspace/status/free');
      setParkingSpaces(data.spaces);
      setError(null);
    } catch (err) {
      console.error('Failed to fetch parking spaces:', err);
      setError(err);
    } finally {
      setLoading(false);
    }
  };

  // Fetch the view count when the component mounts
  const fetchViewCount = async () => {
    try {
      const { data } = await axios.get('http://localhost:8000/parking-space/count'); // 更新为新的端点
      setViewCount(data.count); // 从返回的 JSON 中提取 count
    } catch (err) {
      console.error('Failed to fetch view count:', err);
    }
  };

  // 增加浏览计数并获取停车位信息
  const incrementViewCount = async () => {
    try {
      await axios.get('http://localhost:8000/parking-space'); // 发送请求以增加计数
      fetchViewCount(); // 然后获取当前浏览计数
    } catch (err) {
      console.error('Failed to increment view count:', err);
    }
  };

  useEffect(() => {
    incrementViewCount(); // 在组件挂载时增加浏览计数
    fetchParkingSpaces(); // 获取停车位信息
  }, []);

  return (
    <div className="ParkingSpacePage">
      <Header />
      <h1>空闲车位信息页面</h1>
      <button type="button" onClick={fetchParkingSpaces}>
        查询
      </button>
      {loading && <p>加载中...</p>}
      {error && <p className="error">加载失败: {error.message}</p>}
      {!loading && !error && (
        <div className="two-column-table">
          <div className="column left">
            <table>
              <thead>
                <tr>
                  <th>车位 ID</th>
                  <th>状态</th>
                  <th>停车场 ID</th>
                </tr>
              </thead>
              <tbody>
                {parkingSpaces
                  .filter((_, index) => index % 2 === 0)
                  .map((space) => (
                    <tr key={space.SpaceID}>
                      <td>{space.SpaceID}</td>
                      <td>{space.Status}</td>
                      <td>{space.ParkingLotID}</td>
                    </tr>
                  ))}
              </tbody>
            </table>
          </div>
          <div className="column right">
            <table>
              <thead>
                <tr>
                  <th>车位 ID</th>
                  <th>状态</th>
                  <th>停车场 ID</th>
                </tr>
              </thead>
              <tbody>
                {parkingSpaces
                  .filter((_, index) => index % 2 !== 0)
                  .map((space) => (
                    <tr key={space.SpaceID}>
                      <td>{space.SpaceID}</td>
                      <td>{space.Status}</td>
                      <td>{space.ParkingLotID}</td>
                    </tr>
                  ))}
              </tbody>
            </table>
          </div>
          <div className="vertical-divider" />
        </div>
      )}
      <Footer />
      {/* 浏览计数显示 */}
      <div id="viewCount" className="view-count">
        浏览次数：{viewCount}
      </div>
    </div>
  );
};

export default ParkingSpacePage;
