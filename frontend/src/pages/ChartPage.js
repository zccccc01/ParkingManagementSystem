import React, { useState, useEffect } from 'react';
import * as echarts from 'echarts';
import Header from '../components/Header';
import Footer from '../components/Footer';
import './Chart.scss';

const Chart = () => {
  const [lotId, setLotId] = useState('');
  const [chartData, setChartData] = useState(null);

  const fetchData = async () => {
    try {
      const response = await fetch(`http://localhost:8000/api/parkinglot/status/lot/${lotId}`);
      const data = await response.json();

      if (response.ok) {
        setChartData(data);
      } else {
        console.error(data.error);
      }
    } catch (error) {
      console.error('Error fetching data:', error);
    }
  };

  const renderChart = () => {
    const chartDom = document.getElementById('main');
    const myChart = echarts.init(chartDom);

    const option = {
      title: {
        text: '停车场状态',
        subtext: 'Free, Occupied, Reserved',
        left: 'center',
      },
      tooltip: {
        trigger: 'item',
      },
      legend: {
        orient: 'vertical',
        left: 'left',
      },
      series: [
        {
          name: 'Status',
          type: 'pie',
          radius: '50%',
          data: [
            { value: chartData.free, name: 'Free' },
            { value: chartData.occupied, name: 'Occupied' },
            { value: chartData.reserved, name: 'Reserved' },
          ],
          emphasis: {
            itemStyle: {
              shadowBlur: 10,
              shadowOffsetX: 0,
              shadowColor: 'rgba(0, 0, 0, 0.5)',
            },
          },
        },
      ],
      animation: true,
    };

    myChart.setOption(option);
  };

  useEffect(() => {
    if (chartData) {
      renderChart();
    }
  }, [chartData]);

  return (
    <div className="container">
      <Header />
      <div className="input-container">
        <input
          type="text"
          value={lotId}
          onChange={(e) => setLotId(e.target.value)}
          placeholder="Enter Parking Lot ID"
          className="input"
        />
        <button type="button" onClick={fetchData} className="button">
          获取状态
        </button>
      </div>
      <div id="main" className="chart-container" />
      <Footer />
    </div>
  );
};

export default Chart;
