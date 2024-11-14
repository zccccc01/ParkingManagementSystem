// src/pages/NotFoundPage.js
import React, { useState } from 'react';
import './NotFoundPage.scss';
import tinycolor from 'tinycolor2';

const NotFoundPage = () => {
  const [color, setColor] = useState('#000000');
  const [hue, setHue] = useState(0);

  const handleColorChange = (event) => {
    const newHue = parseInt(event.target.value, 10);
    const newColor = tinycolor({ h: newHue, s: 1, v: 1 }).toHexString();
    setColor(newColor);
    setHue(newHue);
  };

  return (
    <div className="not-found-page">
      <h2>Whoops, that page is gone.</h2>
      <p>
        You can click here to return to{' '}
        <a href="/" className="back-to-home">
          Home{' '}
        </a>
      </p>
      <p>While you’re here, feast your eyes upon these tantalizing</p>

      <p>
        popular designs matching the color{' '}
        <span className="color-value" style={{ color }}>
          {color}
        </span>
        .
      </p>
      <h1 className="artistic-text" style={{ color }}>
        404
      </h1>

      <input
        type="range"
        min="0"
        max="360"
        value={hue}
        onChange={handleColorChange}
        className="color-picker"
        style={{
          background: `linear-gradient(to right, hsl(${hue}, 100%, 50%), hsl(${hue + 180}, 100%, 50%))`,
        }}
      />
    </div>
  );
};

export default NotFoundPage;
