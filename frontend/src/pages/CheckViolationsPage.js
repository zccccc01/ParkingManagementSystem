import React, { useState, useEffect, useRef } from 'react';
import axios from 'axios';
import * as echarts from 'echarts'; // 确保echarts导入在其他组件导入之前
import { Table } from 'antd';
import Header from '../components/AdminHeader';
import Footer from '../components/Footer';
import './CheckViolationsPage.scss';

const CheckViolationsPage = () => {
  const [violationRecords, setViolationRecords] = useState([]);
  const [loading, setLoading] = useState(false);
  const [errorMessage, setErrorMessage] = useState(null);
  const [ViolationType, setViolationType] = useState('');
  const pieChartRef = useRef(null);

  const fetchViolationRecords = async (violationType) => {
    try {
      const response = await axios.get(`/api/violationrecord/violation/${violationType}`);
      console.log('Response data:', response.data);

      if (!response.data || !Array.isArray(response.data.violationRecords)) {
        throw new Error('No violation record data returned or data is not an array');
      }

      return response.data.violationRecords;
    } catch (error) {
      console.error('Failed to fetch violation records:', error);
      throw error;
    }
  };

  const loadViolationRecords = async () => {
    try {
      setLoading(true);
      const records = await fetchViolationRecords(ViolationType);
      setViolationRecords(records);
    } catch (error) {
      setErrorMessage(error.message);
    } finally {
      setLoading(false);
    }
  };

  const initializePieCharts = () => {
    const myChart = echarts.init(pieChartRef.current);

    // 计算每个状态的违规记录数量
    const statusCounts = violationRecords.reduce((acc, record) => {
      if (acc[record.Status]) {
        acc[record.Status] += record.TotalViolations; // 使用TotalViolations作为计数
      } else {
        acc[record.Status] = record.TotalViolations;
      }
      return acc;
    }, {});

    const statusData = Object.entries(statusCounts).map(([status, count]) => ({
      value: count,
      name: status,
    }));

    const option = {
      tooltip: {
        trigger: 'item',
        formatter: '{a} <br/>{b} : {c} ({d}%)',
      },
      legend: {
        orient: 'vertical',
        left: 'left',
        data: Object.keys(statusCounts),
      },
      series: [
        {
          name: '状态分布',
          type: 'pie',
          radius: '50%',
          data: statusData,
          emphasis: {
            itemStyle: {
              shadowBlur: 10,
              shadowOffsetX: 0,
              shadowColor: 'rgba(0, 0, 0, 0.5)',
            },
          },
        },
      ],
    };
    myChart.setOption(option);
  };

  useEffect(() => {
    if (violationRecords.length > 0) {
      initializePieCharts();
    }
  }, [violationRecords]); // 使用violationRecords作为依赖项

  const handleInputChange = (e) => {
    setViolationType(e.target.value);
  };

  const handleSearch = () => {
    if (ViolationType) {
      loadViolationRecords();
    } else {
      setErrorMessage('请选择违章类型');
    }
  };

  const columns = [
    {
      title: '违章类型',
      dataIndex: 'ViolationType',
      key: 'ViolationType',
    },
    {
      title: '状态',
      dataIndex: 'Status',
      key: 'Status',
    },
    {
      title: '总违章数',
      dataIndex: 'TotalViolations',
      key: 'TotalViolations',
    },
    {
      title: '总罚款金额',
      dataIndex: 'TotalFineAmount',
      key: 'TotalFineAmount',
    },
  ];

  return (
    <div className="violations-page">
      <Header />
      <h2>查询违规停车页面</h2>
      <p>统计和分析违规停车行为的数量及其处理情况</p>
      <label>
        违章类型:
        <select name="ViolationType" value={ViolationType} onChange={handleInputChange}>
          <option value="">请筛选违章类型</option>
          <option value="OVERSTAY">OVERSTAY</option>
          <option value="NOPAY">NOPAY</option>
        </select>
      </label>
      <button onClick={handleSearch} disabled={loading} className="search-button" type="button">
        {loading ? '查询中...' : '查询'}
      </button>

      {loading && <p>Loading...</p>}
      {!loading && errorMessage && <p className="error-message">{errorMessage}</p>}
      {!loading && !errorMessage && (
        <div>
          {violationRecords.length === 0 ? (
            <p>请选择违规类型</p>
          ) : (
            <div>
              <div ref={pieChartRef} className="pie-chart-container" />
              <Table
                columns={columns}
                dataSource={violationRecords}
                rowKey="RecordID"
                pagination={{ pageSize: 10 }}
              />
            </div>
          )}
        </div>
      )}

      <Footer />
    </div>
  );
};

export default CheckViolationsPage;
