// src/components/ParkingLotCard.js
import React from 'react';

const ParkingLotCard = ({ name, address }) => {
  return (
    <div className="parking-lot-card">
      <h3>{name}</h3>
      <p>{address}</p>
    </div>
  );
};

export default ParkingLotCard;