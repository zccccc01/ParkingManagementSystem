// src/components/ViolationsList.js
import React from 'react';

const ViolationsList = ({ violations }) => {
  return (
    <ul className="violations-list">
      {violations.map((violation, index) => (
        <li key={index}>
          <p>Vehicle: {violation.vehicle}</p>
          <p>Violation: {violation.violation}</p>
          <p>Date: {violation.date}</p>
        </li>
      ))}
    </ul>
  );
};

export default ViolationsList;
