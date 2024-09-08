// src/components/ParkingSpotCard.js
import React from 'react';

const ParkingSpotCard = ({ spotNumber, status }) => {
  return (
    <div className="parking-spot-card">
      <h3>Spot {spotNumber}</h3>
      <p>Status: {status}</p>
    </div>
  );
};

export default ParkingSpotCard;