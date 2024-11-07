// src/components/SearchBar.js
import React, { useState, useEffect, useRef } from 'react';
import { useNavigate } from 'react-router-dom';
import {
  HomeOutlined,
  DashboardOutlined,
  CarOutlined,
  FileSearchOutlined,
  DollarCircleOutlined,
  CalendarOutlined,
  UserOutlined,
  LockOutlined,
} from '@ant-design/icons';
import './SearchBar.scss';

const menuList = [
  { value: '/', label: '首页', icon: <HomeOutlined /> },
  { value: '/dashboard', label: '用户仪表盘', icon: <DashboardOutlined /> },
  { value: '/parking-records', label: '停车记录', icon: <FileSearchOutlined /> },
  { value: '/create-vehicle', label: '绑定车辆', icon: <CarOutlined /> },
  { value: '/vehicle-info', label: '车辆信息', icon: <CarOutlined /> },
  { value: '/violations', label: '违章停车', icon: <LockOutlined /> },
  { value: '/bookings', label: '停车位预约', icon: <CalendarOutlined /> },
  { value: '/payments', label: '支付', icon: <DollarCircleOutlined /> },
  { value: '/parking-space', label: '空闲车位', icon: <CarOutlined /> },
  { value: '/parking-lots', label: '停车场列表', icon: <HomeOutlined /> },
  { value: '/user', label: '用户页面', icon: <UserOutlined /> },
];

const SearchBar = () => {
  const [query, setQuery] = useState('');
  const [results, setResults] = useState([]);
  const navigate = useNavigate();
  const searchBarRef = useRef(null);

  const handleSearch = (e) => {
    const searchQuery = e.target.value.toLowerCase();
    setQuery(searchQuery);

    const filteredResults = menuList.filter((item) =>
      item.label.toLowerCase().includes(searchQuery)
    );

    setResults(filteredResults);
  };

  const handleResultClick = (result) => {
    navigate(result.value);
    setQuery('');
    setResults([]);
  };

  const handleClickOutside = (event) => {
    if (searchBarRef.current && !searchBarRef.current.contains(event.target)) {
      setQuery('');
      setResults([]);
    }
  };

  useEffect(() => {
    document.addEventListener('mousedown', handleClickOutside);
    return () => {
      document.removeEventListener('mousedown', handleClickOutside);
    };
  }, []);

  return (
    <div ref={searchBarRef} className="search-bar">
      <input
        type="text"
        placeholder="Search..."
        value={query}
        onChange={handleSearch}
        className="search-input"
      />
      {results.length > 0 && (
        <ul className="search-results">
          {results.map((result) => (
            <button
              type="button"
              key={result.value}
              onClick={() => handleResultClick(result)}
              className="search-result-item"
            >
              {result.icon} {result.label}
            </button>
          ))}
        </ul>
      )}
    </div>
  );
};

export default SearchBar;
