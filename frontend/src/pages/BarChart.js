import React, { useState, useEffect, useRef } from 'react';
import * as echarts from 'echarts';
import Header from '../components/AdminHeader';
import Footer from '../components/Footer';
import './BarChart.scss';

const BarChart = () => {
  const chartRef = useRef(null);
  const [year, setYear] = useState('2024');
  const [month, setMonth] = useState('10');
  const [loading, setLoading] = useState(false);
  const [errorMessage, setErrorMessage] = useState(null);

  const fetchData = async (selectedYear, selectedMonth) => {
    try {
      const response = await fetch(
        `http://localhost:8000/api/parkingrecord/month?year=${selectedYear}&month=${selectedMonth}`
      );
      const data = await response.json();
      if (!data || !Array.isArray(data.records)) {
        throw new Error('No records returned or data is not an array');
      }
      return data.records;
    } catch (error) {
      console.error('Error fetching data:', error);
      throw error;
    }
  };

  const initChart = async (selectedYear, selectedMonth) => {
    try {
      setLoading(true);
      const records = await fetchData(selectedYear, selectedMonth);
      const lotIDs = records.map((record) => `Lot ${record.LotID}`);
      const totalIncomes = records.map((record) => record.TotalIncome);

      if (chartRef.current) {
        const chartInstance = echarts.init(chartRef.current);
        const option = {
          tooltip: {
            trigger: 'axis',
            axisPointer: {
              type: 'shadow',
            },
          },
          xAxis: {
            type: 'category',
            data: lotIDs,
          },
          yAxis: {
            type: 'value',
          },
          series: [
            {
              name: 'Total Income',
              type: 'bar',
              data: totalIncomes,
              itemStyle: {
                color: 'rgba(75, 192, 192, 0.6)',
              },
            },
          ],
        };

        chartInstance.setOption(option);
      }
    } catch (error) {
      setErrorMessage(error.message);
    } finally {
      setLoading(false);
    }
  };

  const handleYearChange = (e) => {
    setYear(e.target.value);
  };

  const handleMonthChange = (e) => {
    setMonth(e.target.value);
  };

  const handleSearch = () => {
    if (year && month) {
      initChart(year, month);
    } else {
      setErrorMessage('Please enter both year and month');
    }
  };

  useEffect(() => {
    initChart(year, month);
  }, []);

  return (
    <div>
      <Header />
      <div className="bar-chart-container">
        <div className="chart-title">Total Income by Lot ID</div>
        <div className="chart-subtitle">
          Year: {year}, Month: {month}
        </div>
        <div className="input-container">
          <label>
            Year:
            <input type="number" value={year} onChange={handleYearChange} />
          </label>
          <label>
            Month:
            <input type="number" min="1" max="12" value={month} onChange={handleMonthChange} />
          </label>
          <button onClick={handleSearch} disabled={loading} className="search-button" type="button">
            {loading ? 'Loading...' : 'Search'}
          </button>
        </div>
        {loading && <p>Loading...</p>}
        {!loading && errorMessage && <p className="error-message">{errorMessage}</p>}
        <div ref={chartRef} className="chart-canvas" />
      </div>
      <br />
      <br />
      <br />
      <br />
      <Footer />
    </div>
  );
};

export default BarChart;
