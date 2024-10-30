// src/components/SearchBar.js
import React, { useState, useEffect, useRef } from 'react';
import { useNavigate } from 'react-router-dom';
import './SearchBar.scss';

const menuList = [
  { value: '/', label: '首页', icon: '' },
  { value: '/dashboard', label: '仪表盘', icon: '' },
  { value: '/parking-records', label: '停车记录', icon: '' },
  { value: '/vehicle-info', label: '车辆信息', icon: '' },
  { value: '/violations', label: '违章停车', icon: '' },
  { value: '/parking-space', label: '停车位', icon: '' },
  { value: '/bookings', label: '停车位预约', icon: '' },
  { value: '/parking-lots', label: '停车场列表', icon: '' },
  { value: '/admin-dashboard', label: '管理员仪表盘', icon: '' },
  { value: '/space-status', label: '查询车位状态', icon: '' },
  { value: '/update-space', label: '更新车位状态', icon: '' },
  { value: '/check-space', label: '查看车位', icon: '' },
  { value: '/create-vehicle', label: '创建车位', icon: '' },
  { value: '/user', label: '用户页面', icon: '' },
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
              {result.label}
            </button>
          ))}
        </ul>
      )}
    </div>
  );
};

export default SearchBar;
