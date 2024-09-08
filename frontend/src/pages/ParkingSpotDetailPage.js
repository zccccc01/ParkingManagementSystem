// src/pages/ParkingSpotDetailPage.js
import React from 'react';

const ParkingSpotDetailPage = ({ match }) => {
  const spotId = match.params.id;
  return (
    <div className="parking-spot-detail-page">
      <h1>Parking Spot Detail</h1>
      <p>Spot ID: {spotId}</p>
    </div>
  );
};

export default ParkingSpotDetailPage;