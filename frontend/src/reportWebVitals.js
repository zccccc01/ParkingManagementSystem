// 定义 reportWebVitals 函数
/* eslint-disable no-console */
const reportWebVitals = (onPerfEntry) => {
  if (typeof onPerfEntry === 'function') {
    import('web-vitals').then(({ getCLS, getFID, getFCP, getLCP, getTTFB }) => {
      getCLS(onPerfEntry);
      getFID(onPerfEntry);
      getFCP(onPerfEntry);
      getLCP(onPerfEntry);
      getTTFB(onPerfEntry);
    });
  }
};

// 调用默认的 reportWebVitals 函数
reportWebVitals(console.log);

// 导出自定义的 reportWebVitals 函数
export default reportWebVitals;
