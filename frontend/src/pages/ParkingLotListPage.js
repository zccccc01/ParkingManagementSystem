import React, { useState, useEffect } from 'react';
import axios from 'axios';
import Header from '../components/Header';
import Footer from '../components/Footer';
import './ParkingLotListPage.scss';

const ParkingLotListPage = () => {
  const [parkingLots, setParkingLots] = useState([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);
  const [viewCount, setViewCount] = useState('Loading...'); // 使用单引号

  const fetchParkingLots = async () => {
    setLoading(true);
    try {
      const response = await axios.get('http://localhost:8000/api/parkinglot');
      setParkingLots(response.data);
      setError(null);
    } catch (err) {
      console.error('Failed to fetch parking lots:', err);
      setError(err);
    } finally {
      setLoading(false);
    }
  };

  // Fetch the view count when the component mounts
  const fetchViewCount = async () => {
    try {
      const { data } = await axios.get('http://localhost:8000/parking-lots/count'); // 更新为新的端点
      setViewCount(data.count); // 从返回的 JSON 中提取 count
    } catch (err) {
      console.error('Failed to fetch view count:', err);
    }
  };

  // 增加浏览计数并获取停车场信息
  const incrementViewCount = async () => {
    try {
      await axios.get('http://localhost:8000/parking-lots'); // 发送请求以增加计数
      fetchViewCount(); // 然后获取当前浏览计数
    } catch (err) {
      console.error('Failed to increment view count:', err);
    }
  };

  useEffect(() => {
    incrementViewCount(); // 在组件挂载时增加浏览计数
    fetchParkingLots(); // 获取停车场信息
  }, []);

  return (
    <div className="parking-lot-list-page">
      <Header />
      <h1>停车场列表</h1>
      <button type="button" onClick={fetchParkingLots}>
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
                  <th>名称</th>
                  <th>经度</th>
                  <th>纬度</th>
                  <th>总容量</th>
                  <th>收费标准（元/h）</th>
                </tr>
              </thead>
              <tbody>
                {parkingLots
                  .filter((_, index) => index % 2 === 0)
                  .map((lot) => (
                    <tr key={lot.id}>
                      <td>{lot.ParkingName}</td>
                      <td>{lot.Longitude}</td>
                      <td>{lot.Latitude}</td>
                      <td>{lot.Capacity}</td>
                      <td>{lot.Rates}</td>
                    </tr>
                  ))}
              </tbody>
            </table>
          </div>
          <div className="column right">
            <table>
              <thead>
                <tr>
                  <th>名称</th>
                  <th>经度</th>
                  <th>纬度</th>
                  <th>总容量</th>
                  <th>收费标准（元/h）</th>
                </tr>
              </thead>
              <tbody>
                {parkingLots
                  .filter((_, index) => index % 2 !== 0)
                  .map((lot) => (
                    <tr key={lot.id}>
                      <td>{lot.ParkingName}</td>
                      <td>{lot.Longitude}</td>
                      <td>{lot.Latitude}</td>
                      <td>{lot.Capacity}</td>
                      <td>{lot.Rates}</td>
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

export default ParkingLotListPage;
