import React, { useState, useEffect, useRef } from 'react';
import * as echarts from 'echarts';
import Header from '../components/AdminHeader';
import Footer from '../components/Footer';
import './BarChart.scss';

const BarChart = () => {
  const chartRef = useRef(null);
  const annualChartRef = useRef(null);
  const [year, setYear] = useState('2024');
  const [month, setMonth] = useState('10');
  const [annualYear, setAnnualYear] = useState('2024');
  const [loading, setLoading] = useState(false);
  const [errorMessage, setErrorMessage] = useState(null);
  const [showMonthlyChart, setShowMonthlyChart] = useState(true);

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

  const fetchAnnualData = async (selectedYear) => {
    try {
      const response = await fetch(
        `http://localhost:8000/api/parkingrecord/year?year=${selectedYear}`
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

  const initAnnualChart = async (selectedYear) => {
    try {
      setLoading(true);
      const records = await fetchAnnualData(selectedYear);
      const lotIDs = records.map((record) => `Lot ${record.LotID}`);
      const totalIncomes = records.map((record) => record.TotalIncome);

      if (annualChartRef.current) {
        const chartInstance = echarts.init(annualChartRef.current);
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

  const handleAnnualYearChange = (e) => {
    setAnnualYear(e.target.value);
  };

  const handleSearch = () => {
    if (year && month) {
      initChart(year, month);
    } else {
      setErrorMessage('Please enter both year and month');
    }
  };

  const handleAnnualSearch = () => {
    if (annualYear) {
      initAnnualChart(annualYear);
    } else {
      setErrorMessage('Please enter a year');
    }
  };

  const handleToggleChart = () => {
    setShowMonthlyChart(!showMonthlyChart);
  };

  useEffect(() => {
    if (showMonthlyChart) {
      initChart(year, month);
    } else {
      initAnnualChart(annualYear);
    }
  }, [showMonthlyChart, year, month, annualYear]);

  return (
    <div>
      <Header />
      <div className="bar-chart-container">
        <div className="chart-title">Total Income by Lot ID</div>
        <div className="chart-subtitle">
          {showMonthlyChart ? `Year: ${year}, Month: ${month}` : `Year: ${annualYear}`}
        </div>
        <div className="input-container">
          {showMonthlyChart ? (
            <>
              <label>
                Year:
                <input type="number" value={year} onChange={handleYearChange} />
              </label>
              <label>
                Month:
                <input type="number" min="1" max="12" value={month} onChange={handleMonthChange} />
              </label>
              <button
                onClick={handleSearch}
                disabled={loading}
                className="search-button"
                type="button"
              >
                {loading ? 'Loading...' : 'Search'}
              </button>
            </>
          ) : (
            <>
              <label>
                Year:
                <input type="number" value={annualYear} onChange={handleAnnualYearChange} />
              </label>
              <button
                onClick={handleAnnualSearch}
                disabled={loading}
                className="search-button"
                type="button"
              >
                {loading ? 'Loading...' : 'Search'}
              </button>
            </>
          )}
        </div>
        {loading && <p>Loading...</p>}
        {!loading && errorMessage && <p className="error-message">{errorMessage}</p>}
        <div ref={showMonthlyChart ? chartRef : annualChartRef} className="chart-canvas" />
        <button onClick={handleToggleChart} className="toggle-button" type="button">
          {showMonthlyChart ? 'Annual Chart' : 'Monthly Chart'}
        </button>
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
