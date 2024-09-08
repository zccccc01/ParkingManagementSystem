import logo from './logo.svg';
import './App.css';

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <h1>Parking<br></br>Management<br></br>System</h1>
        <a
          className="App-link"
          href="https://www.baidu.com" // 跳转链接
          target="_blank"
          rel="noopener noreferrer"
        >
          Explore More
        </a>
      </header>
    </div>
  );
}

export default App;
